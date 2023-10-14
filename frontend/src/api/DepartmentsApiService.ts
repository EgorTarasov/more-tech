import axios from 'axios';
import { API_URL } from '../config';
import { IDepartment, IDepartmentDetails } from './models';

export class DepartmentsApiService {
    public async getDepartments(): Promise<IDepartment[]> {
        const response = await axios.post<IDepartment[]>(`${API_URL}/v1/departments`, {
            latitude: 55.755864,
            longitude: 37.617698,
            radius: 50,
        });

        return response.data;
    }

    public async getDepartment(id: string): Promise<IDepartmentDetails> {
        const response = await axios.get<IDepartmentDetails>(`${API_URL}/v1/departments/${id}`);

        return response.data;
    }
}

export const DepartmentsApiServiceInstanse = new DepartmentsApiService();
