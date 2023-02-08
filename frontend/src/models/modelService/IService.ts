import { CustomerInterface } from "../modelCustomer/ICustomer";

export interface ServicesInterface {
    ID?: number,
    CustomerID?: number | null,
    Customer?: CustomerInterface,
    Time?: Date | null,
    FoodID?: number,
    Food?: FoodsInterface,
    DrinkID?: number,
    Drink?: DrinksInterface,
    AccessoriesID?: number,
    Accessories?: AccessoriesInterface;
}

export interface FoodsInterface {
    ID?: number,
    Name?: string,
    Price?: number,
    Item?: number;
}

export interface DrinksInterface {
    ID?: number,
    Name: string,
    Price: number,
    Item: number;
}

export interface AccessoriesInterface {
    ID?: number,
    Name: string,
    Price: number,
    Item: number;
}