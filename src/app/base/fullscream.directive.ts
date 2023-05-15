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
    @Output() setIndex = new EventEmitter();
    private isDown = false;
    private dragDown = false;
    num = 1;
    shiftPosition = { x: 0, y: 0 };
    element: any = null;

    constructor(private el: ElementRef) {
        this.element = this.el.nativeElement;
    }
    ngOnInit(): void {}

    @HostListener('mousedown', ['$event']) onMousedown(event: any) {
        const elementRect = this.element.getBoundingClientRect();
        const draw =
            !this.dragDown &&
            event.clientX > elementRect.left &&
            event.clientX < elementRect.right - 120 &&
            event.clientY > elementRect.top &&
            event.clientY < elementRect.top + 37;
        const resize =
            !this.isDown &&
            event.clientX > elementRect.right - 10 &&
            event.clientX < elementRect.right &&
            event.clientY > elementRect.bottom - 10 &&
            event.clientY < elementRect.bottom;

        if (draw) {
            this.dragDown = true;
            this.num = 1;
        }
        if (resize) {
            this.isDown = true;

            const mask = document.createElement('div');
            mask.style.cssText =
                'position: absolute;top: 0;left: 0;width: 100vw;height: 100vh;z-index: 9999;';
            mask.setAttribute('id', 'mask');
            document.body.append(mask);
        }
    }

    @HostListener('document: mousemove', ['$event']) onMousemove(event: any) {
        if (this.dragDown && this.num) {
            const mask = document.createElement('div');
            mask.style.cssText =
                'position: absolute;top: 0;left: 0;width: 100vw;height: 100vh;z-index: 9999;';
            mask.setAttribute('id', 'mask');
            document.body.append(mask);
            this.num = 0;
            this.setIndex.emit();
        }
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
        if (this.isDown || this.dragDown) {
            this.isDown = false;
            this.dragDown = false;
            const mask = document.getElementById('mask');
            mask && mask.remove();
        }
    }
}
