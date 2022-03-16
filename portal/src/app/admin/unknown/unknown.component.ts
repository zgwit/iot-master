import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import {RequestService} from "../../request.service";

@Component({
  selector: 'app-unknown',
  templateUrl: './unknown.component.html',
  styleUrls: ['./unknown.component.scss']
})
export class UnknownComponent implements OnInit {



  constructor(private router: Router, private rs: RequestService) {

  }

  ngOnInit(): void {
  }

}
