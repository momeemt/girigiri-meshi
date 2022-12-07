export type Position = {
    longitude: number;
    latitude: number;
};

export function GetPosition(): Promise<Position> {
    return new Promise((resolve, reject) => {
        navigator.geolocation.getCurrentPosition(
            (position) => {
                resolve({
                    longitude: position.coords.longitude,
                    latitude: position.coords.latitude,
                });
            },
            (error) => {
                reject(error);
            }
        );
    });
}
