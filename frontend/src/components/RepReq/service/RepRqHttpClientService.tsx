import { CheckInOutInterface } from "../../../models/ICheckInOut";
import { RepairReqInterface } from "../../../models/IRepairReq";

const apiUrl = "http://localhost:8080";
			// //========================= repreq routes
			// //type
			// protected.GET("/repairtype/:id", repreq.GetRepairType)
			// protected.GET("/repairtypes", repreq.ListRepairTypes)
			// //status
			// protected.GET("/repairstatus/:id", repreq.GetRepairStatus)
			// protected.GET("/repairstatuses", repreq.ListRepairStatuses)
			// //main
			// protected.GET("/repairreq/:id", repreq.GetRepairReq)
			// protected.GET("/repairreqs", repreq.ListRepairReqs)
			// protected.POST("/repairreq", repreq.CreateRepairReq)
			// protected.PATCH("/repairreq", repreq.UpdateRepairReq)
			// protected.DELETE("/repairreq/:id", repreq.DeleteRepairReq)

// List Type
async function GetRepairTypes() {
    const requestOptions = {
        method: "GET",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
    };

    let res = await fetch(`${apiUrl}/repairtypes`, requestOptions)
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
async function GetRepairStatuses() {
    const requestOptions = {
        method: "GET",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
    };

    let res = await fetch(`${apiUrl}/repairstatuses`, requestOptions)
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

// List repairreqs
async function GetRepairReqs() {
    const requestOptions = {
        method: "GET",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
    };

    let res = await fetch(`${apiUrl}/repairreqs`, requestOptions)
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

// List customer
async function GetCustomers() {
    const requestOptions = {
        method: "GET",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
    };

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

//Craete repairreq
async function CreateRepairReq(data: RepairReqInterface) {
    const requestOptions = {
        method: "POST",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    };

    let res = await fetch(`${apiUrl}/repairreq`, requestOptions)
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

// Delete /repairreq/:id
async function DeleteRepairReq(data: number) {
    let repairreqID = data;
    const requestOptions = {
        method: "DELETE",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
        body: JSON.stringify(data),
    }
    
    let res = await fetch(`${apiUrl}/repairreq/${repairreqID}`, requestOptions)
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

// Update /repairreq
async function UpdateRepairReq(data: RepairReqInterface) {
    const requestOptions = {
        method: "PATCH",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
        body: JSON.stringify(data),
    }

    let res = await fetch(`${apiUrl}/repairreq`, requestOptions)
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
    GetRepairTypes,
    GetRepairStatuses,
    GetRepairReqs,
    CreateRepairReq,
    UpdateRepairReq,
    DeleteRepairReq,
    GetCustomers,
};

