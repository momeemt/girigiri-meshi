import React from "react";
import { NextComponentType, NextPageContext } from "next";

import Grid from "@mui/material/Grid";

import { Shops, Shop } from "../model/Shops";
import ShopCard from "./ShopCard";

type ShopCardListProps = {
    shopList: Shops;
};

const _ShopCardList: NextComponentType<
    NextPageContext,
    Record<string, unknown>,
    ShopCardListProps
> = (props: ShopCardListProps) => {
    console.log("ShopCardList render start");

    return (
        <Grid container spacing={1}>
            {props.shopList.map((shop: Shop) => {
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
