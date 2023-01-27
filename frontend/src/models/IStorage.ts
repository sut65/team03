import { EmployeeInterface } from "./IEmployee";

//product
export interface ProductInterface {
    ID: number;
    Name?: string;
    Price: number;
}

//Producttype
export interface ProductTypeInterface {
    ID: number;
    Name?: string;
}

//main
export interface StorageInterface {
    ID?: number;
    Quantity: number;
    Time?: Date;
    
    EmployeeID?: number;
    Employee?: EmployeeInterface;

    ProductID?: number;
    Product?: ProductInterface;

    ProductTypeID?: number;
    ProducType?: ProductTypeInterface;

}


