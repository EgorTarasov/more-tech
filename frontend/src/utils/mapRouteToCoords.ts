import { IRoute } from '../api/models';

export function mapRouteToCoords(route: IRoute): [number, number][] {
    return route.features.length
        ? route.features[0].geometry.coordinates.map((coords) => [coords[0], coords[1]])
        : [];
}
