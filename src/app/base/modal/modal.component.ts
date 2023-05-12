import { Component, Input, OnInit, Output, EventEmitter } from '@angular/core';
import { NzMessageService } from 'ng-zorro-antd/message';
import { DomSanitizer, SafeResourceUrl } from '@angular/platform-browser';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';
@Component({
    selector: 'app-modal',
    templateUrl: './modal.component.html',
    styleUrls: ['./modal.component.scss'],
})
export class ModalComponent implements OnInit {
    @Input() title: any;
    @Input() show: any;
    width: any = '60vw';
    height: any = '50vh';
    dynamic = false;
    items: any[] = [];
    @Output() close = new EventEmitter();
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
        this.close.emit(this.title); 
        this.width = '60vw';
        this.height = '50vh';
    }
    // close() {
    //     this.items.filter((item: any, index: any) => {
    //         if (item === this.title) this.items.splice(index, 1);
    //     });
    // }
    addTab() {
       // this.items.push(this.title);
        this.hide.emit(this.title);
    }
    showTab() {}
    fullscrean() {this.dynamic=!this.dynamic
       
        if (this.width === '100vw') {
            this.dynamic = false;
            this.width = '60vw';
            this.height = '50vh';
        } else {
            this.dynamic = true;
            this.width = '100vw';
            this.height = '100vh';
        }
    }
}
