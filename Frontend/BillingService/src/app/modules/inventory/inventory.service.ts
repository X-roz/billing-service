import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { HttpClient } from '@angular/common/http';
import { addInventoryPayloadI, apiResponseI, updateInventoryPayloadI } from './inventory.model';

@Injectable({
  providedIn: 'root'
})
export class InventoryService {
  private BaseAPiUrl = environment.apiURL;
  route  = {
    addInventory: 'inventory/add-item',
    getInventory: 'inventory/get-items',
    updateInventory: 'inventory/update-item'
  }
  constructor(private http: HttpClient) { }
  
  addInventory(payload: addInventoryPayloadI){
    const url = this.BaseAPiUrl + this.route.addInventory
    return this.http.post<apiResponseI<string>>(url,payload)
  }
  getInventory(){
    const url = this.BaseAPiUrl + this.route.getInventory
    return this.http.get<apiResponseI<addInventoryPayloadI[]>>(url)
  }
  updateInventory(payload: updateInventoryPayloadI){
    const url = this.BaseAPiUrl + this.route.updateInventory;
    return this.http.post<apiResponseI>(url,payload)
  }
}
