import { Position } from "./Position";

export type Shop = {
    shopName: string;
    shopImageURL: string;
    shopCloseTime: string;
    starQuantity: number;
    longitude: number;
    latitude: number;
};

export type Shops = Shop[];

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
                    longitude: 137.1121831,
                    latitude: 35.1830169,
                },
                {
                    shopName: "ゴリラやさん1",
                    shopImageURL:
                        "https://2.bp.blogspot.com/-ruMSXp-w-qk/XDXbUFVC3FI/AAAAAAABQ-8/QRyKKr--u9E1-Rvy2SQqt0QPWeq1ME6wgCLcBGAs/s800/animal_gorilla.png",
                    shopCloseTime: "22:00",
                    starQuantity: 4,
                    longitude: 137.1121831,
                    latitude: 35.1830169,
                },
                {
                    shopName: "ゴリラやさん2",
                    shopImageURL:
                        "https://2.bp.blogspot.com/-ruMSXp-w-qk/XDXbUFVC3FI/AAAAAAABQ-8/QRyKKr--u9E1-Rvy2SQqt0QPWeq1ME6wgCLcBGAs/s800/animal_gorilla.png",
                    shopCloseTime: "22:00",
                    starQuantity: 3,
                    longitude: 137.1121831,
                    latitude: 35.1830169,
                },
                {
                    shopName: "ゴリラやさん3",
                    shopImageURL:
                        "https://2.bp.blogspot.com/-ruMSXp-w-qk/XDXbUFVC3FI/AAAAAAABQ-8/QRyKKr--u9E1-Rvy2SQqt0QPWeq1ME6wgCLcBGAs/s800/animal_gorilla.png",
                    shopCloseTime: "22:00",
                    starQuantity: 2,
                    longitude: 137.1121831,
                    latitude: 35.1830169,
                },
                {
                    shopName: "ゴリラやさん4",
                    shopImageURL:
                        "https://2.bp.blogspot.com/-ruMSXp-w-qk/XDXbUFVC3FI/AAAAAAABQ-8/QRyKKr--u9E1-Rvy2SQqt0QPWeq1ME6wgCLcBGAs/s800/animal_gorilla.png",
                    shopCloseTime: "22:00",
                    starQuantity: 1,
                    longitude: 137.1121831,
                    latitude: 35.1830169,
                },
                {
                    shopName: "ゴリラやさん5",
                    shopImageURL:
                        "https://2.bp.blogspot.com/-ruMSXp-w-qk/XDXbUFVC3FI/AAAAAAABQ-8/QRyKKr--u9E1-Rvy2SQqt0QPWeq1ME6wgCLcBGAs/s800/animal_gorilla.png",
                    shopCloseTime: "22:00",
                    starQuantity: 0,
                    longitude: 137.1121831,
                    latitude: 35.1830169,
                },
            ]);
        }, 1000);
    });
}