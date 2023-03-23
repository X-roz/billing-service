import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MessageService } from 'primeng/api';
import { addInventoryPayloadI } from '../inventory.model';
import { InventoryService } from '../inventory.service';

@Component({
  selector: 'kcs-inventory',
  templateUrl: './inventory.component.html',
  styleUrls: ['./inventory.component.scss']
})
export class InventoryComponent implements OnInit {
  visible = false;
  inventoryData: addInventoryPayloadI[] = [];
  inventoryGroup: FormGroup;
  constructor(private fb: FormBuilder, 
    private api: InventoryService, 
    private messageService: MessageService) {
    this.inventoryGroup = this.fb.group({
      name: ['',Validators.required],
      tax: [null,Validators.required],
      price: [null,Validators.required],
      quantity: [null,Validators.required]
    })
   }

  ngOnInit(): void {
    this.getInventory()
  }
  getInventory(){
    this.api.getInventory().subscribe(data => {
      this.inventoryData = data.Data
    })
  }
  addInventory(){
    this.visible = true;
  }
  submitInventory(){
    this.api.addInventory(this.inventoryGroup.value).subscribe(data => {
      console.log(data);
      this.messageService.add({ severity: 'success', summary: 'Success', detail: data.Message });
      this.visible=false;
      this.getInventory()
    })
  }
}
