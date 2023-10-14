import axios from 'axios';
import { API_URL } from '../config';

export class CommonApiService {
    public async getUser(): Promise<void> {
        await axios.get<void>(`${API_URL}/v1/users/1`).catch((err) => console.log(err));

        return;
    }
}

export const CommonApiServiceInstanse = new CommonApiService();
