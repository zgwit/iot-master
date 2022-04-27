import {Component, OnInit} from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {RequestService} from "../../request.service";

@Component({
  selector: 'app-tunnel-detail',
  templateUrl: './tunnel-detail.component.html',
  styleUrls: ['./tunnel-detail.component.scss']
})
export class TunnelDetailComponent implements OnInit {
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
    this.rs.get(`tunnel/${this.id}`).subscribe(res=>{
      this.data = res.data;
      this.loading = false;
    });
    //TODO 监听
  }

  enable($event: any) {

  }


}
