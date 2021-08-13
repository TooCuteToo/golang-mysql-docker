import { ProductComponent } from './products/product/product.component';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { ReactiveFormsModule } from '@angular/forms';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HeaderComponent } from './header/header.component';
import { BannerComponent } from './banner/banner.component';
import { SubhomeComponent } from './subhome/subhome.component';
import { FooterComponent } from './footer/footer.component';
import { ProductsComponent } from './products/products.component';
import { AboutComponent } from './about/about.component';
import { TeamMembersComponent } from './team-members/team-members.component';
import { ServicesComponent } from './services/services.component';
import { CustomersComponent } from './customers/customers.component';
import { OurProductsComponent } from './our-products/our-products.component';
import { ProductDetailComponent } from './product-detail/product-detail.component';

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    BannerComponent,
    SubhomeComponent,
    FooterComponent,
    ProductsComponent,
    ProductComponent,
    AboutComponent,
    TeamMembersComponent,
    ServicesComponent,
    CustomersComponent,
    OurProductsComponent,
    ProductDetailComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ReactiveFormsModule,
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
