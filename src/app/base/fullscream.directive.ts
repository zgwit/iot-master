// import { Directive, HostListener } from '@angular/core';
// import screenfull from 'screenfull';
// @Directive({
//   selector: '[appFullscream]',
// })
// export class FullscreamDirective {
//   @HostListener('click') onClick() { 
//     if (screenfull.isEnabled) {
//       if (screenfull.isFullscreen) {
//         screenfull.toggle(); 
//       } else screenfull.toggle();
//     }
//   }
// }


import { Directive, ElementRef, HostListener, Input, OnInit } from '@angular/core'; 
@Directive({
  selector: '[appFullscream]',
})
export class FullscreamDirective implements OnInit {
  private isDown = false;
  private disX = 0;
  private disY = 0;
  private dom: any;
   
  @Input('appDrag') className!: string;
 
  constructor(private el: ElementRef) {
  }
  ngOnInit(): void {
    
  }
  
  @HostListener('drag' ,['event'])
  dragEvent(  ) {
     
  }

  @HostListener('mousedown' ) onMousedown(event:any) { 
     
    // if (!this.isDown) {
    // //  this.dom = this.getParentRecurse(this.el.nativeElement, this.className);
    //   // 移动区域
    //   this.isDown = true;
    //   this.disX = event.clientX - this.dom.offsetLeft;
    //   this.disY = event.clientY - this.dom.offsetTop;
   
    // }
     
  }
  
 
 
  @HostListener('mousemove' , ['$event']) onMousemove(event:any) {
     
    // if (this.isDown) {
    //   const cw = document.documentElement.clientWidth;
    //   const cy = document.documentElement.clientHeight;
    //   const dw = this.dom.offsetWidth;
    //   const dh = this.dom.offsetHeight;
 
    //   let oLeft = event.clientX - this.disX;
    //   let oTop = event.clientY - this.disY;
 
    //   if (oTop < 0) {
    //     oTop = 0;
    //   } else if (oTop > cy - dh) {
    //     oTop = cy - dh;
    //   }
 
    //   if (oLeft < 0) {
    //     oLeft = 0;
    //   } else if (oLeft > cw - dw) {
    //     oLeft = cw - dw;
    //   }
 
    //   this.dom.style.left = oLeft + 'px';
    //   this.dom.style.top = oTop + 'px';
    //   this.dom.style.position = 'fixed';
    //  // this.el.nativeElement.style.cursor = 'move';
    // }
  }
  
  @HostListener('mouseup', ['$event'] ) onMouseup(event:any) {
    // if (this.isDown) {
    //   document.onmousemove = null;
    //   document.onmouseup = null;
    //   this.isDown = false;
    // }
   
  }
 
}