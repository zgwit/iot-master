import {Component, Input, OnInit} from '@angular/core';
import {NzUploadChangeParam} from "ng-zorro-antd/upload";

@Component({
  selector: 'app-import',
  templateUrl: './import.component.html',
  styleUrls: ['./import.component.scss']
})
export class ImportComponent implements OnInit {
  @Input() action: string = ''

  constructor() {
  }

  ngOnInit(): void {
  }

  uploadChange($event: NzUploadChangeParam) {
    console.log($event)

  }
}
