import axios from 'axios';
import { API_URL } from '../config';
import { IDepartment } from './models';

export class DepartmentsApiService {
    public async getDepartments(): Promise<IDepartment[]> {
        const response = await axios.post<IDepartment[]>(`${API_URL}/v1/departments`, {
            latitude: 55.755864,
            longitude: 37.617698,
            radius: 3,
        });

        return response.data;
    }
}

export const DepartmentsApiServiceInstanse = new DepartmentsApiService();
