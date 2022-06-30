import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from "@angular/router";
import {RequestService} from "../../request.service";
import {NzMessageService} from "ng-zorro-antd/message";
import cryptoRandomString from "crypto-random-string";

@Component({
  selector: 'app-hmi-edit',
  templateUrl: './hmi-edit.component.html',
  styleUrls: ['./hmi-edit.component.scss']
})
export class HmiEditComponent implements OnInit {
  id: any;
  submitting = false;

  data: any = {width:800, height: 600, entities:[]}

  constructor(private route: ActivatedRoute, private router: Router, private rs: RequestService, private message: NzMessageService) {
    this.id = route.snapshot.paramMap.get('id');
    if (this.id) this.load();
    else this.data.id = cryptoRandomString({length: 20, type: 'alphanumeric'})
  }

  ngOnInit(): void {
  }


  load(): void {
    this.rs.get('hmi/' + this.id).subscribe(res => {
      this.data = res.data;
    })
  }

  submit(): void {
  }

  onSave(hmi: any) {
    console.log('save', hmi)
    this.submitting = true
    const uri = this.id ? 'hmi/' + this.id : 'hmi/create';
    this.rs.post(uri, hmi).subscribe(res => {
      this.message.success("提交成功");
      this.router.navigate(['/admin/hmi/detail/' + res.data.id]);
    }).add(() => {
      this.submitting = false;
    })
  }
}
