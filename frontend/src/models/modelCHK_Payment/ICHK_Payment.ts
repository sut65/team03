import { PaymentsInterface } from "../modelPayment/IPayment"
import { EmployeeInterface } from "../IEmployee"
import { CHK_PaymentStatusesInterface } from "./IStatus"

export interface CHK_PaymentsInterface {
    ID?: number,
    PaymentID?: number,
    Payment?: PaymentsInterface,

    CHK_PaymentStatusID?: number,
    CHK_PaymentStatus?: CHK_PaymentStatusesInterface,

    Date_time?: Date | null,
    Amount?: number,
    Description?: string,

    EmployeeID?: number | null,
    Employee?: EmployeeInterface,
}