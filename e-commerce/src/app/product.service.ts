import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Product } from './product';

@Injectable({
  providedIn: 'root',
})
export class ProductService {
  private url: string = 'http://localhost:8080/products';
	 //private url: string = 'https://ou3e0b2209.execute-api.ap-southeast-1.amazonaws.com/dev/products'

  httpOption = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
  };

  constructor(private httpClient: HttpClient) {}

  getProducts(): Observable<Product[]> {
    return this.httpClient.get<Product[]>(this.url);
  }

  getProduct(id: number): Observable<Product> {
    const productUrl = `${this.url}/${id}`;
    return this.httpClient.get<Product>(productUrl);
  }

  updateProduct(product: Product): Observable<Product> {
    const productUrl = `${this.url}/${product.id}`;
    return this.httpClient.put<Product>(productUrl, product, this.httpOption);
  }

  deleteProduct(product: Product): Observable<Product> {
    const productUrl = `${this.url}/${product.id}`;
    return this.httpClient.delete<Product>(productUrl);
  }
}
