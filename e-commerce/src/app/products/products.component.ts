import { ProductService } from './../product.service';
import { Observable } from 'rxjs';
import { Component, OnInit } from '@angular/core';
import { Location } from '@angular/common';
import { map } from 'rxjs/operators';

import { Product } from '../product';

@Component({
  selector: 'app-products',
  templateUrl: './products.component.html',
  styleUrls: ['./products.component.scss'],
})
export class ProductsComponent implements OnInit {
  products!: Product[];
  currentPath!: string;

  constructor(
    private productService: ProductService,
    private location: Location
  ) {}

  ngOnInit(): void {
    this.productService
      .getProducts()
      .subscribe((products) => (this.products = products));
    this.currentPath = this.location.path();
  }

  deleteProduct(product: Product) {
    this.productService
      .deleteProduct(product)
      .subscribe(
        () =>
          (this.products = this.products.filter(
            (item) => item.id !== product.id
          ))
      );
  }
}
