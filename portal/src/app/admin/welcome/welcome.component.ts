import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import {RequestService} from "../../request.service";

@Component({
  selector: 'app-welcome',
  templateUrl: './welcome.component.html',
  styleUrls: ['./welcome.component.scss']
})
export class WelcomeComponent implements OnInit {



  constructor(private router: Router, private rs: RequestService) {

  }

  ngOnInit(): void {
  }

}
