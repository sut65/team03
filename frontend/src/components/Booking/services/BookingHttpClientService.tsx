import { BookingsInterface } from "../../interfaces/IBooking";

const apiUrl = "http://localhost:8080";
/* -----------------------------------------------------------------------Booking--------------------------------------------------------------*/

//List Booking
async function GetBookings() {
    const requestOptions = {
        method: "GET",
        //   headers: {
        //     Authorization: `Bearer ${localStorage.getItem("token")}`,
        //     "Content-Type": "application/json",
        //   },
    };

    let res = await fetch(`${apiUrl}/bookings`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                return res.data;
            } else {
                return false;
            }
        });

    return res;
}

//***List Booking by user ID***
async function GetBookingsBYUID() {
    let uid = localStorage.getItem('uid');
    const requestOptions = {
        method: "GET",
        // headers: {
        //     Authorization: `Bearer ${localStorage.getItem("token")}`,
        //     "Content-Type": "application/json",
        // },
    };

    let res = await fetch(`${apiUrl}/bookings/user/${uid}`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                return res.data;
            } else {
                return false;
            }
        });

    return res;
}

//Craete Bookings
async function Bookings(data: BookingsInterface) {
    const requestOptions = {
        method: "POST",
        // headers: {
        //     Authorization: `Bearer ${localStorage.getItem("token")}`,
        //     "Content-Type": "application/json",
        // },
        body: JSON.stringify(data),
    };

    let res = await fetch(`${apiUrl}/bookings`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                return res.data;
            } else {
                return false;
            }
        });

    return res;
}

//GET List Branchs
async function GetBranchs() {
    const requestOptions = {
        method: "GET",
        //   headers: {
        //     Authorization: `Bearer ${localStorage.getItem("token")}`,
        //     "Content-Type": "application/json",
        //   },
    };

    let res = await fetch(`${apiUrl}/branchs`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                return res.data;
            } else {
                return false;
            }
        });

    return res;
}

// Delete Booking
async function DeleteBooking(data: BookingsInterface) {
    let booking_id = data.ID;
    const requestOptions = {
        method: "POST",
        //   headers: {
        //     Authorization: `Bearer ${localStorage.getItem("token")}`,
        //     "Content-Type": "application/json",
        //   },
        body: JSON.stringify(data),
    }
    
    let res = await fetch(`${apiUrl}/bookings/${booking_id}`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                return res.data;
            } else {
                return false;
            }
        });

    return res;
}

// Update Booking
async function UppdateBooking(data: BookingsInterface) {
    const requestOptions = {
        method: "POST",
        //   headers: {
        //     Authorization: `Bearer ${localStorage.getItem("token")}`,
        //     "Content-Type": "application/json",
        //   },
        body: JSON.stringify(data),
    }

    let res = await fetch(`${apiUrl}/bookings`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                return res.data;
            } else {
                return false;
            }
        });

    return res;
}

export {
    Bookings,
    GetBookings,
    DeleteBooking,
    GetBookingsBYUID, //****Special get */
    GetBranchs,
};