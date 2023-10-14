export default function distanceConverter(distance: number): string {
    if (distance < 1) {
        return (distance * 1000).toFixed(0) + ' м';
    } else {
        return distance.toFixed(1) + ' км';
    }
}
