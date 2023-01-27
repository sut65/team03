import { RoomInterface } from "./IRoom";
import { CustomerInterface } from "./modelCustomer/ICustomer";

//type
export interface RepairTypeInterface {
    ID: number;
    Name: string;
}

//main
export interface RepairReqInterface {
    ID?: number;
    Note?: string | null;
    Time?: Date | null;

    RoomID?: number | null;
    Room?: RoomInterface;
    
    RepairTypeID?: number | null;
	RepairType?:   RepairTypeInterface;

    CustomerID?: number | null;
    Customer?: CustomerInterface; 
}