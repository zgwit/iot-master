import { Component, OnInit } from '@angular/core';
import {AdminComponent} from "../admin.component";

@Component({
  selector: 'app-dash',
  templateUrl: './dash.component.html',
  styleUrls: ['./dash.component.scss']
})
export class DashComponent implements OnInit {

  constructor(admin: AdminComponent) {

  }

  ngOnInit(): void {
  }

}
