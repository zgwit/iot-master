import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';

@Component({
  selector: 'app-common-bar',
  templateUrl: './common-bar.component.html',
  styleUrls: ['./common-bar.component.scss']
})
export class CommonBarComponent implements OnInit {
  @Output() search = new EventEmitter<string>();
  @Output() refresh = new EventEmitter();
  @Output() create = new EventEmitter();

  @Input() loading: any = false;
  @Input() noCreate: any = false;
  @Input() createButtonText = "新建";
  @Input() noSearch: any = false;
  @Input() noView: any = false;

  //视图
  @Input() view: string = 'card';
  @Output() viewChange = new EventEmitter<string>();

  keyword = '';

  constructor() {
  }

  ngOnInit(): void {
  }

  onViewChange($event: any) {
    //console.log("onViewChange($event)", $event)
    this.viewChange.emit($event)
  }
}
