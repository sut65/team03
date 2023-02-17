import { CustomerInterface } from "../modelCustomer/ICustomer";

export interface PaymentsInterface {
    ID?: number ;
    CustomerID?: number | null;
    Customer?: CustomerInterface;
    PaymentMethodID?: number | null;
    PaymentMethod?: PaymentMethodsInterface;
    MethodID?: number | null;
    Method?: MethodsInterface;
    PlaceID?: number | null;
    Place?: PlacesInterface;
    Price?: number ;
    Time?: Date | null;
    Picture?: string | ArrayBuffer | null;
}

export interface PaymentMethodsInterface {
    ID?: number,
    Name: string;
}

export interface MethodsInterface {
    ID?: number,
    Name: string,
    Destination: string,
    Picture: string,
    PaymentMethodID: number,
    PaymentMethod: PaymentMethodsInterface;
}

export interface PlacesInterface {
    ID?: number,
    Name: string,
}