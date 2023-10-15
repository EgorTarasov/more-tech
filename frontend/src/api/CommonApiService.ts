import axios from 'axios';
import { API_URL } from '../config';
import { ISearchResponse } from './models';

export class CommonApiService {
    public async getUser(): Promise<void> {
        await axios.get<void>(`${API_URL}/v1/users/1`).catch((err) => console.log(err));

        return;
    }

    async search(text: string, latitude: number, longitude: number): Promise<ISearchResponse> {
        const response = await axios.post<ISearchResponse>(`${API_URL}/v1/search?`, {
            text,
            coordinates: {
                latitude,
                longitude,
            },
            test: false,
        });

        return response.data;
    }
}

export const CommonApiServiceInstanse = new CommonApiService();
