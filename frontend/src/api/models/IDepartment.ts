export interface IDepartment {
    _id: string;
    shortName: string;
    schedulefl: string;
    schedulejurl: string;
    special: ISpecial;
    distance: number;
    location: ILocation;
    schedule: ISchedule[];
}

export interface ISpecial {
    vipZone: number;
    vipOffice: number;
    ramp: number;
    person: number;
    juridical: number;
    Prime: number;
}

export interface ICoordinates {
    latitude: number;
    longitude: number;
}

export interface ILocation {
    type: string;
    coordinates: ICoordinates;
}

export interface ISchedule {
    day: string;
    loadhours: Iloadhours[];
}

export interface Iloadhours {
    hour: string;
    load: number;
}
