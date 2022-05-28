import {Component, OnInit} from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {RequestService} from "../../request.service";

@Component({
  selector: 'app-product-detail',
  templateUrl: './product-detail.component.html',
  styleUrls: ['./product-detail.component.scss']
})
export class ProductDetailComponent implements OnInit {
  id = '';
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
    this.rs.get(`product/${this.id}`).subscribe(res=>{
      this.data = res.data;
      this.loading = false;
    });
  }

}
