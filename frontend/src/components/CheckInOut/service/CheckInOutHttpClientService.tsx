import { CheckInOutInterface } from "../../../models/ICheckInOut";

const apiUrl = "http://localhost:8080";
	// //========================= checkInOut routes
	// //status
	// r.GET("/checkinoutstatus/:id", check.GetCheckInOutStatus)
	// r.GET("/checkinoutstatuses", check.ListCheckInOutStatuses)
	// r.POST("/checkinoutstatus", check.CreateCheckInOutStatus)
	// r.PATCH("/checkinoutstatus", check.UpdateCheckInOutStatus)
	// r.DELETE("/checkinoutstatus/:id", check.DeleteCheckInOutStatus)
	// //main
	// r.GET("/checkinout/:id", check.GetCheckInOut)
	// r.GET("/checkinouts", check.ListCheckInOuts)
	// r.POST("/checkinout", check.CreateCheckInOut)
	// r.PATCH("/checkinout", check.UpdateCheckInOut)
	// r.DELETE("/checkinout/:id", check.DeleteCheckInOut)

// List CheckInOut
async function GetCheckInOut() {
    const requestOptions = {
        method: "GET",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
    };

    let res = await fetch(`${apiUrl}/checkinouts`, requestOptions)
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

// List CheckInOut
async function GetBookingByDate() {
    const requestOptions = {
        method: "GET",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
    };

    let res = await fetch(`${apiUrl}/bookingbydate`, requestOptions)
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

// List Status
async function GetIOStatus() {
    const requestOptions = {
        method: "GET",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
    };

    let res = await fetch(`${apiUrl}/checkinoutstatuses`, requestOptions)
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

//Craete CheckInout
async function CreateCheckInOut(data: CheckInOutInterface) {
    const requestOptions = {
        method: "POST",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    };

    let res = await fetch(`${apiUrl}/checkinout`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                return { status: true, message: res.data };
            } else {
                console.log(res.error);
                return { status: false, message: res.error };
            }
        });

    return res;
}

// Delete CheckInOut
async function DeleteCheckInOut(data: number) {
    let checkInOutID = data;
    const requestOptions = {
        method: "DELETE",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
        body: JSON.stringify(data),
    }
    
    let res = await fetch(`${apiUrl}/checkinout/${checkInOutID}`, requestOptions)
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

// Update CheckInOut
async function UpdateCheckInOut(data: CheckInOutInterface) {
    const requestOptions = {
        method: "PATCH",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
        body: JSON.stringify(data),
    }

    let res = await fetch(`${apiUrl}/checkinout`, requestOptions)
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

//protected.PATCH("/checkinout/:id", check.CheckOut)
async function CheckOut(data: number) {
    let checkInOutID = data;
    const requestOptions = {
        method: "PATCH",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
        body: JSON.stringify(data),
    }

    let res = await fetch(`${apiUrl}/checkinout/${checkInOutID}`, requestOptions)
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

//update checkIn
async function UpdateCheckIn(data: CheckInOutInterface) {
    const requestOptions = {
        method: "PATCH",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
        body: JSON.stringify(data),
    }

    let res = await fetch(`${apiUrl}/checkin`, requestOptions)
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

//update checkOut
async function UpdateCheckOut(data: CheckInOutInterface) {
    const requestOptions = {
        method: "PATCH",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
        body: JSON.stringify(data),
    }

    let res = await fetch(`${apiUrl}/checkout`, requestOptions)
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

// List Status
async function GetEmps() {
    const requestOptions = {
        method: "GET",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
    };

    let res = await fetch(`${apiUrl}/Employees`, requestOptions)
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
    GetCheckInOut,
    GetIOStatus,
    CreateCheckInOut,
    DeleteCheckInOut,
    UpdateCheckInOut,
    UpdateCheckIn,
    UpdateCheckOut,
    GetEmps,
    CheckOut,
    GetBookingByDate,
};

