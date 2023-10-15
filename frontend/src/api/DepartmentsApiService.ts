import axios from 'axios';
import { API_URL } from '../config';
import { IDepartment, IDepartmentDetails } from './models';

export class DepartmentsApiService {
    public async getDepartments(latitude: number, longitude: number): Promise<IDepartment[]> {
        const response = await axios.post<IDepartment[]>(`${API_URL}/v1/departments`, {
            latitude,
            longitude,
            radius: 20,
        });

        return response.data;
    }

    public async getDepartment(
        id: string,
        startLatitude: number,
        startLongitude: number
    ): Promise<IDepartmentDetails> {
        const response = await axios.get<IDepartmentDetails>(`${API_URL}/v1/departments/${id}`, {
            params: {
                startLatitude,
                startLongitude,
            },
        });

        return response.data;
    }

    public async postDepartmentRating(departmentId: string, rating: number): Promise<void> {
        const response = await axios.post<void>(`${API_URL}/v1/departments/rating`, {
            departmentId,
            rating,
            text: 'text',
        });

        return response.data;
    }

    public async setAsFavorite(departmentId: string): Promise<void> {
        const response = await axios.post<void>(
            `${API_URL}/v1/departments/favourite/${departmentId}`
        );

        return response.data;
    }
}

export const DepartmentsApiServiceInstanse = new DepartmentsApiService();
