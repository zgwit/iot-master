import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import {RequestService} from "../../request.service";

@Component({
  selector: 'app-protocol',
  templateUrl: './protocol.component.html',
  styleUrls: ['./protocol.component.scss']
})
export class ProtocolComponent implements OnInit {
  datum: any[] = [];

  loading = false;


  constructor(private router: Router, private rs: RequestService) {
  }

  ngOnInit(): void {
    this.load();
  }


  load(): void {
    this.loading = true;
    this.rs.get('system/protocols').subscribe(res => {
      console.log('res', res);
      this.datum = res.data;
    }).add(() => {
      this.loading = false;
    });
  }

}
