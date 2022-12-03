import React from "react";
import { NextComponentType, NextPageContext } from "next";

import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import CardMedia from "@mui/material/CardMedia";
import Typography from "@mui/material/Typography";
import StarIcon from "@mui/icons-material/Star";
import StarBorderIcon from "@mui/icons-material/StarBorder";
import Grid from "@mui/material/Grid";
import IconButton from "@mui/material/IconButton";
import FmdGoodIcon from "@mui/icons-material/FmdGood";

import type { Shop } from "./ShopCardList";
import MapModal from "./MapModal";

type ShopCardProps = {
    shop: Shop;
};

const _ShopCard: NextComponentType<
    NextPageContext,
    Record<string, unknown>,
    ShopCardProps
> = (props: ShopCardProps) => {
    console.log("ShopCard render start");
    const [isMapOpen, setIsMapOpen] = React.useState(false);
    const handleMapOpen = () => setIsMapOpen(true);
    const handleMapClose = () => setIsMapOpen(false);

    const MAX_STAR_QUANTITY = 5;

    const { shop } = props;
    const shopCloseTimeText = shop.shopCloseTime + "まで";

    return (
        <Card variant="outlined">
            <CardMedia
                component="img"
                height="200"
                image={shop.shopImageURL}
                alt={shop.shopName}
            />
            <CardContent>
                <Typography gutterBottom variant="h5" component="div">
                    {shop.shopName}
                </Typography>
                <Typography variant="body1" color="text.secondary">
                    {shopCloseTimeText}
                </Typography>
                <div style={{ position: "relative", bottom: "-0.5em" }}>
                    {[...Array(shop.starQuantity)].map((_, index) => {
                        return (
                            <StarIcon
                                style={{ color: "#ff6666" }}
                                key={index}
                            />
                        );
                    })}
                    {[...Array(MAX_STAR_QUANTITY - shop.starQuantity)].map(
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
                <Grid container justifyContent="center">
                    <IconButton
                        style={{ color: "#006699" }}
                        onClick={handleMapOpen}
                    >
                        <FmdGoodIcon />
                        <Typography variant="body1">地図を表示</Typography>
                    </IconButton>
                    <MapModal
                        isMapOpen={isMapOpen}
                        handleMapClose={handleMapClose}
                        longitude={shop.longitude}
                        latitude={shop.latitude}
                    ></MapModal>
                </Grid>
            </CardActions>
        </Card>
    );
};

const ShopCard = React.memo(_ShopCard);
export default ShopCard;
