import { Product } from './../product';
import { Location } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormBuilder } from '@angular/forms';
import { ActivatedRoute } from '@angular/router';

import { ProductService } from '../product.service';

@Component({
  selector: 'app-product-detail',
  templateUrl: './product-detail.component.html',
  styleUrls: ['./product-detail.component.scss'],
})
export class ProductDetailComponent implements OnInit {
  product: Product = {
    id: 1,
    name: '',
    description: '',
    price: 10,
    rate: 2,
    image: '',
  };
  checkoutForm = this.formBuilder.group({
    name: '',
    desc: '',
    price: '',
  });

  constructor(
    private formBuilder: FormBuilder,
    private route: ActivatedRoute,
    private productService: ProductService,
    private location: Location
  ) {}

  ngOnInit(): void {
    const id = Number(this.route.snapshot.paramMap.get('id'))!;
    this.productService
      .getProduct(id)
      .subscribe((product) => (this.product = product));
  }

  onSubmit() {
    this.product.price = Number(this.product.price);
    this.productService
      .updateProduct(this.product)
      .subscribe(() => this.location.back(), (err) => console.log(err));
  }

  log() {
    console.log(this.product);
  }
}
