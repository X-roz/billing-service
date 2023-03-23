import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { HttpClient } from '@angular/common/http';
import { addInventoryPayloadI, apiResponseI } from './inventory.model';

@Injectable({
  providedIn: 'root'
})
export class InventoryService {
  private BaseAPiUrl = environment.apiURL;
  route  = {
    addInventory: 'billing-service/add-item',
    getInventory: 'billing-service/get-item'
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
}
