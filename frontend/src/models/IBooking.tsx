import { UsersInterface } from "./IUser"

export interface BookingsInterface {
    ID?: string,
    BranchID?: number,
    Branch?: BrachsInterface,

    // RoomID?: number,
    // Room?: RoomsInterface, // waiting for Rooms Interface

    Start?: Date | null,
    Stop?: Date | null,

    UserID?: number,
    User?: UsersInterface,
}

export interface BrachsInterface {
    ID: number,
    Name: string,
}