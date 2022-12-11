import React, { useState } from "react";
import { NextComponentType, NextPageContext } from "next";

import Grid from "@mui/material/Grid";
import IconButton from "@mui/material/IconButton";
import FmdGoodIcon from "@mui/icons-material/FmdGood";
import Typography from "@mui/material/Typography";

import MapModal from "./MapModal";
import { MapModalPinProps } from "./MapModal";

const _CardMapButton: NextComponentType<
    NextPageContext,
    Record<string, unknown>,
    MapModalPinProps
> = (props: MapModalPinProps) => {
    console.log("CardMapButton render start");

    const [isMapOpen, setIsMapOpen] = useState(false);
    const handleMapOpen = () => setIsMapOpen(true);
    const handleMapClose = () => setIsMapOpen(false);

    return (
        <Grid container justifyContent="center">
            <IconButton style={{ color: "#006699" }} onClick={handleMapOpen}>
                <FmdGoodIcon />
                <Typography variant="body1"fontFamily="Zen Maru Gothic">地図を表示</Typography>
            </IconButton>

            <MapModal
                isMapOpen={isMapOpen}
                onClose={handleMapClose}
                shopPins={props.shopPins}
                center={props.shopPins[0].position}
            ></MapModal>
        </Grid>
    );
};

const CardMapButton = React.memo(_CardMapButton);
export default CardMapButton;
