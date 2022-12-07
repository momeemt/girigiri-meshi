import { atom } from "recoil";

export type Position = [number, number];

export const PositionAtom = atom({
    key: "position",
    default: {} as Position,
});

export function GetUserPosition(): Promise<Position> {
    return new Promise((resolve, reject) => {
        if (!navigator.geolocation) {
            reject("Geolocation is not supported by your browser");
        }
        navigator.geolocation.getCurrentPosition(
            (position) => {
                resolve([position.coords.latitude, position.coords.longitude]);
            },
            (error) => {
                reject(error.message);
            }
        );
    });
}
