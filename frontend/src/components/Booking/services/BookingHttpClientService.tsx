import { BookingsInterface } from "../../../models/modelBooking/IBooking";
import { BranchsInterface } from "../../../models/modelBooking/IBranch";


const apiUrl = "http://localhost:8080";

async function GetCustomerByUID() {
    let uid = localStorage.getItem('id');
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
    }

    let res = await fetch(`${apiUrl}/customer/${uid}`, requestOptions)
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

async function GetCustomers() {
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
    }

    let res = await fetch(`${apiUrl}/customers`, requestOptions)
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

/* -----------------------------------------------------------------------Booking--------------------------------------------------------------*/
// //=================================================== Booking Routes
// protected.GET("/bookings", booking.ListBookings)
// protected.GET("/booking/:id", booking.GetBooking)
// protected.GET("/bookings/user/:id", booking.ListBookingsByUID)
// protected.POST("/bookings", booking.CreateBooking)
// protected.PATCH("/bookings/:id", booking.UpdateBooking)
// protected.DELETE("/bookings/:id", booking.DeleteBooking)
// protected.GET("/bookingbydate", booking.ListBookingsBydate)
// protected.GET("/bookingtotalgroupbydate", booking.ListBookingsTotalbyCID)

//List Booking
async function GetBookings() {
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
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

//Get Booking
async function GetBooking(data: string | undefined) {
    let b_id = data;
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
    };

    let res = await fetch(`${apiUrl}/booking/${b_id}`, requestOptions)
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
    let uid = localStorage.getItem('id');
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
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

//***List Booking Total by customerID**
async function GetBookingsSumTotal() {
    let uid = localStorage.getItem('id');
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
    };

    let res = await fetch(`${apiUrl}/bookingtotalgroupbydate`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                return res.data;
                console.log(res.data);
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
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    };

    let res = await fetch(`${apiUrl}/bookings`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                return { status: true, message: res.data };
            } else {
                return { status: false, message: res.error };
            }
        });

    return res;
}

// Delete Booking
async function DeleteBooking(data: BookingsInterface) {
    let booking_id = data.ID;
    const requestOptions = {
        method: "DELETE",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    }

    let res = await fetch(`${apiUrl}/bookings/${booking_id}`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                return { status: true, message: res.data };
            } else {
                return { status: false, message: res.error };
            }
        });

    return res;
}

// Update Booking
async function UppdateBooking(data: BookingsInterface) {
    let b_id = data.ID;
    const requestOptions = {
        method: "PATCH",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    }

    let res = await fetch(`${apiUrl}/bookings/${b_id}`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                return { status: true, message: res.data };
            } else {
                return { status: false, message: res.error };
            }
        });

    return res;
}

// // ---Branch---
//	r.GET("/branchs", booking.ListBranchs)
	// r.GET("/branch/:id", booking.GetBranch)
	// r.POST("/branchs", booking.CreateBranch)
	// r.PATCH("/branchs", booking.UpdateBranch)
	// r.DELETE("/branchs/:id", booking.DeleteBranch)
// //=================================================== Booking Routes

//GET List Branchs
async function GetBranchs() {
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
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

//GET Get Branchs
async function GetBranch(data: BranchsInterface) {
    let br_id = data.ID;
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
    };

    let res = await fetch(`${apiUrl}/branch/${br_id}`, requestOptions)
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

//Craete Branchs
async function Branchs(data: BranchsInterface) {
    const requestOptions = {
        method: "POST",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
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

// Delete Branch
async function DeleteBranch(data: BranchsInterface) {
    let branch_id = data.ID;
    const requestOptions = {
        method: "DELETE",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    }

    let res = await fetch(`${apiUrl}/branchs/${branch_id}`, requestOptions)
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

// Update Branch
async function UppdateBranch(data: BranchsInterface) {
    const requestOptions = {
        method: "PATCH",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    }

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

export {
    GetCustomerByUID,
    GetCustomers,

    Bookings,
    GetBookings,
    GetBooking,
    GetBookingsBYUID, //****Special get */
    GetBookingsSumTotal,
    DeleteBooking,
    UppdateBooking,

    Branchs,
    GetBranchs,
    GetBranch,
    UppdateBranch,
    DeleteBranch,
};