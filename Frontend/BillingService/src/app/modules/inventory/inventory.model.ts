export interface addInventoryPayloadI{
    name: string,
    tax: number
    price: number,
    quantity: number
}
export interface apiResponseI<T={}>{
    Success: boolean;
    Message: string;
    Data?: T
}
export interface updateInventoryPayloadI{
  "ID":number,
  "name":string,
  "tax": number,
  "price":number, 
  "quantity":number
}