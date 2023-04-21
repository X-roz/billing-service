import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BillingHomeComponent } from './billing-home/billing-home.component';
import { BillingRoutingModule } from './billing-routing.module';
import { TabViewModule } from 'primeng/tabview';
import { BillingFormComponent } from './billing-form/billing-form.component';
import { DropdownModule } from 'primeng/dropdown';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { TableModule } from 'primeng/table';


@NgModule({
  declarations: [
    BillingHomeComponent,
    BillingFormComponent
  ],
  imports: [
    CommonModule,
    BillingRoutingModule,
    ReactiveFormsModule,
    FormsModule,
    TabViewModule,
    DropdownModule,
    TableModule
  ]
})
export class BillingModule { }
