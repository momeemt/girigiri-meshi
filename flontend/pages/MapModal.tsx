import dynamic from "next/dynamic";
import React, { useMemo } from "react";
import { NextComponentType, NextPageContext } from "next";

import Card from "@mui/material/Card";
import Modal from "@mui/material/Modal";

type MapModalProps = {
    handleMapClose: () => void;
    isMapOpen: boolean;
};

const mapModalStyle = {
    position: "absolute",
    top: "50%",
    left: "50%",
    transform: "translate(-50%, -50%)",
    width: 800,
    height: 800,
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

    const position: [number, number] = [51.505, -0.09];
    const zoom = 13;

    const Map = useMemo(
        () =>
            dynamic(() => import("./map"), {
                loading: () => <p>A map is loading</p>,
                ssr: false,
            }),
        []
    );

    return (
        <Modal
            open={props.isMapOpen}
            onClose={props.handleMapClose}
            aria-labelledby="modal-modal-title"
            aria-describedby="modal-modal-description"
        >
            <Card sx={mapModalStyle}>
                <Map
                    style={{ height: "100%", width: "100%" }}
                    center={position}
                    zoom={zoom}
                ></Map>
            </Card>
        </Modal>
    );
};

const MapModal = React.memo(_MapModal);
export default MapModal;
