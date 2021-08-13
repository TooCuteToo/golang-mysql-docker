import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';

import { Product } from '../../product';

@Component({
  selector: 'app-product',
  templateUrl: './product.component.html',
  styleUrls: ['./product.component.scss'],
})
export class ProductComponent implements OnInit {
  @Input() product!: Product;
  @Output() delete = new EventEmitter();

  rates!: number[];
  currentParam!: string | null;

  constructor() {}

  ngOnInit(): void {

    this.rates = Array(Math.round(this.product.rate)).fill('');
  }

  click(product: Product) {
    this.delete.emit(product);
  }
}
