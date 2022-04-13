import {Component, OnInit} from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {RequestService} from "../../request.service";

@Component({
  selector: 'app-link-detail',
  templateUrl: './link-detail.component.html',
  styleUrls: ['./link-detail.component.scss']
})
export class LinkDetailComponent implements OnInit {
  id: any = '';
  data: any = {};
  loading = false;

  constructor(private router: ActivatedRoute, private rs: RequestService) {
    this.id = router.snapshot.params['id'];
    this.load();
  }

  ngOnInit(): void {
  }

  load(): void {
    this.loading = true;
    this.rs.get(`link/${this.id}/compose`).subscribe(res=>{
      this.data = res.data;
      this.loading = false;
    });
    //TODO 监听
  }

  disabled($event: any) {

  }

}
