import { DamageInterface } from "./modelCheckroom/IDamage";
import { StatusInterface } from "./modelCheckroom/IStatus";
import { ProductInterface } from "./IStorage";
import { EmployeeInterface } from "./IEmployee";
import { RoomInterface } from "./IRoom";

export interface CheckroomInterface {
    ID?: number;
    RoomID?: number;
    Room?: RoomInterface;

    ProductID?: number;
    Product?: ProductInterface;

    DamageID?: number;
    Damage?: DamageInterface;

    StatusID?: number;
    Status?: StatusInterface;
    Date: Date | null,

    EmployeeID?: number;
    Employee?: EmployeeInterface;
  }