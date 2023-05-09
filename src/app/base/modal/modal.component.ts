import { Component, Input, OnInit } from '@angular/core';
import { NzMessageService } from 'ng-zorro-antd/message';
import { DomSanitizer, SafeResourceUrl } from "@angular/platform-browser";
@Component({
    selector: 'app-modal',
    templateUrl: './modal.component.html',
    styleUrls: ['./modal.component.scss'],
})
export class ModalComponent implements OnInit {
    @Input() title: any;
    @Input() show: any;
    constructor(private msg: NzMessageService,private san: DomSanitizer) {}
    
    ngOnInit(): void {}
    tabData: any;

    @Input() set entries(arr: any) {
      arr.forEach((item: { url: SafeResourceUrl; path: string; }) => {
        item.url = this.san.bypassSecurityTrustResourceUrl(item.path);
      });
      this.tabData = arr;
    }
    cancel() {
        this.show = false;
    }
 
}
