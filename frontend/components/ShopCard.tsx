import React from "react";
import { NextComponentType, NextPageContext } from "next";

import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import CardMedia from "@mui/material/CardMedia";
import Typography from "@mui/material/Typography";
import StarIcon from "@mui/icons-material/Star";
import StarBorderIcon from "@mui/icons-material/StarBorder";

import type { Shop } from "../model/Shops";
import CardMapButton from "./CardMapButton";

type ShopCardProps = {
    shop: Shop;
};

const _ShopCard: NextComponentType<
    NextPageContext,
    Record<string, unknown>,
    ShopCardProps
> = (props: ShopCardProps) => {
    console.log("ShopCard render start");

    const MAX_STAR_QUANTITY = 5;
    const { shop } = props;

    let ratingDom;
    if (shop.rating === undefined) {
        ratingDom = <>評価なし</>;
    } else {
        const rating = Math.round(shop.rating);
        ratingDom = (
            <>
                {[...Array(rating)].map((_, index) => {
                    return (
                        <StarIcon style={{ color: "#ff6666" }} key={index} />
                    );
                })}
                {[...Array(MAX_STAR_QUANTITY - rating)].map((_, index) => {
                    return (
                        <StarBorderIcon
                            style={{ color: "#ff6666" }}
                            key={index}
                        />
                    );
                })}
            </>
        );
    }

    const closeDate = new Date(shop.closeTime);
    const minute = closeDate.getMinutes();
    const minuteStr = minute < 10 ? "0" + minute : minute;
    const hour = closeDate.getHours();
    const shopCloseTimeText = hour + ":" + minuteStr + "まで";

    return (
        <Card variant="outlined">
            <CardMedia
                component="img"
                height="200"
                image={shop.photoUrl}
                alt={shop.name}
            />
            <CardContent>
                <Typography gutterBottom variant="h5" component="div">
                    {shop.name}
                </Typography>
                <Typography variant="body1" color="text.secondary">
                    {shopCloseTimeText}
                </Typography>
                <div style={{ position: "relative", bottom: "-0.5em" }}>
                    {ratingDom}
                </div>
            </CardContent>
            <CardActions>
                <CardMapButton
                    shopPins={[
                        {
                            description: shop.name,
                            position: shop.location,
                            photoURL: shop.photoUrl,
                        },
                    ]}
                />
            </CardActions>
        </Card>
    );
};

const ShopCard = React.memo(_ShopCard);
export default ShopCard;
