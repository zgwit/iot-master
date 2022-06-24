import { Component } from '@angular/core';
import {InfoService} from "./info.service";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  constructor(private is: InfoService) {
  }
}
