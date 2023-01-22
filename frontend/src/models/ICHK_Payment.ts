// import { PaymentsInterface } from "./IPayment"
import { UsersInterface } from "./IUser"

export interface CHK_PaymentsInterface {
    ID?: number,
    // PaymentID?: number,
    // Payment?: PaymentsInterface,

    StatusID?: number,
    Status?: StatusesInterface,

    Date_time?: Date | null,
    Amount?: number,
    Description?: string,

    EmployeeID?: number,
    Employee?: UsersInterface,
}

export interface StatusesInterface {
    ID: number,
    Type: string,
}