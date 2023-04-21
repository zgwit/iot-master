import { Component, Input } from '@angular/core';
import { DomSanitizer, SafeResourceUrl } from "@angular/platform-browser";

@Component({
  selector: 'app-window',
  templateUrl: './window.component.html',
  styleUrls: ['./window.component.scss']
})
export class WindowComponent {
  tabData: any;

  @Input() set entries(arr: any) {
    arr.forEach((item: { url: SafeResourceUrl; path: string; }) => {
      item.url = this.san.bypassSecurityTrustResourceUrl(item.path);
    });
    this.tabData = arr;
  }

  constructor(private san: DomSanitizer) {
    //this._url = san.bypassSecurityTrustResourceUrl("http://image.baidu.com")
  }
}

