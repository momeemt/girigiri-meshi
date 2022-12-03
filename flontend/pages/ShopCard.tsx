import React from "react";
import { NextComponentType, NextPageContext } from "next";

import type { Shop } from "./ShopCardList";

type ShopCardProps = {
    shop: Shop;
};

const _ShopCard: NextComponentType<
    NextPageContext,
    Record<string, unknown>,
    ShopCardProps
> = (props: ShopCardProps) => {
    console.log("ShopCard render start");
    return <div>{props.shop.shopName}</div>;
};

const ShopCard = React.memo(_ShopCard);
export default ShopCard;
