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
    num = 0;
    constructor(private el: ElementRef) {
        this.element = this.el.nativeElement;
    }
    ngOnInit(): void {}

    @HostListener('document:click', ['$event']) onClick(event: any) {
        if (this.num === 1) {
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
        this.num = 1;
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
        }
    }

    @HostListener('document: mousemove', ['$event']) onMousemove(event: any) {
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
        }
    }
}
