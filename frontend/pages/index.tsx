import Head from "next/head";
import React, { useState, useEffect } from "react";

import { Shops, FetchShops } from "../model/Shops";
import ShopCardList from "../components/ShopCardList";

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
        FetchShops(longitude, latitude).then(
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
