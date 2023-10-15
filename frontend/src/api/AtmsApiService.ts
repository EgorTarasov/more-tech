import axios from 'axios';
import { API_URL } from '../config';
import { IAtm } from './models/IAtm';

export class AtmApiService {
    public async getAtms(latitude: number, longitude: number): Promise<IAtm[]> {
        const response = await axios.post<IAtm[]>(`${API_URL}/v1/atms`, {
            latitude,
            longitude,
            radius: 20,
        });

        return response.data;
    }
}

export const AtmApiServiceInstanse = new AtmApiService();
