function distanceBetweenBusStops(distance: number[], start: number, destination: number): number {
    const s = Math.min(start, destination);
    const d = Math.max(start, destination);
    
    let totalDistance = 0;
    let clockwiseDistance = 0;
    
    for (let i = 0; i < distance.length; i++) {
        totalDistance += distance[i];
        if (i >= s && i < d) {
            clockwiseDistance += distance[i];
        }
    }
    
    return Math.min(clockwiseDistance, totalDistance - clockwiseDistance);
}