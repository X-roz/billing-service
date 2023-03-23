import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { BillingHomeComponent } from './billing-home/billing-home.component';

const routes: Routes = [
  {
    path: '',
    component: BillingHomeComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class BillingRoutingModule { }
