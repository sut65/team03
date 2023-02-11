import { PaymentsInterface } from "../../../models/modelPayment/IPayment";

const apiUrl = "http://localhost:8080";

const requestOptionsGet = {
    method: "GET",
    headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
    },
}

async function AddPayment(data: PaymentsInterface) {
    const requestOptions = {
        method: "POST",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    };

    let res = await fetch(`${apiUrl}/payment`, requestOptions)
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

async function UpdatePayment(data: PaymentsInterface) {
    const requestOptions = {
        method: "PATCH",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    }

    let res = await fetch(`${apiUrl}/payments`, requestOptions)
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

async function GetPayment(id: string) {
    let res = await fetch(`${apiUrl}/payment/customer/${id}`, requestOptionsGet)
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

async function GetDestination(id: string) {
    let res = await fetch(`${apiUrl}/method/${id}`, requestOptionsGet)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                return res.data.Destination;
            } else {
                return false;
            }
        });
    return res;
}

async function GetPicture(id: string) {
    let res = await fetch(`${apiUrl}/method/${id}`, requestOptionsGet)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                return res.data.Picture;
            } else {
                return false;
            }
        });
    return res;
}

async function GetPaymentMethods() {
    let res = await fetch(`${apiUrl}/paymentmethods`, requestOptionsGet)
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

async function GetPlaces() {
    let res = await fetch(`${apiUrl}/places`, requestOptionsGet)
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

async function GetMethodP(id: string) {
    let res = await fetch(`${apiUrl}/methods/paymet/${id}`, requestOptionsGet)
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

async function GetPaymentByID(id?: string) {
    let res = await fetch(`${apiUrl}/payment/${id}`, requestOptionsGet)
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
    AddPayment,
    UpdatePayment,
    GetPayment,
    GetDestination,
    GetPicture,
    GetPaymentMethods,
    GetPlaces,
    GetMethodP,
    GetPaymentByID,
};