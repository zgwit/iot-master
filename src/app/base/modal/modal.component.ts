import { Component, Input, OnInit, Output, EventEmitter } from '@angular/core';
import { NzMessageService } from 'ng-zorro-antd/message';
import { DomSanitizer, SafeResourceUrl } from '@angular/platform-browser';
@Component({
    selector: 'app-modal',
    templateUrl: './modal.component.html',
    styleUrls: ['./modal.component.scss'],
})
export class ModalComponent implements OnInit {
    @Input() title: any;
    @Input() show: any;
    width: any = '95vw';
    height: any = '50vh';
    @Output() hide = new EventEmitter();
    constructor(private msg: NzMessageService, private san: DomSanitizer) {}

    ngOnInit(): void {}
    tabData: any;

    @Input() set entries(arr: any) {
        arr.forEach((item: { url: SafeResourceUrl; path: string }) => {
            item.url = this.san.bypassSecurityTrustResourceUrl(item.path);
        });
        this.tabData = arr;
    }
    cancel() {
        this.hide.emit();
    }

    fullscrean() {
        if (this.width === '100vw') {
            this.width = '95vw';
            this.height = '50vh';
        } else {
            this.width = '100vw';
            this.height = '100vh';
        }
    }
}
