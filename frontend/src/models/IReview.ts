import { DepartmentInterface } from "./IEmployee";

export interface ReviewInterface {

    ID: number;
    Comment: string;
	Start: number;
	Reviewdate: Date;
    Reviewimega: string;

    // entity
    CustomerID : number;
    Customer: CustomerInterface;

    DepartmentID : number;
    Department: DepartmentInterface;

    SystemworkID : number;
    Systemwork: SystemworkInterface;

   //แสดงข้อมูลมาแสดงมาจาก หลังบ้าน
}

export interface OfficerInterface {

    ID: number;
    Officername: string;

   //แสดงข้อมูลมาแสดงมาจาก หลังบ้าน
}

export interface SystemworkInterface {

    ID: number;
    Name: string;
//หลังบ้าน ไปเว็บ เว็บมา models
   //แสดงข้อมูลมาแสดงมาจาก หลังบ้าน
}

export interface CustomerInterface {

    ID: number;
    FirstName: string;
//หลังบ้าน ไปเว็บ เว็บมา models
   //แสดงข้อมูลมาแสดงมาจาก หลังบ้าน
}