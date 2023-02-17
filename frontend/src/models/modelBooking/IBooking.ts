import { CustomerInterface } from "../modelCustomer/ICustomer"
import { RoomInterface } from "../IRoom"
import { BranchsInterface } from "./IBranch"

export interface BookingsInterface {
    ID?: number,
    Booking_Number?: string,
    Tx_No?: string,
    
    BranchID?: number,
    Branch?: BranchsInterface,

    RoomID?: number,
    Room?: RoomInterface,

    Start?: Date | null,
    Stop?: Date | null,

    CustomerID?: number,
    Customer?: CustomerInterface,
}