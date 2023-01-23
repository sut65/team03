import { time } from "console";
import { BookingsInterface } from "./IBooking";
import { EmployeeInterface } from "./IEmployee";

//status
export interface ChecKInOutStatusInterface {
    ID: number;
    Name: string;
}

//main
export interface CheckInOutInterface {
    ID?: number;
    CheckInTime: Date;
    CheckOutTime?: Date;

    BookingID: number;
    Booking: BookingsInterface;

    CheckInOutStatusID: number;
    CheckInOutStatus: ChecKInOutStatusInterface;

    EmployeeID: number;
    Employee: EmployeeInterface;
}