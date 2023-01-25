// import { PaymentsInterface } from "./IPayment"
// import { EmployeeInterface } from "./IEmployee"

export interface CHK_PaymentsInterface {
    ID?: number,
    // PaymentID?: number,
    // Payment?: PaymentsInterface,

    StatusID?: number,
    Status?: StatusesInterface,

    Date_time?: Date | null,
    Amount?: number,
    Description?: string,

    // EmployeeID?: number,
    // Employee?: EmployeeInterface,
}

export interface StatusesInterface {
    ID: number,
    Type: string,
}