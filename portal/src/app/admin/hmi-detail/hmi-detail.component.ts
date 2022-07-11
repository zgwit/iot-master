import {Component, OnInit} from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {RequestService} from "../../request.service";

@Component({
  selector: 'app-hmi-detail',
  templateUrl: './hmi-detail.component.html',
  styleUrls: ['./hmi-detail.component.scss']
})
export class HmiDetailComponent implements OnInit {
  id = '';
  data: any = {};
  loading = false;
  hmi: any = {};

  constructor(private router: ActivatedRoute, private rs: RequestService) {
    this.id = router.snapshot.params['id'];
    this.load();
  }

  ngOnInit(): void {
  }

  load(): void {
    this.loading = true;
    this.rs.get(`hmi/${this.id}`).subscribe(res=>{
      this.data = res.data;
      this.loading = false;
    });
    this.rs.get(`hmi/${this.id}/manifest`).subscribe(res=>{
      this.hmi = res;
    });
  }
}
