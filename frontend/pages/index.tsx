import Head from "next/head";
import Image from "next/image"
import React, { useState, useEffect } from "react";

import { useSetRecoilState } from "recoil";

import Grid from "@mui/material/Grid";

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
            <Grid container alignItems="center" justifyContent="center" className="fadeIn" style={{ width: '100%' }}>
                <Grid item style={{ maxWidth: '400px', width: '80%', maxHeight: '400px', height: '60dvw' }}>
                    <div style={{ position: 'relative', width: '100%', height: '100%', textAlign: 'center' }}>
                        <Image src="/logo.png" alt="logo" layout="fill" objectFit="contain" />
                    </div>
                </Grid>
            </Grid>

            <div style={{ padding: '0 10%', backgroundColor: '#F6D60F' }}>
                <ShopCardList></ShopCardList>
            </div>
            <AllPinMapButton></AllPinMapButton>
        </div>
    );
}
