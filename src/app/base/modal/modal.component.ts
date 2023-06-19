import {
    Component,
    Input,
    OnInit,
    Output,
    EventEmitter,
    ViewChild,
} from '@angular/core';
import { NzMessageService } from 'ng-zorro-antd/message';
import { DomSanitizer, SafeResourceUrl } from '@angular/platform-browser';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';
@Component({
    selector: 'app-modal',
    templateUrl: './modal.component.html',
    styleUrls: ['./modal.component.scss'],
})
export class ModalComponent implements OnInit {
    // index = 0;
    @Input() index: any;
    @Input() title: any;
    @Input() show: any;
    width: any = '60vw';
    height: any = '50vh';
    dragPosition = { x: 0, y: 0 };
    dynamic = false;
    items: any[] = [];
    @Output() close = new EventEmitter();
    @Output() hide = new EventEmitter();
    @Output() setIndex = new EventEmitter();
    constructor(private msg: NzMessageService, private san: DomSanitizer) {}

    ngOnInit(): void {}
    tabData: any;
    zindex() { 
        this.setIndex.emit(this.title);
    }
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

    addTab() {
        this.hide.emit(this.title);
    }
    showTab() {}
    fullscreen() {
        this.dynamic = !this.dynamic;
        this.dragPosition = { x: 0, y: 0 };
        if (this.dynamic) {
            this.width = '100vw';
            this.height = '100vh';
        } else {
            this.width = '60vw';
            this.height = '50vh';
        }
    }
}
