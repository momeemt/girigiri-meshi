import React, { useEffect } from "react";
import { NextComponentType } from "next";

const _Location: NextComponentType = () => {
    console.log("location render start");

    const [error, setError] = React.useState<string>("");

    useEffect(() => {
        // get longitude and latitude
        let longitude = 0;
        let latitude = 0;
        if (navigator.geolocation) {
            navigator.geolocation.getCurrentPosition(
                (position) => {
                    longitude = position.coords.longitude;
                    latitude = position.coords.latitude;
                    console.log(longitude, latitude);
                },
                (error) => {
                    setError(error.message);
                }
            );
        } else {
            setError("Geolocation is not supported by this browser.");
        }
    }, []);

    return <div>{error}</div>;
};

const Location = React.memo(_Location);
export default Location;
