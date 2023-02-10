import { DepartmentInterface } from "./IEmployee";
import { CustomerInterface } from "./modelCustomer/ICustomer";

export interface ReviewInterface {

    ID: number;
    Comment: string;
	Star: number;
	Reviewdate: Date;
    Reviewimage: string;

    // entity
    CustomerID : number;
    Customer: CustomerInterface;

    DepartmentID : number;
    Department: DepartmentInterface;

    SystemworkID : number;
    Systemwork: SystemworkInterface;

   //แสดงข้อมูลมาแสดงมาจาก หลังบ้าน
}

export interface SystemworkInterface {

    ID: number;
    Name: string;
//หลังบ้าน ไปเว็บ เว็บมา models
   //แสดงข้อมูลมาแสดงมาจาก หลังบ้าน
}