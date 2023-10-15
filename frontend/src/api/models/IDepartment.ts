export interface IDepartment {
    _id: string;
    address: string;
    shortName: string;
    schedulefl: string;
    schedulejurl: string;
    special: ISpecial;
    distance: number;
    rating: number;
    location: ILocation;
    schedule: ISchedule[];
    favourite: boolean;
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

export interface IDepartmentDetails {
    _id: string;
    address: string;
    shortName: string;
    schedulefl: string;
    schedulejurl: string;
    special: ISpecial;
    distance: number;
    rating: number;
    location: ILocation;
    schedule: ISchedule[];
    workload: IWorkload[];
    estimatedTimeCar: number;
    estimatedTimeWalk: number;
}

export interface IWorkload {
    day: string;
    loadHours: Iloadhours[];
}

export interface ISearchResponse {
    _id: string;
    text: string;
    userId: string;
    coordinates: ICoordinates;
    createdAt: string;
    special: ISpecialBool;
    atm: boolean;
    online: boolean;
}

export interface ISpecialBool {
    vipZone: boolean;
    vipOffice: boolean;
    ramp: boolean;
    person: boolean;
    juridical: boolean;
    Prime: boolean;
}
