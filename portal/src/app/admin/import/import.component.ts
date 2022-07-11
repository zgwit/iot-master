import {Component, Input, OnInit} from '@angular/core';
import {NzUploadChangeParam} from "ng-zorro-antd/upload";
import {NzMessageService} from "ng-zorro-antd/message";

@Component({
  selector: 'app-import',
  templateUrl: './import.component.html',
  styleUrls: ['./import.component.scss']
})
export class ImportComponent implements OnInit {
  @Input() action: string = ''
  @Input() type: string = '.zip'
  @Input() accept: string = '.zip'

  constructor(private ms: NzMessageService) {
  }

  ngOnInit(): void {
  }

  uploadChange($event: NzUploadChangeParam) {
    console.log($event)
    if ($event.type == "success") {
      if ($event.file.response?.error) {
        this.ms.error($event.file.response.error)
      }
    }

  }
}
