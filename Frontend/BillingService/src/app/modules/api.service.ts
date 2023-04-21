import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { HttpClient } from '@angular/common/http';
import { addInventoryPayloadI, apiResponseI } from './inventory/inventory.model';

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  private BaseAPiUrl = environment.apiURL;
  route  = {
    getInventory: 'inventory/get-items',
  }
  constructor(private http: HttpClient) { }
  getInventory(){
    const url = this.BaseAPiUrl + this.route.getInventory
    return this.http.get<apiResponseI<addInventoryPayloadI[]>>(url)
  }
}
