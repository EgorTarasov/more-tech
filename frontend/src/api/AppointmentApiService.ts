import axios from 'axios';
import { API_URL } from '../config';

export class AppointmentApiService {
    async getAppointment(id: string): Promise<void> {
        await axios.get<void>(`${API_URL}/v1/tickets/${id}`).catch((err) => console.log(err));

        return;
    }

    async createAppointment(
        departmentId: string,
        timeSlot: string,
        startLatitude: number,
        startLongitude: number
    ): Promise<void> {
        const response = await axios.post<void>(`${API_URL}/v1/tickets`, {
            departmentId,
            timeSlot,
            startLatitude,
            startLongitude,
        });

        return response.data;
    }
}

export const AppointmentApiServiceInstanse = new AppointmentApiService();
