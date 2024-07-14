import {Component, Input, OnInit} from '@angular/core';
import {NzMessageService} from 'ng-zorro-antd/message';
import {NzModalService} from 'ng-zorro-antd/modal';
import {SmartRequestService} from '../../../request.service';

@Component({
    selector: 'app-object-attach',
    standalone: true,
    imports: [],
    templateUrl: './object-attach.component.html',
    styleUrl: './object-attach.component.scss'
})
export class ObjectAttachComponent implements OnInit {
    constructor(
        private rs: SmartRequestService,
        private msg: NzMessageService,
        private ms: NzModalService
    ) {
    }

    ngOnInit(): void {
        throw new Error('Method not implemented.');
    }

    total = 0;
    pageIndex = 1;
    pageSize = 10;
    value = '';
    item!: any
    @Input() id: string = '';
    @Input() type: string = '';

    load() {
        let query;

        query = {
            limit: this.pageSize,
            skip: (this.pageIndex - 1) * this.pageSize,
        };

        this.value ? (query = {...query, filter: {id: this.value}}) : '';
        this.rs.post(`${this.type}/${this.id}/attach`, query).subscribe(
            (res) => {

                this.item = res.data;
                this.total = res.total;

            },
            (err) => {
                console.log('err:', err);
            }
        );
    }
}
