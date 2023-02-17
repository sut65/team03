import { StorageInterface } from "../IStorage";
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

    StorageID?: number,
    Storage?: StorageInterface,
    StorageItem?: number;
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