import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from "@angular/router";
import {RequestService} from "../../request.service";
import {NzMessageService} from "ng-zorro-antd/message";

@Component({
  selector: 'app-hmi-edit-content',
  templateUrl: './hmi-edit-content.component.html',
  styleUrls: ['./hmi-edit-content.component.scss']
})
export class HmiEditContentComponent implements OnInit {
  id: any;
  submitting = false;

  data: any = {width:800, height: 600, entities:[]}

  constructor(private route: ActivatedRoute, private router: Router, private rs: RequestService, private message: NzMessageService) {
    this.id = route.snapshot.paramMap.get('id');
    this.load();
  }

  ngOnInit(): void {
  }


  load(): void {
    this.rs.get('hmi/' + this.id + '/manifest').subscribe(res => {
      this.data = res;
    })
  }

  submit(): void {
  }

  onSave(hmi: any) {
    console.log('save', hmi)
    this.submitting = true
    const uri = 'hmi/' + this.id + '/manifest';
    this.rs.post(uri, hmi).subscribe(res => {
      this.message.success("提交成功");
      this.router.navigate(['/admin/hmi/detail/' + this.id]);
    }).add(() => {
      this.submitting = false;
    })
  }
}
