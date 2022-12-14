import dynamic from "next/dynamic";
import React, { useMemo } from "react";
import { NextComponentType, NextPageContext } from "next";

import Card from "@mui/material/Card";
import Modal from "@mui/material/Modal";

import { Position } from "../model/Position";
import { Pin } from "../model/Pin";

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

export interface MapModalPinProps {
    shopPins: Pin[];
}

interface MapContorolProps {
    center: Position;
    isMapOpen: boolean;
    onClose: () => void;
}

type MapModalProps = MapContorolProps & MapModalPinProps;

const _MapModal: NextComponentType<
    NextPageContext,
    Record<string, unknown>,
    MapModalProps
> = (props: MapModalProps) => {
    console.log("MapModal render start");

    const zoom = 17;

    const Map = useMemo(
        () =>
            dynamic(() => import("../components/map"), {
                loading: () => <p>A map is loading</p>,
                ssr: false,
            }),
        []
    );

    return (
        <Modal open={props.isMapOpen} onClose={props.onClose}>
            <Card sx={mapModalStyle}>
                <Map
                    style={{ height: "100%", width: "100%" }}
                    center={props.center}
                    zoom={zoom}
                    shopPins={props.shopPins}
                ></Map>
            </Card>
        </Modal>
    );
};

const MapModal = React.memo(_MapModal);
export default MapModal;
