import { CHK_PaymentsInterface } from "../../../models/modelCHK_Payment/ICHK_Payment";
import { CHK_PaymentStatusesInterface } from "../../../models/modelCHK_Payment/IStatus";


const apiUrl = "http://localhost:8080";

async function GetEmployeeByUID() {
    let uid = localStorage.getItem('id');
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
    }

    let res = await fetch(`${apiUrl}/Employee/${uid}`, requestOptions)
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

// //=================================================== Check Payment Routes
// protected.GET("/chk_payments", chk_payment.ListCHK_Payments)
// protected.GET("/chk_payment/:id", chk_payment.GetCHK_Payment)
// protected.POST("/chk_payments", chk_payment.CreateCHK_Payment)
// protected.PATCH("/chk_payments", chk_payment.UpdateCHK_Payment)
// protected.DELETE("/chk_payments/:id", chk_payment.DeleteCHK_Payment)

// List CHK_Payment
async function GetCHK_Payments() {
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
    };

    let res = await fetch(`${apiUrl}/chk_payments`, requestOptions)
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

//Get CHK_Payment
async function GetCHK_Payment(data: string | null) {
    let b_id = data;
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
    };

    let res = await fetch(`${apiUrl}/chk_payment/${b_id}`, requestOptions)
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

//Craete CHK_Payments
async function CHK_Payments(data: CHK_PaymentsInterface) {
    const requestOptions = {
        method: "POST",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    };

    let res = await fetch(`${apiUrl}/chk_payments`, requestOptions)
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

// Delete CHK_Payment
async function DeleteCHK_Payment(data: CHK_PaymentsInterface) {
    let chkp_id = data.ID;
    const requestOptions = {
        method: "DELETE",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    }

    let res = await fetch(`${apiUrl}/chk_payments/${chkp_id}`, requestOptions)
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

// Update CHK_Payment
async function UppdateCHK_Payment(data: CHK_PaymentsInterface) {
    let chkp_id = data.ID;
    const requestOptions = {
        method: "PATCH",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    }

    let res = await fetch(`${apiUrl}/chk_payments/${chkp_id}`, requestOptions)
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

// // ---Status---
// protected.GET("/chk_payment/statuses", chk_payment.ListStatuses)
// protected.GET("/chk_payment/status/:id", chk_payment.GetStatus)
// protected.POST("/chk_payment/statuses", chk_payment.CreateStatus)
// protected.PATCH("/chk_payment/statuses", chk_payment.UpdateStatus)
// protected.DELETE("/chk_payment/statuses/:id", chk_payment.DeleteStatus)

async function GetStatuses() {
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
    };

    let res = await fetch(`${apiUrl}/chk_payment/statuses`, requestOptions)
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

//Get Status
async function GetStatus(data: CHK_PaymentStatusesInterface) {
    let s_id = data.ID;
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
    };

    let res = await fetch(`${apiUrl}/chk_payment/status/${s_id}`, requestOptions)
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

//Craete Statuses
async function Statuses(data: CHK_PaymentStatusesInterface) {
    const requestOptions = {
        method: "POST",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    };

    let res = await fetch(`${apiUrl}/chk_payment/statuses`, requestOptions)
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

// Delete Status
async function DeleteStatus(data: CHK_PaymentStatusesInterface) {
    let booking_id = data.ID;
    const requestOptions = {
        method: "DELETE",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    }

    let res = await fetch(`${apiUrl}/chk_payment/statuses/${booking_id}`, requestOptions)
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

// Update Status
async function UpdateStatus(data: CHK_PaymentStatusesInterface) {
    const requestOptions = {
        method: "PATCH",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    }

    let res = await fetch(`${apiUrl}/chk_payment/statuses`, requestOptions)
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

// ===================================================================== ETC
async function GetPayments() {
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
    };

    let res = await fetch(`${apiUrl}/payments`, requestOptions)
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
    GetEmployeeByUID,

    GetCHK_Payments,
    GetCHK_Payment,
    CHK_Payments,
    UppdateCHK_Payment,
    DeleteCHK_Payment,

    GetStatuses,
    GetStatus,
    Statuses,
    UpdateStatus,
    DeleteStatus,

    GetPayments,
};

