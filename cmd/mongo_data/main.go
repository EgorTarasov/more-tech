package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Office struct {
	SalePointName       string                `json:"salePointName"`
	Address             string                `json:"address"`
	Status              string                `json:"status"`
	OpenHours           []OpenHours           `json:"openHours"`
	Rko                 string                `json:"rko"`
	OpenHoursIndividual []OpenHoursIndividual `json:"openHoursIndividual"`
	OfficeType          string                `json:"officeType"`
	SalePointFormat     string                `json:"salePointFormat"`
	SuoAvailability     string                `json:"suoAvailability"`
	HasRamp             string                `json:"hasRamp"`
	Latitude            float64               `json:"latitude"`
	Longitude           float64               `json:"longitude"`
	MetroStation        string                `json:"metroStation"`
	Distance            int                   `json:"distance"`
	Kep                 bool                  `json:"kep"`
	MyBranch            bool                  `json:"myBranch"`
}

type OpenHours struct {
	Days  string `json:"days"`
	Hours string `json:"hours"`
}

type OpenHoursIndividual struct {
	Days  string `json:"days"`
	Hours string `json:"hours"`
}

type HourWorkload struct {
	Hour string  `json:"hour"`
	Load float64 `json:"load"`
}

type Workload struct {
	Day       string         `json:"day"`
	LoadHours []HourWorkload `json:"loadHours"`
}

type Location struct {
	Type        string      `json:"type"`
	Coordinates Coordinates `json:"coordinates"`
}

type Department struct {
	ID           int         `json:"id"`
	BiskvitID    string      `json:"Biskvit_id"`
	ShortName    string      `json:"shortName"`
	Address      string      `json:"address"`
	City         string      `json:"city"`
	ScheduleFl   string      `json:"scheduleFl"`
	ScheduleJurL string      `json:"scheduleJurL"`
	Special      Special     `json:"special"`
	Coordinates  Coordinates `json:"coordinates"`
	Location     Location    `json:"location"`
	Workload     []Workload  `json:"workload"`
}

