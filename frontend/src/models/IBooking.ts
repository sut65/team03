import { CustomersInterface } from "./ICustomer"
import { RoomsInterface } from "./IRoom"

export interface BookingsInterface {
    ID?: string,
    BranchID?: number,
    Branch?: BrachsInterface,

    RoomID?: number,
    Room?: RoomsInterface, // waiting for Rooms Interface

    Start?: Date | null,
    Stop?: Date | null,

    CustomerID?: number,
    Customerr?: CustomersInterface,
}

export interface BrachsInterface {
    ID: number,
    Name: string,
}
