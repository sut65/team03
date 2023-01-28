import { PaymentsInterface } from "../modelPayment/IPayment"
import { EmployeeInterface } from "../IEmployee"
import { StatusesInterface } from "./IStatus"

export interface CHK_PaymentsInterface {
    ID?: number,
    PaymentID?: number,
    Payment?: PaymentsInterface,

    StatusID?: number,
    Status?: StatusesInterface,

    Date_time?: Date | null,
    Amount?: number,
    Description?: string,

    EmployeeID?: number,
    Employee?: EmployeeInterface,
}