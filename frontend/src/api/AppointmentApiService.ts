import axios from 'axios';
import { API_URL } from '../config';
import { ITicket } from './models/ITicket';

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
    ): Promise<ITicket> {
        const response = await axios.post<ITicket>(`${API_URL}/v1/tickets`, {
            departmentId,
            timeSlot,
            startLatitude,
            startLongitude,
        });

        return response.data;
    }
}

export const AppointmentApiServiceInstanse = new AppointmentApiService();
