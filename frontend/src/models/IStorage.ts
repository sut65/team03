import { EmployeeInterface } from "./IEmployee";

//product
export interface ProductInterface {
    ID?: number;
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
    ID?: number | null;
    Quantity?: number | null;
    Time?: Date | null;
    
    EmployeeID?: number | null;
    Employee?: EmployeeInterface;

    ProductID?: number | null;
    Product?: ProductInterface;

    ProductTypeID?: number | null;
    ProducType?: ProductTypeInterface;

}


