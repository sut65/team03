import { CHK_PaymentsInterface } from "../../../models/ICHK_Payment";


const apiUrl = "http://localhost:8080";
// //=================================================== Check Payment Routes
// r.GET("/chk_payments", chk_payment.ListCHK_Payments)
// r.GET("/chk_payment/:id", chk_payment.GetCHK_Payment)
// r.POST("/chk_payments", chk_payment.CreateCHK_Payment)
// r.PATCH("/chk_payments", chk_payment.UpdateCHK_Payment)
// r.DELETE("/chk_payments/:id", chk_payment.DeleteCHK_Payment)
// // ---Status---
// r.GET("/chk_payment/statuses", chk_payment.ListCHK_Payments)
// r.GET("/chk_payment/status/:id", chk_payment.GetCHK_Payment)
// r.POST("/chk_payment/statuses", chk_payment.CreateCHK_Payment)
// r.PATCH("/chk_payment/statuses", chk_payment.UpdateCHK_Payment)
// r.DELETE("/chk_payment/statuses/:id", chk_payment.DeleteCHK_Payment)
// //=================================================== Check Payment Routes
// List CHK_Payment
async function GetCHK_Payments() {
    const requestOptions = {
        method: "GET",
        //   headers: {
        //     Authorization: `Bearer ${localStorage.getItem("token")}`,
        //     "Content-Type": "application/json",
        //   },
    };

    let res = await fetch(`${apiUrl}/chk_payments`, requestOptions)
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
        // headers: {
        //     Authorization: `Bearer ${localStorage.getItem("token")}`,
        //     "Content-Type": "application/json",
        // },
        body: JSON.stringify(data),
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

async function GetStatuses() {
    const requestOptions = {
        method: "GET",
        //   headers: {
        //     Authorization: `Bearer ${localStorage.getItem("token")}`,
        //     "Content-Type": "application/json",
        //   },
    };

    let res = await fetch(`${apiUrl}/statuses`, requestOptions)
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

async function GetPayments() {
    const requestOptions = {
        method: "GET",
        //   headers: {
        //     Authorization: `Bearer ${localStorage.getItem("token")}`,
        //     "Content-Type": "application/json",
        //   },
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
    GetCHK_Payments,
    GetStatuses,
    GetPayments,
    CHK_Payments,
};

