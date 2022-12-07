import Head from "next/head";
import React, { useState, useEffect } from "react";

import ShopCardList from "../components/ShopCardList";

export type Shop = {
    shopName: string;
    shopImageURL: string;
    shopCloseTime: string;
    starQuantity: number;
    longitude: number;
    latitude: number;
};

export type Shops = Shop[];

const dummyShopListFetch = (
    longitude: number,
    latitude: number
): Promise<Shops> => {
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
};

export default function Home() {
    console.log("index render start");

    const [shopList, setShopList] = useState<Shops>([]);
    const [isLoaded, setIsLoaded] = useState<boolean>(false);
    const [error, setError] = useState<string>("");

    useEffect(() => {
        console.log("get longitude and latitude");
        // get longitude and latitude
        let longitude = 0;
        let latitude = 0;
        if (navigator.geolocation) {
            navigator.geolocation.getCurrentPosition(
                (position) => {
                    longitude = position.coords.longitude;
                    latitude = position.coords.latitude;
                },
                (error) => {
                    setError(error.message);
                }
            );
        } else {
            setError("Geolocation is not supported by this browser.");
        }

        // fetch shop list
        // TODO: 本番用の関数に差し替え
        dummyShopListFetch(longitude, latitude).then(
            (result) => {
                console.log("promise resolved");
                setShopList(result);
                setIsLoaded(true);
            }
            // (error) => {
            //     setIsLoaded(true);
            //     setError(error);
            // }
        );
        console.log("get longitude and latitude end");
    }, []);

    if (!isLoaded) {
        return <div>Loading...</div>;
    }

    if (error) {
        return <div>Error: {error}</div>;
    }

    return (
        <div>
            <Head>
                <title>Create Next App</title>
                <meta
                    name="description"
                    content="Generated by create next app"
                />
                <link rel="icon" href="/favicon.ico" />
            </Head>
            <ShopCardList shopList={shopList}></ShopCardList>
        </div>
    );
}
