import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BillingHomeComponent } from './billing-home/billing-home.component';
import { BillingRoutingModule } from './billing-routing.module';



@NgModule({
  declarations: [
    BillingHomeComponent
  ],
  imports: [
    CommonModule,
    BillingRoutingModule
  ]
})
export class BillingModule { }
