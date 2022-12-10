import { FC } from "react";

import L from "leaflet";
import { MapContainer, Marker, Popup, TileLayer } from "react-leaflet";
import "leaflet/dist/leaflet.css";
import iconUrl from "public/pin.png";
import shadowUrl from "leaflet/dist/images/marker-shadow.png";

import { Pin } from "../model/Pin";

type MapProps = {
    center: [number, number];
    zoom: number;
    style: React.CSSProperties;
    shopPins: Pin[];
};

L.Icon.Default.mergeOptions({
    iconRetinaUrl: iconUrl.src,
    iconUrl: iconUrl.src,
    iconSize: [30, 30],
    shadowUrl: shadowUrl.src,
});

const Map: FC<MapProps> = (props: MapProps) => {
    console.log("Map render start");

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
                        key={
                            shopPin.description +
                            shopPin.position[0] +
                            shopPin.position[1]
                        }
                    >
                        <Popup>
                            <h2>{shopPin.description}</h2>
                            <br></br>
                            <img
                                src={shopPin.photoURL}
                                style={{ height: "10em" }}
                            ></img>
                        </Popup>
                    </Marker>
                );
            })}
        </MapContainer>
    );
};

export default Map;
