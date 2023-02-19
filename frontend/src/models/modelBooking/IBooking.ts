import { CustomerInterface } from "../modelCustomer/ICustomer"
import { RoomInterface } from "../IRoom"
import { BranchsInterface } from "./IBranch"

export interface BookingsInterface {
    ID?: number,

    BranchID?: number,
    Branch?: BranchsInterface,

    RoomID?: number,
    Room?: RoomInterface,

    Start?: Date | null,
    Stop?: Date | null,

    // Auto generate in backend
    Booking_Number?: string,
    Tx_No?: string,
    Total?: number, // TypeScript's number type represents a floating-point number and can store integers and floating-point numbers.
    DayEech?: Date | null,
    TotalAmount?: number,
    Num_Of_Day?: number,
    // Auto generate in backend

    CustomerID?: number,
    Customer?: CustomerInterface,
}