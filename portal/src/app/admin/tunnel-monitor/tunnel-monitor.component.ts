import {Component, ElementRef, Input, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {RequestService} from "../../request.service";
import {ActivatedRoute, Router} from "@angular/router";
import {DomSanitizer, SafeUrl} from "@angular/platform-browser";
import {environment} from "../../../environments/environment";

@Component({
  selector: 'app-tunnel-monitor',
  templateUrl: './tunnel-monitor.component.html',
  styleUrls: ['./tunnel-monitor.component.scss']
})
export class TunnelMonitorComponent implements OnInit, OnDestroy {
  id = '';

  @ViewChild('read') read: ElementRef | undefined;

  ws: any;

  text = '';
  isHex = false;
  crlf = false;
  transfer: SafeUrl | undefined;

  constructor(private router: ActivatedRoute, private rs: RequestService, private san: DomSanitizer) {
    this.id = router.snapshot.params['id'];
  }

  watch() {
    const types = {
      read: '收到',
      write: '发送',
      error: '错误',
      close: '下线',
      data: '下线',
    }

    //此处Angular框架的proxy.conf.json开发模式下不能正常使用，需要替换成原始地址
    const host = environment.production ? location.origin.replace(/^http/, 'ws') : 'ws://localhost:8080';

    this.ws = new WebSocket(`${host}/api/tunnel/${this.id}/watch`);
    //this.transfer = this.san.bypassSecurityTrustUrl(`open-vcom://${host}/api/tunnel/${this._id}/transfer`);

    this.ws.onmessage = (e: any) => {
      console.log('websocket onmessage', e.data)
      const msg: any = JSON.parse(e.data);

      //创建表达DIV元素，避免动态框架的性能损耗
      const div = document.createElement("div");

      const date = new Date();
      const time = document.createElement("span");
      time.append(date.toTimeString().substr(0, 8));
      //time.style.backgroundColor = '#F0F0F0';
      time.style.margin = '5px 10px';

      const type = document.createElement("span");
      // @ts-ignore
      type.append(msg.event)
      type.style.margin = '5px 10px';
      div.append(time, type);

      if (msg.data) {
        const data = document.createElement("span");
        data.append('<', msg.data, '>')
        data.style.margin = '5px 10px';
        div.append(data);
      }

      if (this.read) {
        const container = this.read.nativeElement;
        container.appendChild(div)
        //.scrollIntoView();
        container.scrollTop = container.scrollHeight;

        //TODO 需要考虑最大长度
      }
    }
    this.ws.onerror = (e: any) => {
      console.log('websocket onerror', e)
    }
    this.ws.onclose = (e: any) => {
      console.log('websocket onclose', e)
    }
  }


  ngOnInit(): void {
    console.log('init');
    this.watch()
  }

  ngOnDestroy(): void {
    this.ws?.close();
  }

  send() {
    this.ws?.send(JSON.stringify({
      type: 'write',
      isHex: this.isHex,
      data: this.crlf ? this.text + '\r\n' : this.text
    }));
  }

  clear(){
    // this.read?.nativeElement.childNodes.forEach((c: any)=>{
    //   this.read?.nativeElement.removeChild(c);
    // })

    // @ts-ignore
    this.read?.nativeElement.innerHTML = "";
  }
}
