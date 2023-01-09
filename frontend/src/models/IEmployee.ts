export interface EmployeeInterface {

    ID?: number;
    PersonalID: number;
    Name: string;
    Email?: string;
    Eusername?: string;
    Password?: string;
    Salary: number;
    Phonenumber: string;
    Gender: string;
    DateOfBirth: Date | null;
    YearOfStart: Date | null;
    Address: string;
    Blood:  string;


    // entity
    OfficerID : number;
    Officer: OfficerInterface;

    DepartmentID : number;
    Department: DepartmentInterface;

    PositionID : number;
    Position: PositionInterface;

    LocationID : number;
    Location: LocationInterface;

   //แสดงข้อมูลมาแสดงมาจาก หลังบ้าน
}

export interface OfficerInterface {

    ID: number;
    Name: string;

   //แสดงข้อมูลมาแสดงมาจาก หลังบ้าน
}

export interface DepartmentInterface {

    ID: number;
    Name: string;
//หลังบ้าน ไปเว็บ เว็บมา models
   //แสดงข้อมูลมาแสดงมาจาก หลังบ้าน
}

export interface PositionInterface {

    ID: number;
    Name: string;
//หลังบ้าน ไปเว็บ เว็บมา models
   //แสดงข้อมูลมาแสดงมาจาก หลังบ้าน
}

export interface LocationInterface {

    ID: number;
    Name: string;

   //แสดงข้อมูลมาแสดงมาจาก หลังบ้าน
}