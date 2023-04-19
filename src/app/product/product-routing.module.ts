import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {ProductsComponent} from "./products/products.component";
import {ProductEditComponent} from "./product-edit/product-edit.component";
import {PageNotFoundComponent} from "../base/page-not-found/page-not-found.component";
import {ProductDetailComponent} from "./product-detail/product-detail.component";

const routes: Routes = [
  {path: '', pathMatch: "full", redirectTo: "list"},
  {path: 'list', component: ProductsComponent},
  {path: 'detail/:id', component: ProductDetailComponent},
  {path: 'edit/:id', component: ProductEditComponent},
  {path: 'create', component: ProductEditComponent},
  {path: '**', component: PageNotFoundComponent}
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ProductRoutingModule {
}
