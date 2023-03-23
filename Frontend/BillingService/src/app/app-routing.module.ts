import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  {
    path: 'inventory',
    loadChildren: ()=>import('./modules/inventory/inventory.module').then(module=>module.InventoryModule)
  },
  {
    path: 'billing',
    loadChildren: ()=>import('./modules/billing/billing.module').then(module=> module.BillingModule)
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
