import { EmployeeInterface } from "./IEmployee";

//roomtype
export interface RoomTypeInterface {
    ID: number;
    Size:    string;
	Bedsize: string;
	Bedtype: string;
    Amount: number;
}

//roomzone
export interface RoomZoneInterface {
    ID: number;
    Name?: string;
}

//state
export interface StateInterface {
    ID: number;
    Name?: string;
}

//main
export interface RoomInterface {
    ID?: number;
    Room_No?: string;
    Amount?: number ;
    Time?: Date | null;
    
    EmployeeID?: number | null;
    Employee?: EmployeeInterface;

    RoomTypeID?: number | null;
    RoomType?: RoomTypeInterface;

    RoomZoneID?: number | null;
    RoomZone?: RoomZoneInterface;

    StateID?: number | null;
    State?: StateInterface;
}


