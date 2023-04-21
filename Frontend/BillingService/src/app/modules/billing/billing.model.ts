export interface salesDataPayloadI{
    Name: string,
    tax: number
    price: number,
    quantity: number,
    total: number
}
export interface salesApiPayloadI{
    customerName: string,
    customerPhoneNumber: string,
    salesDetails: salesDataPayloadI[],
    summaryTotal: number
}