import { CustomerInterface } from "../modelCustomer/ICustomer";

export interface ServicesInterface {
    ID?: number,
    CustomerID?: number | null,
    Customer?: CustomerInterface,
    Time?: Date | null,
    FoodID?: number,
    Food?: FoodsInterface,
    FoodItem?: number,
    DrinkID?: number,
    Drink?: DrinksInterface,
    DrinkItem?: number,
    AccessoriesID?: number,
    Accessories?: AccessoriesInterface;
    AccessoriesItem?: number,
}

export interface FoodsInterface {
    ID?: number,
    Name?: string,
    Price?: number,
    Item?: number;
}

export interface DrinksInterface {
    ID?: number,
    Name?: string,
    Price?: number,
    Item?: number;
}

export interface AccessoriesInterface {
    ID?: number,
    Name?: string,
    Price?: number,
    Item?: number;
}