import { BookingsInterface } from "./modelBooking/IBooking";
import { EmployeeInterface } from "./IEmployee";

//main
export interface CheckInOutInterface {
    ID?: number;

    BookingID?: number | null;
    Booking?: BookingsInterface;

    CheckInOutStatusID?: number | null;
    CheckInOutStatus?: CheckInOutStatusInterface;

    EmployeeID?: number | null;
    Employee?: EmployeeInterface;

    CheckInTime?: Date | null;
    CheckOutTime?: Date | null;
}

//status
export interface CheckInOutStatusInterface {
    ID: number;
    Name: string;
}
