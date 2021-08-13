import { Component, OnInit } from '@angular/core';
import { Location } from '@angular/common';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss'],
})
export class HeaderComponent implements OnInit {
  currentPath!: string;

  constructor(private location: Location) {}

  ngOnInit(): void {
    this.currentPath = this.location.path();
  }
}
