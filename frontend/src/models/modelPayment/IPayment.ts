import { CustomerInterface } from "../modelCustomer/ICustomer";

export interface PaymentsInterface {
    ID?: number,
    CustomerID: number,
    Customer: CustomerInterface;
    PaymentMethodID: number,
    PaymentMethod: PaymentMethodsInterface;
    MethodID: number,
    Method: MethodsInterface
    PlaceID: number,
    Place: PlacesInterface
    Time: Date | null,
    Picture: string;
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