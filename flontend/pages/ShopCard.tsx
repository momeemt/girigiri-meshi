import React from "react";
import { NextComponentType } from "next";

const _ShopCard: NextComponentType = () => {
    console.log("ShopCardList render start");
    return <div></div>;
};

const ShopCard = React.memo(_ShopCard);
export default ShopCard;
