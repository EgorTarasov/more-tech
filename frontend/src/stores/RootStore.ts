import { makeAutoObservable, observable, runInAction } from 'mobx';
import { IDepartment, IDepartmentDetails } from '../api/models';
import { DepartmentsApiServiceInstanse } from '../api/DepartmentsApiService';
import { ILineString, IMapLocation } from '../models';
import { OpenMapsAipServiceInstanse } from '../api/OpenMapsApiService';
import { mapCoordsToString } from '../utils/mapCoordsToString';
import { mapRouteToCoords } from '../utils/mapRouteToCoords';
import { IFilter } from '../models/Filters';
import { CommonApiServiceInstanse } from '../api/CommonApiService';

export class RootStore {
    departments: IDepartment[] = [];
    filteredDepartments: IDepartment[] = [];
    selectedDepartment: IDepartment | null = null;
    selectedDepartmentDetails: IDepartmentDetails | null = null;
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
    openFilterTrigger: boolean | null = null;
    openAppointmentTrigger: boolean | null = null;
    filters: IFilter = {
        special: {
            vipZone: null,
            vipOffice: null,
            ramp: null,
            person: null,
            juridical: null,
            Prime: null,
        },
        raitingMoreThan4: null,
        raitingMoreThan45: null,
    };

    constructor() {
        makeAutoObservable(this, {
            departments: observable,
            filteredDepartments: observable,
            mapLocation: observable,
            selectedDepartment: observable,
            polylyne: observable,
            start: observable,
            openFilterTrigger: observable,
            openAppointmentTrigger: observable,
            filters: observable,
        });
    }

    setDepartments(departments: IDepartment[]) {
        runInAction(() => {
            this.departments = departments;
            this.filteredDepartments = departments;
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

    setFilters(filters: IFilter) {
        runInAction(() => {
            this.filteredDepartments = this.departments.filter((department) => {
                if (filters.raitingMoreThan4 && department.rating < 4) {
                    return false;
                }

                if (filters.raitingMoreThan45 && department.rating < 4.5) {
                    return false;
                }

                if (filters.special.vipZone && department.special.vipZone === 0) {
                    return false;
                }

                if (filters.special.vipOffice && department.special.vipOffice === 0) {
                    return false;
                }

                if (filters.special.ramp && department.special.ramp === 0) {
                    return false;
                }

                if (filters.special.person && department.special.person === 0) {
                    return false;
                }

                if (filters.special.juridical && department.special.juridical === 0) {
                    return false;
                }

                if (filters.special.Prime && department.special.Prime === 0) {
                    return false;
                }

                return true;
            });
        });

        runInAction(() => {
            this.filters = filters;
        });
    }

    triggerFilter() {
        runInAction(() => {
            this.openFilterTrigger = !this.openFilterTrigger;
        });
    }

    triggerAppointment() {
        runInAction(() => {
            this.openAppointmentTrigger = !this.openAppointmentTrigger;
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

    async fetchUser() {
        const user = await CommonApiServiceInstanse.getUser();

        return user;
    }

    async fetchDepartmentDetails() {
        if (this.selectedDepartment === null) {
            return;
        }

        const details = await DepartmentsApiServiceInstanse.getDepartment(
            this.selectedDepartment?._id ?? ''
        );

        runInAction(() => {
            this.selectedDepartmentDetails = details;
        });

        return details;
    }
}
