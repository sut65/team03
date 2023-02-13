import { CheckroomInterface } from "../../../models/ICheckroom" 
import { DamageInterface } from "../../../models/modelCheckroom/IDamage" 
import { StatusInterface } from "../../../models/modelCheckroom/IStatus" 
import { ProductInterface } from "../../../models/IStorage" 
import { EmployeeInterface } from "../../../models/IEmployee" 
import { RoomInterface } from "../../../models/IRoom";

const apiUrl = "http://localhost:8080";

async function GetRoom() {
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
    };

    let res = await fetch(`${apiUrl}/rooms`, requestOptions)
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

async function GetProduct() {
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
    };

    let res = await fetch(`${apiUrl}/products`, requestOptions)
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

async function GetDamage() {
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
    };

    let res = await fetch(`${apiUrl}/damages`, requestOptions)
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

async function GetStatus() {
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
    };

    let res = await fetch(`${apiUrl}/status`, requestOptions)
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

async function Checkrooms(data: CheckroomInterface) {
    const requestOptions = {
        method: "POST",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    };

    let res = await fetch(`${apiUrl}/checkrooms`, requestOptions)
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


async function GetChecklists() {
    const requestOptions = {
        method: "GET",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
    };

    let res = await fetch(`${apiUrl}/checkrooms`, requestOptions)
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

async function GetCheckroomslist(id : any) {
    const requestOptions = {
        method: "GET",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
    };

    let res = await fetch(`${apiUrl}/checkroom/${id}`, requestOptions)
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


async function DeleteCheckroom(data: number) {
    let checkroom = data;
    const requestOptions = {
        method: "DELETE",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
        body: JSON.stringify(data),
    }
    
    let res = await fetch(`${apiUrl}/checkrooms/${checkroom}`, requestOptions)
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

async function UpdateCheckroom(data: CheckroomInterface) {
    let cr_id = data.ID
    const requestOptions = {
        method: "PATCH",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
        body: JSON.stringify(data),
    }

    let res = await fetch(`${apiUrl}/checkroomsupdate/${cr_id}`, requestOptions)
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
    UpdateCheckroom,
    DeleteCheckroom,
    GetCheckroomslist,
    Checkrooms,
    GetRoom,
    GetProduct,
    GetDamage, 
    GetStatus,
    GetChecklists,
};