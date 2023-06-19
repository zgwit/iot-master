import { Component, Input } from '@angular/core';
import { DomSanitizer, SafeResourceUrl } from '@angular/platform-browser';
import { AppService } from '../app.service';
@Component({
    selector: 'app-window',
    templateUrl: './window.component.html',
    styleUrls: ['./window.component.scss'],
})
export class WindowComponent {
    tabData: any; 
    constructor(private san: DomSanitizer, private app: AppService) {
        let oem: any = localStorage.getItem('window');
        this.tabData = JSON.parse(oem).entries;
        this.tabData.forEach((item: { url: SafeResourceUrl; path: string }) => {
            item.url = this.san.bypassSecurityTrustResourceUrl(item.path);
        });
        //this._url = san.bypassSecurityTrustResourceUrl("http://image.baidu.com")
    }
}
