export interface IFilter {
    special: {
        vipZone: boolean | null;
        vipOffice: boolean | null;
        ramp: boolean | null;
        person: boolean | null;
        juridical: boolean | null;
        Prime: boolean | null;
    };
    raitingMoreThan4: boolean | null;
    raitingMoreThan45: boolean | null;
}
