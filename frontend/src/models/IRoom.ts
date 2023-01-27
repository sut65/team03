import { EmployeeInterface } from "./IEmployee";

//roomtype
export interface RoomTypeInterface {
    ID: number;
    Size:    string;
	Bedsize: string;
	Bedtype: string;
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
    Time?: Date;
    
    EmployeeID?: number;
    Employee?: EmployeeInterface;

    RoomTypeID?: number;
    RoomType?: RoomTypeInterface;

    RoomZoneID?: number;
    RoomZone?: RoomZoneInterface;

    StateID?: number;
    State?: StateInterface;
}


