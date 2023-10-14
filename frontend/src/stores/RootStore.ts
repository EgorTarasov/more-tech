import { makeAutoObservable, observable, runInAction } from 'mobx';
import { IDepartment } from '../api/models';
import { DepartmentsApiServiceInstanse } from '../api/DepartmentsApiService';
import { ILineString, IMapLocation } from '../models';
import { OpenMapsAipServiceInstanse } from '../api/OpenMapsApiService';
import { mapCoordsToString } from '../utils/mapCoordsToString';
import { mapRouteToCoords } from '../utils/mapRouteToCoords';

export class RootStore {
    departments: IDepartment[] = [];
    selectedDepartment: IDepartment | null = null;
    mapLocation: IMapLocation = {
        center: [37.617698, 55.755864],
        zoom: 11,
    };
    polylyne: ILineString | null = {
        id: 'route',
        geometry: {
            type: 'LineString',
            coordinates: [],
        },
        style: { stroke: [{ color: '#092896', width: 4 }] },
    };
    start: [number, number] = [37.617698, 55.755864];

    constructor() {
        makeAutoObservable(this, {
            departments: observable,
            mapLocation: observable,
            selectedDepartment: observable,
            polylyne: observable,
            start: observable,
        });
    }

    setDepartments(departments: IDepartment[]) {
        runInAction(() => {
            this.departments = departments;
        });
    }

    setSelectedDepartment(department: IDepartment | null) {
        runInAction(() => {
            this.selectedDepartment = department;
        });
    }

    setMapLocation(mapLocation: IMapLocation) {
        runInAction(() => {
            this.mapLocation = mapLocation;
        });
    }

    setPolylyne(polylyne: ILineString | null) {
        runInAction(() => {
            this.polylyne = polylyne;
        });
    }

    async fetchDepartments() {
        const departments = await DepartmentsApiServiceInstanse.getDepartments();

        this.setDepartments(departments.sort((a, b) => a.distance - b.distance));

        return departments;
    }

    async fetchRoute() {
        if (this.selectedDepartment === null) {
            return;
        }

        const route = await OpenMapsAipServiceInstanse.fetchRoute(
            mapCoordsToString(this.start),
            mapCoordsToString([
                this.selectedDepartment?.location.coordinates.longitude ?? 0,
                this.selectedDepartment?.location.coordinates.latitude ?? 0,
            ])
        );

        this.setPolylyne({
            id: 'route',
            geometry: {
                type: 'LineString',
                coordinates: mapRouteToCoords(route),
            },
            style: { stroke: [{ color: '#092896', width: 4 }] },
        });

        this.setMapLocation({
            center: [
                this.selectedDepartment.location.coordinates.longitude,
                this.selectedDepartment.location.coordinates.latitude,
            ],
            zoom: 16,
        });

        return route;
    }
}
