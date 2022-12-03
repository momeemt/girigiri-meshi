import React from "react";
import { NextComponentType } from "next";

const _ShopCardList: NextComponentType = () => {
    console.log("ShopCardList render start");
    return <div></div>;
};

const ShopCardList = React.memo(_ShopCardList);
export default ShopCardList;
