import React from "react";
import { NextComponentType } from "next";

import { useRecoilValue } from "recoil";
import Grid from "@mui/material/Grid";

import { Shop, ShopsAtom } from "../model/Shops";
import ShopCard from "./ShopCard";

const _ShopCardList: NextComponentType = () => {
    console.log("ShopCardList render start");
    const shops = useRecoilValue(ShopsAtom);

    return (
        <Grid container spacing={1}>
            {shops.map((shop: Shop) => {
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
