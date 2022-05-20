import {AfterViewInit, Directive, ElementRef, Input} from '@angular/core';

@Directive({
  selector: '[app-html]',
})
export class HtmlDirective implements AfterViewInit {

  _html = '[html]'

  @Input("app-html")
  set html(html: string) {
    this._html = html
    this.elementRef.nativeElement.innerHTML = html
  }

  constructor(private elementRef: ElementRef) {

  }

  ngAfterViewInit(): void {
    this.elementRef.nativeElement.innerHTML = this._html
  }

}
