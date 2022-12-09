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
    const shopCloseTimeText = shop.closeTime + "まで";
    const rating = Math.round(shop.rating);

    console.log(shop)

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
                    {[...Array(rating)].map((_, index) => {
                        return (
                            <StarIcon
                                style={{ color: "#ff6666" }}
                                key={index}
                            />
                        );
                    })}
                    {[...Array(MAX_STAR_QUANTITY - rating)].map(
                        (_, index) => {
                            return (
                                <StarBorderIcon
                                    style={{ color: "#ff6666" }}
                                    key={index}
                                />
                            );
                        }
                    )}
                </div>
            </CardContent>
            <CardActions>
                <CardMapButton
                    shopPins={[
                        {
                            description: shop.name,
                            position: shop.location,
                        },
                    ]}
                />
            </CardActions>
        </Card>
    );
};

const ShopCard = React.memo(_ShopCard);
export default ShopCard;
