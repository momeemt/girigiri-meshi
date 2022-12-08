import { atom } from "recoil";

import { Position } from "./Position";

export type Shop = {
    shopName: string;
    shopImageURL: string;
    shopCloseTime: string;
    starQuantity: number;
    position: Position;
};

export type Shops = Shop[];

export const ShopsAtom = atom({
    key: "shops",
    default: [] as Shops,
});

export function FetchShops(position: Position): Promise<Shops> {
    return new Promise((resolve) => {
        setTimeout(() => {
            resolve([
                {
                    shopName: "ゴリラやさん",
                    shopImageURL:
                        "https://2.bp.blogspot.com/-ruMSXp-w-qk/XDXbUFVC3FI/AAAAAAABQ-8/QRyKKr--u9E1-Rvy2SQqt0QPWeq1ME6wgCLcBGAs/s800/animal_gorilla.png",
                    shopCloseTime: "22:00",
                    starQuantity: 5,
                    position: [35.1830169, 137.1121831],
                },
                {
                    shopName: "ゴリラやさん1",
                    shopImageURL:
                        "https://2.bp.blogspot.com/-ruMSXp-w-qk/XDXbUFVC3FI/AAAAAAABQ-8/QRyKKr--u9E1-Rvy2SQqt0QPWeq1ME6wgCLcBGAs/s800/animal_gorilla.png",
                    shopCloseTime: "22:00",
                    starQuantity: 4,
                    position: [35.1810169, 137.1121831],
                },
                {
                    shopName: "ゴリラやさん2",
                    shopImageURL:
                        "https://2.bp.blogspot.com/-ruMSXp-w-qk/XDXbUFVC3FI/AAAAAAABQ-8/QRyKKr--u9E1-Rvy2SQqt0QPWeq1ME6wgCLcBGAs/s800/animal_gorilla.png",
                    shopCloseTime: "22:00",
                    starQuantity: 3,
                    position: [35.1830169, 137.1101831],
                },
                {
                    shopName: "ゴリラやさん3",
                    shopImageURL:
                        "https://2.bp.blogspot.com/-ruMSXp-w-qk/XDXbUFVC3FI/AAAAAAABQ-8/QRyKKr--u9E1-Rvy2SQqt0QPWeq1ME6wgCLcBGAs/s800/animal_gorilla.png",
                    shopCloseTime: "22:00",
                    starQuantity: 2,
                    position: [35.1810169, 137.1101831],
                },
                {
                    shopName: "ゴリラやさん4",
                    shopImageURL:
                        "https://2.bp.blogspot.com/-ruMSXp-w-qk/XDXbUFVC3FI/AAAAAAABQ-8/QRyKKr--u9E1-Rvy2SQqt0QPWeq1ME6wgCLcBGAs/s800/animal_gorilla.png",
                    shopCloseTime: "22:00",
                    starQuantity: 1,
                    position: [35.1850169, 137.1121831],
                },
                {
                    shopName: "ゴリラやさん5",
                    shopImageURL:
                        "https://2.bp.blogspot.com/-ruMSXp-w-qk/XDXbUFVC3FI/AAAAAAABQ-8/QRyKKr--u9E1-Rvy2SQqt0QPWeq1ME6wgCLcBGAs/s800/animal_gorilla.png",
                    shopCloseTime: "22:00",
                    starQuantity: 0,
                    position: [35.1830169, 137.1141831],
                },
                {
                    shopName: "ゴリラやさん6",
                    shopImageURL:
                        "https://2.bp.blogspot.com/-ruMSXp-w-qk/XDXbUFVC3FI/AAAAAAABQ-8/QRyKKr--u9E1-Rvy2SQqt0QPWeq1ME6wgCLcBGAs/s800/animal_gorilla.png",
                    shopCloseTime: "22:00",
                    starQuantity: 0,
                    position: [35.1850169, 137.1141831],
                },
            ]);
        }, 1000);
    });
}