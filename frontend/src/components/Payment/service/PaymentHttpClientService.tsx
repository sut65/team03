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

async function GetTotalPriceByCID(id?: string | null) {
    let res = await fetch(`${apiUrl}/totalprice/customer/${id}`, requestOptionsGet)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                return res.data.TotalAmount;
            } else {
                return false;
            }
        });
    return res;
}

async function GetPriceFood(id?: string) {
    let res = await fetch(`${apiUrl}/pricefood/customer/${id}`, requestOptionsGet)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                return res.data.Price;
            } else {
                return false;
            }
        });
    return res;
}
async function GetFoodItem(id?: string) {
    let res = await fetch(`${apiUrl}/fooditem/customer/${id}`, requestOptionsGet)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                return res.data.FoodItem;
            } else {
                return false;
            }
        });
    return res;
}
async function GetPriceDrink(id?: string) {
    let res = await fetch(`${apiUrl}/pricedrink/customer/${id}`, requestOptionsGet)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                return res.data.Price;
            } else {
                return false;
            }
        });
    return res;
}
async function GetDrinkItem(id?: string) {
    let res = await fetch(`${apiUrl}/drinkitem/customer/${id}`, requestOptionsGet)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                return res.data.DrinkItem;
            } else {
                return false;
            }
        });
    return res;
}
async function GetPriceAccessorie(id?: string) {
    let res = await fetch(`${apiUrl}/priceaccessorie/customer/${id}`, requestOptionsGet)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                return res.data.Price;
            } else {
                return false;
            }
        });
    return res;
}
async function GetAccessorieItem(id?: string) {
    let res = await fetch(`${apiUrl}/accessorieitem/customer/${id}`, requestOptionsGet)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                return res.data.StorageItem;
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
    GetTotalPriceByCID,
    GetPriceFood,
    GetFoodItem,
    GetPriceDrink,
    GetDrinkItem,
    GetPriceAccessorie,
    GetAccessorieItem,
};