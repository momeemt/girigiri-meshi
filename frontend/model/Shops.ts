import { atom } from "recoil";

import { Position } from "./Position";
import axios from "axios";

export type Shop = {
    name: string;
    photoUrl: string;
    closeTime: string;
    rating: number | undefined;
    location: Position;
};

type FetchShop = {
    name: string;
    photoUrl: string;
    closeTime: string;
    rating: number | undefined;
    location: {
        latitude: number;
        longitude: number;
    };
};

export type Shops = Shop[];

export const ShopsAtom = atom({
    key: "shops",
    default: [] as Shops,
});

export function FetchShops(position: Position): Promise<Shops> {
    return new Promise((resolve, reject) => {
        if (process.env.NEXT_PUBLIC_SHOPS_FETCH_SERVER === undefined) {
            reject("NEXT_PUBLIC_SHOPS_FETCH_SERVER is not defined");
            return;
        }

        const body = {
            latitude: position[0],
            longitude: position[1],
        };

        const header = {
            headers: {
                accept: "application/json",
                "Content-Type": "application/json",
            },
        };

        console.log(body);

        axios
            .post(process.env.NEXT_PUBLIC_SHOPS_FETCH_SERVER, body, header)
            .then((response) => {
                const res = response.data.map((shop: FetchShop) => {
                    return {
                        name: shop.name,
                        photoUrl: shop.photoUrl,
                        closeTime: shop.closeTime,
                        rating: shop.rating,
                        location: [shop.location.latitude, shop.location.longitude],
                    };
                });
                resolve(res);
            })
            .catch((error) => {
                reject(error.message);
            });
    });
}
