import { LineStringGeometry } from '@yandex/ymaps3-types';

export interface ILineString {
    id: string;
    geometry: LineStringGeometry;
    style: unknown;
}
