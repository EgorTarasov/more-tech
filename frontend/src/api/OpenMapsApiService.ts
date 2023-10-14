import axios from 'axios';
import { IRoute } from './models';
import { OPEN_MAPS_API_URL } from '../config';

export class OpenMapsAipService {
    async fetchRoute(start: string, end: string): Promise<IRoute> {
        const response = await axios.get<IRoute>(
            `${OPEN_MAPS_API_URL}/ors/v2/directions/foot-walking?start=${start}&end=${end}`
        );

        return response.data;
    }
}

export const OpenMapsAipServiceInstanse = new OpenMapsAipService();
