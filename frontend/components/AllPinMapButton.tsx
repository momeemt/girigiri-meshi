import React, { useState } from "react";
import { NextComponentType } from "next";

import { useRecoilValue } from "recoil";
import IconButton from "@mui/material/IconButton";
import MapIcon from "@mui/icons-material/Map";

import MapModal from "./MapModal";
import { ShopsAtom } from "../model/Shops";

const AllPinMapButtonStyle = {
    position: "fixed",
    right: "1.5em",
    bottom: "1.5em",
    backgroundColor: "#006699",
    color: "white",
    "&:hover": {
        backgroundColor: "#006699",
        color: "white",
        opacity: 0.8,
    },
};

const _AllPinMapButton: NextComponentType = () => {
    console.log("AllPinMapButton render start");

    const [isMapOpen, setIsMapOpen] = useState(false);
    const handleMapOpen = () => setIsMapOpen(true);
    const handleMapClose = () => setIsMapOpen(false);

    const shops = useRecoilValue(ShopsAtom);
    const shopPins = shops.map((shop) => {
        return {
            description: shop.shopName,
            position: shop.position,
        };
    });

    return (
        <>
            <IconButton
                sx={AllPinMapButtonStyle}
                onClick={handleMapOpen}
                size="large"
            >
                <MapIcon />
            </IconButton>

            <MapModal
                isMapOpen={isMapOpen}
                onClose={handleMapClose}
                shopPins={shopPins}
            ></MapModal>
        </>
    );
};

const AllPinMapButton = React.memo(_AllPinMapButton);
export default AllPinMapButton;
