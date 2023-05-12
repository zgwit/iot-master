import {
    Directive,
    ElementRef,
    EventEmitter,
    HostListener,
    Input,
    OnInit,
    Output,
} from '@angular/core';
@Directive({
    selector: '[appFullscream]',
})
export class FullscreamDirective implements OnInit {
    @Output() mes = new EventEmitter();
    private isDown = false;
    shiftPosition = { x: 0, y: 0 };
    element: any = null;

    constructor(private el: ElementRef) {
        this.element = this.el.nativeElement;
    }
    ngOnInit(): void {}

    @HostListener('document:dblclick', ['$event']) ondblClick(event: any) {
        const elementRect = this.element.getBoundingClientRect();
        const x = event.clientX;
        const y = event.clientY;
        if (
            x < elementRect.left - 10 ||
            x > elementRect.right + 10 ||
            y < elementRect.top - 10 ||
            y > elementRect.bottom + 10
        )
            this.mes.emit();
    }

    @HostListener('mousedown', ['$event']) onMousedown(event: any) {
        const elementRect = this.element.getBoundingClientRect();

        if (
            !this.isDown &&
            event.clientX > elementRect.right - 10 &&
            event.clientX < elementRect.right &&
            event.clientY > elementRect.bottom - 10 &&
            event.clientY < elementRect.bottom
        ) {
            this.isDown = true;

            const mask = document.createElement('div');
            mask.style.cssText =
                'position: absolute;top: 0;left: 0;width: 100vw;height: 100vh;z-index: 9999;';
            mask.setAttribute('id', 'mask');
            document.body.append(mask);
        }
    }

    @HostListener('document: mousemove', ['$event']) onMousemove(event: any) {
        //console.log(1)
        if (this.isDown) {
            const elementRect = this.element.getBoundingClientRect();

            this.shiftPosition.x = elementRect.right - event.clientX;

            this.shiftPosition.y = elementRect.bottom - event.clientY;

            this.element.style.width =
                elementRect.width - this.shiftPosition.x + 'px';
            this.element.style.height =
                elementRect.height - this.shiftPosition.y + 'px';
        }
    }

    @HostListener('document:mouseup', ['$event']) onMouseup(event: any) {
        if (this.isDown) {
            this.isDown = false;
            const mask = document.getElementById('mask');
            mask && mask.remove();
        }
    }
}
