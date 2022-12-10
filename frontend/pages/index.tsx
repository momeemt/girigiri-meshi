import Head from "next/head";
import React, { useState, useEffect } from "react";

import { useSetRecoilState } from "recoil";

import { PositionAtom, GetUserPosition } from "../model/Position";
import { ShopsAtom, FetchShops } from "../model/Shops";
import ShopCardList from "../components/ShopCardList";
import AllPinMapButton from "../components/AllPinMapButton";
import Loading from "../components/Loading";

export default function Home() {
    console.log("index render start");

    const setUserPosition = useSetRecoilState(PositionAtom);
    const setShops = useSetRecoilState(ShopsAtom);
    const [isLoaded, setIsLoaded] = useState<boolean>(false);
    const [error, setError] = useState<string>("");

    useEffect(() => {
        (async () => {
            const nowUserPosition = await GetUserPosition().catch((error) => {
                setError(error);
                setIsLoaded(true);
                return null;
            });
            if (nowUserPosition === null) {
                return;
            }

            const shops = await FetchShops(nowUserPosition).catch((error) => {
                setError(error);
                setIsLoaded(true);
                return null;
            });
            if (shops === null) {
                return;
            }

            setUserPosition(nowUserPosition);
            setShops(shops);
            setIsLoaded(true);
        })();
    }, []);

    if (!isLoaded) {
        return <Loading />;
    }

    if (error) {
        return <div>Error: {error}</div>;
    }

    return (
        <div className="fadeIn">
            <Head>
                <title>GiriGiriMeshi</title>
                <meta
                    name="description"
                    content="ギリギリで駆け込める飯屋を探してギリギリで駆け込もう"
                />
                <link rel="icon" href="/favicon.ico" />
            </Head>
            <ShopCardList></ShopCardList>
            <AllPinMapButton></AllPinMapButton>
        </div>
    );
}
