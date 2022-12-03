import React, { useState, useEffect } from "react";
import { NextComponentType } from "next";

import Grid from "@mui/material/Grid";

import ShopCard from "./ShopCard";

export type Shop = {
    shopName: string;
    shopImageURL: string;
    shopCloseTime: string;
    starQuantity: number;
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
                },
                {
                    shopName: "ゴリラやさん1",
                    shopImageURL:
                        "https://2.bp.blogspot.com/-ruMSXp-w-qk/XDXbUFVC3FI/AAAAAAABQ-8/QRyKKr--u9E1-Rvy2SQqt0QPWeq1ME6wgCLcBGAs/s800/animal_gorilla.png",
                    shopCloseTime: "22:00",
                    starQuantity: 4,
                },
                {
                    shopName: "ゴリラやさん2",
                    shopImageURL:
                        "https://2.bp.blogspot.com/-ruMSXp-w-qk/XDXbUFVC3FI/AAAAAAABQ-8/QRyKKr--u9E1-Rvy2SQqt0QPWeq1ME6wgCLcBGAs/s800/animal_gorilla.png",
                    shopCloseTime: "22:00",
                    starQuantity: 3,
                },
                {
                    shopName: "ゴリラやさん3",
                    shopImageURL:
                        "https://2.bp.blogspot.com/-ruMSXp-w-qk/XDXbUFVC3FI/AAAAAAABQ-8/QRyKKr--u9E1-Rvy2SQqt0QPWeq1ME6wgCLcBGAs/s800/animal_gorilla.png",
                    shopCloseTime: "22:00",
                    starQuantity: 2,
                },
                {
                    shopName: "ゴリラやさん4",
                    shopImageURL:
                        "https://2.bp.blogspot.com/-ruMSXp-w-qk/XDXbUFVC3FI/AAAAAAABQ-8/QRyKKr--u9E1-Rvy2SQqt0QPWeq1ME6wgCLcBGAs/s800/animal_gorilla.png",
                    shopCloseTime: "22:00",
                    starQuantity: 1,
                },
                {
                    shopName: "ゴリラやさん5",
                    shopImageURL:
                        "https://2.bp.blogspot.com/-ruMSXp-w-qk/XDXbUFVC3FI/AAAAAAABQ-8/QRyKKr--u9E1-Rvy2SQqt0QPWeq1ME6wgCLcBGAs/s800/animal_gorilla.png",
                    shopCloseTime: "22:00",
                    starQuantity: 0,
                },
            ]);
        }, 1000);
    });
};

const _ShopCardList: NextComponentType = () => {
    console.log("ShopCardList render start");

    const [shopList, setShopList] = useState<Shops>([]);
    const [isLoaded, setIsLoaded] = useState<boolean>(false);
    const [error, setError] = useState<string>("");

    useEffect(() => {
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
                setShopList(result);
                setIsLoaded(true);
            }
            // (error) => {
            //     setIsLoaded(true);
            //     setError(error);
            // }
        );
    }, []);

    if (!isLoaded) {
        return <div>Loading...</div>;
    }

    if (error) {
        return <div>Error: {error}</div>;
    }

    return (
        <Grid container spacing={1}>
            {shopList.map((shop: Shop) => {
                return (
                    <Grid
                        item
                        xs={12}
                        sm={6}
                        md={4}
                        lg={3}
                        key={shop.shopName + shop.shopImageURL}
                    >
                        <ShopCard shop={shop} />
                    </Grid>
                );
            })}
        </Grid>
    );
};

const ShopCardList = React.memo(_ShopCardList);
export default ShopCardList;