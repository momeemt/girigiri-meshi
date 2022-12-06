import dynamic from "next/dynamic";
import React, { useMemo } from "react";
import { NextComponentType, NextPageContext } from "next";

import Grid from "@mui/material/Grid";
import IconButton from "@mui/material/IconButton";
import FmdGoodIcon from "@mui/icons-material/FmdGood";
import Typography from "@mui/material/Typography";
import Card from "@mui/material/Card";
import Modal from "@mui/material/Modal";

type MapModalProps = {
    longitude: number;
    latitude: number;
    description: string;
};

const mapModalStyle = {
    position: "absolute",
    top: "50%",
    left: "50%",
    transform: "translate(-50%, -50%)",
    width: "70%",
    height: "70%",
    bgcolor: "background.paper",
    boxShadow: 24,
    p: 4,
};

const _MapModal: NextComponentType<
    NextPageContext,
    Record<string, unknown>,
    MapModalProps
> = (props: MapModalProps) => {
    console.log("MapModal render start");

    const [isMapOpen, setIsMapOpen] = React.useState(false);
    const handleMapOpen = () => setIsMapOpen(true);
    const handleMapClose = () => setIsMapOpen(false);

    const position: [number, number] = [props.latitude, props.longitude];
    const zoom = 14;

    const Map = useMemo(
        () =>
            dynamic(() => import("../components/map"), {
                loading: () => <p>A map is loading</p>,
                ssr: false,
            }),
        []
    );

    return (
        <Grid container justifyContent="center">
            <IconButton style={{ color: "#006699" }} onClick={handleMapOpen}>
                <FmdGoodIcon />
                <Typography variant="body1">地図を表示</Typography>
            </IconButton>
            <Modal open={isMapOpen} onClose={handleMapClose}>
                <Card sx={mapModalStyle}>
                    <Map
                        style={{ height: "100%", width: "100%" }}
                        center={position}
                        zoom={zoom}
                        description={props.description}
                    ></Map>
                </Card>
            </Modal>
        </Grid>
    );
};

const MapModal = React.memo(_MapModal);
export default MapModal;