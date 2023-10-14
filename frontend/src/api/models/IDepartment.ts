export interface IDepartment {
    _id: string;
    shortName: string;
    schedulefl: string;
    schedulejurl: string;
    special: ISpecial;
    distance: number;
    Location: ILocation;
    schedule: ISchedule[];
}

export interface ISpecial {
    vipzone: number;
    vipoffice: number;
    ramp: number;
    person: number;
    juridical: number;
    prime: number;
}

export interface ICoordinates {
    latitude: number;
    longitude: number;
}

export interface ILocation {
    type: string;
    Coordinates: ICoordinates;
}

export interface ISchedule {
    day: string;
    loadhours: Iloadhours[];
}

export interface Iloadhours {
    hour: string;
    load: number;
}