type Special struct {
	VipZone   int `json:"vipZone"`
	VipOffice int `json:"vipOffice"`
	Ramp      int `json:"ramp"`
	Person    int `json:"person"`
	Juridical int `json:"juridical"`
	Prime     int `json:"Prime"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Atm struct {
	Address   string   `json:"address"`
	Latitude  float64  `json:"latitude"`
	Longitude float64  `json:"longitude"`
	AllDay    bool     `json:"allDay"`
	Services  Services `json:"services"`
}

type Wheelchair struct {
	ServiceCapability string `json:"serviceCapability"`
	ServiceActivity   string `json:"serviceActivity"`
}

type Blind struct {
	ServiceCapability string `json:"serviceCapability"`
	ServiceActivity   string `json:"serviceActivity"`
}

type NfcForBankCards struct {
	ServiceCapability string `json:"serviceCapability"`
	ServiceActivity   string `json:"serviceActivity"`
}

type QrRead struct {
	ServiceCapability string `json:"serviceCapability"`
	ServiceActivity   string `json:"serviceActivity"`
}

type SupportsUsd struct {
	ServiceCapability string `json:"serviceCapability"`
	ServiceActivity   string `json:"serviceActivity"`
}

type SupportsChargeRub struct {
	ServiceCapability string `json:"serviceCapability"`
	ServiceActivity   string `json:"serviceActivity"`
}

type SupportsEur struct {
	ServiceCapability string `json:"serviceCapability"`
	ServiceActivity   string `json:"serviceActivity"`
}

type SupportsRub struct {
	ServiceCapability string `json:"serviceCapability"`
	ServiceActivity   string `json:"serviceActivity"`
}

type Services struct {
	Wheelchair        Wheelchair        `json:"wheelchair"`
	Blind             Blind             `json:"blind"`
	NfcForBankCards   NfcForBankCards   `json:"nfcForBankCards"`
	QrRead            QrRead            `json:"qrRead"`
	SupportsUsd       SupportsUsd       `json:"supportsUsd"`
	SupportsChargeRub SupportsChargeRub `json:"supportsChargeRub"`
	SupportsEur       SupportsEur       `json:"supportsEur"`
	SupportsRub       SupportsRub       `json:"supportsRub"`
}

func getOffice(doc Department, offices *[]Office) (*Office, error) {
	for idx, office := range *offices {
		if office.SalePointName == doc.ShortName {
			*offices = append((*offices)[:idx], (*offices)[idx+1:]...)
			fmt.Println(len(*offices))
			return &office, nil
		}
	}
	return nil, errors.New("office not found")
}

func loadAtms(client *mongo.Client) {
	data, err := os.ReadFile("data/atms.json")
	if err != nil {
		panic(err)
	}

	// TODO: сделать нормальный парсинг
	atms := []map[string]any{}
	err = json.Unmarshal([]byte(data), &atms)
	if err != nil {
		panic(err)
	}
	atms_result := make([]interface{}, len(atms))
	for idx, atm := range atms {
		atms_result[idx] = atm
	}

	_, err = client.Database("dev").Collection("atms").InsertMany(context.Background(), atms_result)
	if err != nil {
		panic(err)
	}
}

func loadDepartments(client *mongo.Client) {
	data, err := os.ReadFile("data/data.json")
	if err != nil {
		panic(err)
	}

	departments := []Department{}
	err = json.Unmarshal([]byte(data), &departments)
	if err != nil {
		panic(err)
	}

	data, err = os.ReadFile("data/offices.json")
	if err != nil {
		panic(err)
	}

	offices := []Office{}
	err = json.Unmarshal([]byte(data), &offices)
	if err != nil {
		panic(err)
	}

	result := make([]interface{}, 0)
	for _, doc := range departments {
		office, err := getOffice(doc, &offices)
		if err != nil {
			continue
		}

		openHours := office.OpenHoursIndividual

		workload := []Workload{}
		for _, day := range openHours {
			hoursLoad := []HourWorkload{}

			hours := day.Hours
			if hours != "выходной" && hours != "" {
				startString, endString := strings.Split(hours, "-")[0], strings.Split(hours, "-")[1]

				startHours, err := strconv.Atoi(strings.Split(startString, ":")[0])
				if err != nil {
					panic(err)
				}

				endHours, err := strconv.Atoi(strings.Split(endString, ":")[0])
				if err != nil {
					panic(err)
				}

				startMinutes := strings.Split(startString, ":")[1]
				endMinutes := strings.Split(endString, ":")[1]

				for i := startHours; i < endHours; i++ {
					if i == endHours-1 {
						hoursLoad = append(hoursLoad, HourWorkload{Hour: fmt.Sprintf("%d:%s-%d:%s", i, startMinutes, i+1, endMinutes), Load: rand.Float64()})
					} else {
						hoursLoad = append(hoursLoad, HourWorkload{Hour: fmt.Sprintf("%d:%s-%d:%s", i, startMinutes, i+1, startMinutes), Load: rand.Float64()})
					}
				}
			}

			workload = append(workload, Workload{
				Day:       day.Days,
				LoadHours: hoursLoad,
			})
		}
		doc.Workload = workload

		doc.Location = Location{
			Type:        "Point",
			Coordinates: doc.Coordinates,
		}
		result = append(result, doc)
	}

	_, err = client.Database("dev").Collection("departments").InsertMany(context.Background(), result)
	if err != nil {
		panic(err)
	}

	// save offices into unmatched.json
	unmatched := make([]interface{}, len(offices))
	for idx, office := range offices {
		unmatched[idx] = office
	}
	data, err = json.Marshal(unmatched)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("data/unmatched.json", data, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	mongoConfig := make(map[string]string)
	mongoConfig["user"] = "mongouser"
	mongoConfig["password"] = "mongopass"
	mongoConfig["host"] = "localhost"
	mongoConfig["port"] = "27017"

	for _, arg := range os.Args {
		splitted := strings.Split(arg, "=")
		if len(splitted) < 2 {
			continue
		}
		key, value := splitted[0], splitted[1]
		mongoConfig[key] = value
	}

	ctx := context.Background()

	connString := fmt.Sprintf("mongodb://%s:%s@%s:%s/", mongoConfig["user"], mongoConfig["password"], mongoConfig["host"], mongoConfig["port"])
	fmt.Println(connString)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connString))
	if err != nil {
		panic(err)
	}

	loadAtms(client)
	loadDepartments(client)

}
