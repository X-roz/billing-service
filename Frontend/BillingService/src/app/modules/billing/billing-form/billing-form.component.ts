import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { ApiService } from '../../api.service';
import { dropdownOptionsI } from '../../commonModel';
import { addInventoryPayloadI } from '../../inventory/inventory.model';
import { salesApiPayloadI, salesDataPayloadI } from '../billing.model';

@Component({
  selector: 'kcs-billing-form',
  templateUrl: './billing-form.component.html',
  styleUrls: ['./billing-form.component.scss']
})
export class BillingFormComponent implements OnInit {
  inventoryOptions: dropdownOptionsI[] = [];
  salesFormGroup: FormGroup;
  inventoryData: addInventoryPayloadI[] = [];
  salesData: salesDataPayloadI[] = []
  summaryTotal: number = 0;
  get customerDetailsGroup(){
    return this.salesFormGroup.get('customerDetails') as FormGroup
  } 
  get salesDetailsGroup(){
    return this.salesFormGroup.get('salesDetails') as FormGroup
  } 
  get inventoryNameControl(){
    return this.salesDetailsGroup.get('Name') as FormControl
  }
  get quantityControl(){
    return this.salesDetailsGroup.get('quantity') as FormControl
  }
  get totalControl(){
    return this.salesDetailsGroup.get('total') as FormControl
  }
  constructor(private fb: FormBuilder,
    private commonApi: ApiService){
    this.salesFormGroup = this.fb.group({
      customerDetails: this.fb.group({
        customerName: [],
        customerPhoneNumber: []
      }),
      salesDetails: this.fb.group({
        Name: [''],
        price: [{value: '',disabled: true}],
        tax: [{value: '',disabled: true}],
        quantity: [''],
        total: ['']
      })
    });
  }
  ngOnInit(){
    this.commonApi.getInventory().subscribe(data => {
      this.inventoryData = data.Data ? data.Data : []
      this.inventoryOptions = this.inventoryData.map(data => {
        return {
          Text: data.name,
          Value: data.name
        }
      })
    })
    this.valueChangesRegister();
  }
  valueChangesRegister(){
    this.inventoryNameControl.valueChanges.subscribe(data =>{
      const selectedInventory = this.inventoryData?.find(inventory => inventory.name === data)
      this.salesDetailsGroup.patchValue({
        tax: selectedInventory?.tax,
        price: selectedInventory?.price
      })
    });
    this.quantityControl.valueChanges.subscribe(data => {
      const total = data * this.salesDetailsGroup.get('price')?.value
      this.totalControl.setValue(total);
    })
  }
  addSales(){
    this.salesData.push(this.salesDetailsGroup.getRawValue());
    this.summaryTotal = this.summaryTotal + this.totalControl.value
    this.salesDetailsGroup.reset();
  }
  submitSalesBill(){
    const Payload: salesApiPayloadI = {
      ...this.customerDetailsGroup.value,
      salesDetails: this.salesData,
      summaryTotal: +this.summaryTotal
    }
    console.log(Payload)
  }
}
