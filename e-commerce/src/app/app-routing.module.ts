import { ProductDetailComponent } from './product-detail/product-detail.component';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { OurProductsComponent } from './our-products/our-products.component';
import { AboutComponent } from './about/about.component';
import { SubhomeComponent } from './subhome/subhome.component';

const routes: Routes = [
  {
    path: 'home',
    component: SubhomeComponent,
  },
  {
    path: 'products',
    component: OurProductsComponent,
  },
  {
    path: 'products/:id',
    component: ProductDetailComponent,
  },
  {
    path: 'about',
    component: AboutComponent,
  },
  {
    path: '',
    redirectTo: 'home',
    pathMatch: 'full',
  },
];
@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
