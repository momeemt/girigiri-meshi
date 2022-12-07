export type Position = {
    longitude: number;
    latitude: number;
};

export function GetPosition(): Promise<Position> {
    return new Promise((resolve, reject) => {
        if (!navigator.geolocation) {
            reject("Geolocation is not supported by your browser");
        }
        navigator.geolocation.getCurrentPosition(
            (position) => {
                resolve({
                    longitude: position.coords.longitude,
                    latitude: position.coords.latitude,
                });
            },
            (error) => {
                reject(error.message);
            }
        );
    });
}
