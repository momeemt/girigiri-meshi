import { FC } from "react";
import { useRecoilValue } from "recoil";

import L from "leaflet";
import { MapContainer, Marker, Popup, TileLayer } from "react-leaflet";
import "leaflet/dist/leaflet.css";
import originalIconUrl from "public/pin.png";
import currentUserIconUrl from "public/running.png";
import shadowUrl from "leaflet/dist/images/marker-shadow.png";

import { useWindowSize } from "../hooks/useWindowsSize";
import { Pin } from "../model/Pin";
import { PositionAtom } from "../model/Position";

type MapProps = {
    center: [number, number];
    zoom: number;
    style: React.CSSProperties;
    shopPins: Pin[];
};

const originalIconExtended = L.Icon.extend({
    options: {
        iconRetinaUrl: originalIconUrl.src,
        iconUrl: originalIconUrl.src,
        iconSize: [30, 30],
        iconAnchor: [13, 30],
        shadowUrl: shadowUrl.src
    }
});

const currentUserIconExtended = L.Icon.extend({
    options: {
        iconRetinaUrl: currentUserIconUrl.src,
        iconUrl: currentUserIconUrl.src,
        iconSize: [40, 75],
        iconAnchor: [13, 30],
    }
});

const Map: FC<MapProps> = (props: MapProps) => {
    console.log("Map render start");
    const [width] = useWindowSize();
    const originalIcon = new originalIconExtended();
    const currentUserIcon = new currentUserIconExtended();
    const userPosition = useRecoilValue(PositionAtom);

    return (
        <MapContainer
            center={props.center}
            zoom={props.zoom}
            style={props.style}
        >
            <TileLayer
                attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
            />
            {props.shopPins.map((shopPin) => {
                return (
                    <Marker
                        position={shopPin.position}
                        icon={originalIcon}
                        key={
                            shopPin.description +
                            shopPin.position[0] +
                            shopPin.position[1]
                        }
                    >
                        <Popup maxWidth={width * 0.5}>
                            <h2>{shopPin.description}</h2>
                            <br></br>
                            <img
                                src={shopPin.photoURL}
                                style={{ height: "8em", width: "100%" }}
                            ></img>
                        </Popup>
                    </Marker>
                );
            })}
            <Marker icon={currentUserIcon} position={userPosition} key="current-user-position" />
        </MapContainer>
    );
};

export default Map;
