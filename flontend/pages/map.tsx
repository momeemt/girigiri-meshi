import { FC } from "react";

import L from "leaflet";
import { MapContainer, Marker, Popup, TileLayer } from "react-leaflet";
import "leaflet/dist/leaflet.css";
import iconRetinaUrl from "leaflet/dist/images/marker-icon-2x.png";
import iconUrl from "leaflet/dist/images/marker-icon.png";
import shadowUrl from "leaflet/dist/images/marker-shadow.png";

type MapProps = {
    center: [number, number];
    zoom: number;
    style: React.CSSProperties;
    description: string;
};

L.Icon.Default.mergeOptions({
    iconRetinaUrl: iconRetinaUrl.src,
    iconUrl: iconUrl.src,
    shadowUrl: shadowUrl.src,
});

const Map: FC<MapProps> = (props: MapProps) => {
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
            <Marker position={props.center}>
                <Popup>{props.description}</Popup>
            </Marker>
        </MapContainer>
    );
};

export default Map;
