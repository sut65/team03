import { UsersInterface } from "./IUser";

//type
export interface RepairTypeInterface {
    ID: number;
    Name: string;
}

//status
export interface RepairStatusInterface {
    ID: number;
    Name: string;
}

//main
export interface RepairReqInterface {
    ID: number;
    Note: string;
    Time: Date;

    // RoomID: number;
    // Room: RoomInterface;
    
    RepairStatusID: number;
    RepairStatus: RepairStatusInterface;

    UserID: number;
    User: UsersInterface; 
}