import { RoomInterface, RoomTypeInterface, RoomZoneInterface, StateInterface } from "../../../models/IRoom";

const apiUrl = "http://localhost:8080";
	//protected.GET("/rooms", room.ListRooms)
    // protected.GET("/room/:id", room.GetRoom)
    // protected.POST("/rooms", room.CreateRoom)
    // protected.PATCH("/rooms", room.UpdateRoom)
    // protected.DELETE("/rooms/:id", room.DeleteRoom)

    // protected.GET("/room_types", room.ListRoomTypes)
    // protected.GET("/room_types/:id", room.GetRoomType)
    // protected.POST("/room_types", room.CreateRoomType)
    // protected.PATCH("/room_types", room.UpdateRoomType)
    // protected.DELETE("/room_typss/:id", room.DeleteRoomType)

    // protected.GET("/room_zones", room.ListRoomZones)
    // protected.GET("/room_zone/:id", room.GetRoomZone)
    // protected.POST("/room_zones", room.CreateRoomZone)
    // protected.PATCH("/room_zones", room.UpdateRoomZone)
    // protected.DELETE("/room_zones/:id", room.DeleteRoomZone)

    // protected.GET("/states", room.ListStates)
    // protected.GET("/state/:id", room.GetState)
    // protected.POST("/states", room.CreateState)
    // protected.PATCH("/states", room.UpdateState)
    // protected.DELETE("/states/:id", room.DeleteState)
// List Room
async function GetRooms() {
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

// List Room
async function GetRoom(id : any) {
    const requestOptions = {
        method: "GET",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
    };

    let res = await fetch(`${apiUrl}/room/${id}`, requestOptions)
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

// List RoomType
async function GetRoomTypes() {
    const requestOptions = {
        method: "GET",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
    };

    let res = await fetch(`${apiUrl}/room_types`, requestOptions)
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

// List RoomZone
async function GetRoomZones() {
    const requestOptions = {
        method: "GET",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
    };

    let res = await fetch(`${apiUrl}/room_zones`, requestOptions)
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

// List State
async function GetStates() {
    const requestOptions = {
        method: "GET",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
    };

    let res = await fetch(`${apiUrl}/states`, requestOptions)
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

// List Employee
async function GetEmployees() {
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

//Craete room
async function CreateRoom(data: RoomInterface) {
    const requestOptions = {
        method: "POST",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    };

    let res = await fetch(`${apiUrl}/rooms`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                return {status: true, message: res.data};
            } else {
                return {status: false, message: res.error};
            }
        });

    return res;
}

//protected.PATCH("/rooms/id", room.Room)
async function Room(data: number) {
    let RoomID = data;
    const requestOptions = {
        method: "PATCH",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
        body: JSON.stringify(data),
    }

    let res = await fetch(`${apiUrl}/rooms/${RoomID}`, requestOptions)
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

// Delete Room
async function DeleteRoom(data: number) {
    let RoomID = data;
    const requestOptions = {
        method: "DELETE",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
        body: JSON.stringify(data),
    }
    
    let res = await fetch(`${apiUrl}/rooms/${RoomID}`, requestOptions)
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

//update room
async function UpdateRoom(data: RoomInterface) {
    const requestOptions = {
        method: "PUT",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
        body: JSON.stringify(data),
    }

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

//update roomtype
async function UpdateRoomType(data: RoomTypeInterface) {
    const requestOptions = {
        method: "PUT",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
        body: JSON.stringify(data),
    }

    let res = await fetch(`${apiUrl}/room_types`, requestOptions)
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

//update roomzone
async function UpdateRoomZone(data: RoomZoneInterface) {
    const requestOptions = {
        method: "PUT",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
        body: JSON.stringify(data),
    }

    let res = await fetch(`${apiUrl}/room_zones`, requestOptions)
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

//update state
async function UpdateState(data: StateInterface) {
    const requestOptions = {
        method: "PUT",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
        body: JSON.stringify(data),
    }

    let res = await fetch(`${apiUrl}/states`, requestOptions)
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
    GetRooms,
    GetRoomTypes,
    GetRoomZones,
    GetStates,
    CreateRoom,
    GetEmployees,
    Room,
    DeleteRoom,
    UpdateRoom,
    UpdateRoomType,
    UpdateRoomZone,
    UpdateState,
    GetRoom,
};
