import { makeAutoObservable, observable, runInAction } from 'mobx';
import { IDepartment } from '../api/models';
import { DepartmentsApiServiceInstanse } from '../api/DepartmentsApiService';
import { IMapLocation } from '../models';

export class RootStore {
    departments: IDepartment[] = [];
    selectedDepartment: IDepartment | null = null;
    mapLocation: IMapLocation = {
        center: [37.617698, 55.755864],
        zoom: 11,
    };

    constructor() {
        makeAutoObservable(this, {
            departments: observable,
            mapLocation: observable,
            selectedDepartment: observable,
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

    async fetchDepartments() {
        const departments = await DepartmentsApiServiceInstanse.getDepartments();

        this.setDepartments(departments.sort((a, b) => a.distance - b.distance));

        return departments;
    }
}
