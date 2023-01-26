import { GenderInterface } from "./IGender";
import { ProvinceInterface } from "./IProvince";
import { NametitleInterface } from "./INametitle";

export interface CustomerInterface {
  ID?: number;
  FirstName?: string;
  LastName?: string;
  Password?: string;
  Age?: number;
  Phone?: number;
  Email?: string;
  Gender?: GenderInterface;
  Gender_ID?: number;
  Nametitle?: NametitleInterface;
  Nametitle_ID?: number;
  Province_ID?: number;
  Province?: ProvinceInterface;
}
