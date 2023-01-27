
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
// List CheckInOut
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
export {
    GetRooms,
};