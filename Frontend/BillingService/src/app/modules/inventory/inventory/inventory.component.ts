import { Component, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MessageService } from 'primeng/api';
import { Table } from 'primeng/table';
import { addInventoryPayloadI, updateInventoryPayloadI } from '../inventory.model';
import { InventoryService } from '../inventory.service';

@Component({
  selector: 'kcs-inventory',
  templateUrl: './inventory.component.html',
  styleUrls: ['./inventory.component.scss']
})
export class InventoryComponent implements OnInit {
  @ViewChild('dt1') inventoryTable!: Table;
  visible = false;
  inventoryData: addInventoryPayloadI[] = [];
  inventoryGroup: FormGroup;
  searchControl: FormControl;
  constructor(private fb: FormBuilder, 
    private api: InventoryService, 
    private messageService: MessageService) {
    this.searchControl = this.fb.control('')
    this.inventoryGroup = this.fb.group({
      name: ['',Validators.required],
      tax: [null,Validators.required],
      price: [null,Validators.required],
      quantity: [null,Validators.required]
    })
   }

  ngOnInit(): void {
    this.getInventory();
    this.searchControl.valueChanges.subscribe(value => {
      this.inventoryTable.filterGlobal(value,'contains')
    })
  }
  getInventory(){
    this.api.getInventory().subscribe(data => {
      this.inventoryData = data.Data!
    })
  }
  addInventory(){
    this.visible = true;
  }
  submitInventory(){
    this.api.addInventory(this.inventoryGroup.value).subscribe(data => {
      this.messageService.add({ severity: 'success', summary: 'Success', detail: data.Message });
      this.visible=false;
      this.getInventory()
    })
  }
  onRowEditSave(inventory: updateInventoryPayloadI) {
    this.api.updateInventory(inventory).subscribe(data => {
      console.log(data);
      this.messageService.add({ severity: 'success', summary: 'Success', detail: data.Message });
    },error =>{
      this.messageService.add({ severity: 'error', summary: 'Error', detail: error });
    })
  }
}
