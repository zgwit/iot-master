import {AfterViewInit, Component, HostListener, OnInit} from '@angular/core';
import {
  Circle,
  Container,
  Ellipse,
  ForeignObject,
  Image,
  Line,
  Polygon,
  Polyline,
  Rect,
  Svg,
  SVG,
  Text
} from '@svgdotjs/svg.js';
import '@svgdotjs/svg.draggable.js'
import {GetComponent, GroupedComponents} from "../components";
import {
  CreateComponentObject,
  GetPropertiesDefault,
  HmiComponent,
  HmiEntity,
  HmiPropertyItem,
  HmiElement, GetComponentAllProperties
} from "../hmi";
import {CreateElement} from "../create";

@Component({
  selector: 'hmi-editor',
  templateUrl: './editor.component.html',
  styleUrls: ['./editor.component.scss']
})
export class EditorComponent implements OnInit, AfterViewInit {
  // @ts-ignore
  canvas: Svg;

  // @ts-ignore
  baseLayer: Container;

  // @ts-ignore
  grid: Rect;

  // @ts-ignore
  mainLayer: Container;

  // @ts-ignore
  editLayer: Container;

  currentComponent: HmiComponent | undefined = undefined;

  groupedComponents = GroupedComponents

  entities: Array<HmiEntity> = []
  current: HmiEntity | undefined;

  properties: any = {}
  $properties: Array<HmiPropertyItem> = []

  color = "#FFFFFF"
  fill = "#FFFFFF"
  stroke = 2

  constructor() {
  }

  @HostListener("document:keyup", ['$event'])
  onKeyup(event: KeyboardEvent) {
    console.log("onKeyup", event)
    switch (event.key) {
      case "Delete": //删除
        this.current?.$element.remove();
        this.editLayer.clear();
        let index = this.entities.findIndex(v => v == this.current)
        this.current = undefined
        if (index > -1)
          this.entities.splice(index, 1)
        break;
      case "s": //保存
        if (event.ctrlKey) {

        }
        break;
      case "c": //复制
        if (event.ctrlKey) {

        }
        break;
      //case "v": break;
      case "ArrowUp":
        if (this.current) {
          this.current.$element.dmove(0, -5)
          this.current.properties.y -= 5
          this.editEntity(this.current)
        }
        break;
      case "ArrowDown":
        if (this.current) {
          this.current.$element.dmove(0, 5)
          this.current.properties.y += 5
          this.editEntity(this.current)
        }
        break;
      case "ArrowLeft":
        if (this.current) {
          this.current.$element.dmove(-5, 0)
          this.current.properties.x -= 5
          this.editEntity(this.current)
        }
        break;
      case "ArrowRight":
        if (this.current) {
          this.current.$element.dmove(5, 0)
          this.current.properties.x += 5
          this.editEntity(this.current)
        }
        break;
    }
  }

  ngOnInit(): void {

  }

  initGrid(): void {
    //网格线
    let gridSize = 10
    let gridColor = "#202020"
    let pattern = this.baseLayer.pattern(gridSize, gridSize, pattern => {
      pattern.line(0, 0, gridSize, 0).stroke(gridColor)
      pattern.line(0, 0, 0, gridSize).stroke(gridColor)
    })
    this.grid = this.baseLayer.rect().size("100%", "100%").fill(pattern).stroke(gridColor);
  }

  ngAfterViewInit(): void {
    // @ts-ignore
    this.canvas = SVG().addTo('#hmi-editor-canvas').size(800, 600);
    this.baseLayer = this.canvas.group();
    this.initGrid();

    this.mainLayer = this.canvas.group();
    this.editLayer = this.canvas.group();

    this.entities.forEach(entity => {
      let cmp = GetComponent(entity.component)
      if (!cmp) return
      entity.$element = CreateElement(this.canvas, cmp)
      entity.$object = cmp.data ? cmp.data() : {}
      entity.$object.__proto__ = {
        $element: entity.$element,
        $component: cmp,
        $properties: entity.properties,
      }
      cmp.create?.call(entity.$object, entity.properties)
      cmp.setup.call(entity.$object, entity.properties)

      this.makeEntityEditable(entity);
    })
  }

  makeEntityEditable(entity: HmiEntity) {
    let element = entity.$element;
    // @ts-ignore
    element.draggable().on('dragmove', (e) => {
      //console.log("move", e)
      this.onMove(entity)
    });

    element.on('dragstart', e => {
      this.editLayer.clear()
    })

    element.on('dragend', e => {
      this.editEntity(entity)
    })

    element.on('click', (e) => {
      if (this.current == entity) {
        //TODO 取消编辑
        return
      }
      this.current = entity
      //this.editLayer.clear()
      this.editEntity(entity)
    })

  }

  draw(cmp: HmiComponent) {
    //清空编辑区
    this.editLayer.clear()
    //清空选择
    this.current = undefined;

    this.currentComponent = cmp;

    this.$properties = GetComponentAllProperties(cmp);
    this.properties = GetPropertiesDefault(cmp)

    let element = CreateElement(this.mainLayer, cmp)
    //使用默认填充色 线框
    if (cmp.color) {
      this.properties.fill = this.fill
      element.fill(this.fill)
    }
    if (cmp.stroke) {
      this.properties.color = this.color
      this.properties.stroke = this.stroke
      element.stroke({color: this.color, width: this.stroke})
    }

    let entity: HmiEntity = {
      name: "",
      component: cmp.uuid,
      properties: this.properties,
      handlers: {},
      bindings: {},

      $element: element,
      $component: cmp,
      $object: cmp.data ? cmp.data() : {},
    }
    entity.$object.__proto__ = {
      $element: element,
      $component: cmp,
      $properties: this.properties,
    }

    // @ts-ignore
    entity.__proto__ = {
      //$emit: ()=>{}
    }


    //this.entities.push(entity)

    //画
    this.drawEntity(entity);

    this.makeEntityEditable(entity);

    //组件初始化
    cmp.create?.call(entity.$object, entity.properties)
    cmp.setup.call(entity.$object, entity.properties)
  }


  StopDraw() {
    this.canvas.off('click.draw')
    this.canvas.off('mousemove.draw')
  }

  getLinePoints(line: Line | Polyline | Polygon, properties: any) {
    properties.points = line.plot().toArray()
  }

  getRectPosition(rect: Rect | Ellipse | Image | Svg | ForeignObject, properties: any) {
    properties.x = rect.x()
    properties.y = rect.y()
    properties.width = rect.width()
    properties.height = rect.height()
  }

  getCirclePosition(circle: Circle, properties: any) {
    properties.x = circle.cx()
    properties.y = circle.cy()
    // @ts-ignore
    properties.radius = circle.width() / 2;
  }

  onMove(entity: HmiEntity) {
    const type = entity.$component.type || "svg"
    switch (type) {
      case "rect" :
      case "image" :
      case "text" :
      case "svg" :
      case "object":
      case "ellipse" :
        // @ts-ignore
        this.getRectPosition(entity.$element, entity.properties)
        break
      case "circle" :
        // @ts-ignore
        this.getCirclePosition(entity.$element, entity.properties)
        break
      case "line" :
      case "polyline" :
      case "polygon" :
        // @ts-ignore
        this.getLinePoints(entity.$element, entity.properties)
        break
      case "path" :
      default:
        throw new Error("不支持的控件类型：" + type)
    }
  }

  //drawLine(line: Line, properties: any) {
  drawLine(entity: HmiEntity) {
    // @ts-ignore
    let line: Line = entity.$element

    let startX = 0;
    let startY = 0;
    let firstClick = true;

    // @ts-ignore
    this.canvas.on('click.draw', (e: MouseEvent) => {
      if (firstClick) {
        firstClick = false
        startX = e.offsetX
        startY = e.offsetY
        line.addTo(this.mainLayer)
        this.entities.push(entity)

        // @ts-ignore
        this.canvas.on('mousemove.draw', (e: MouseEvent) => {
          line.plot(startX, startY, e.offsetX, e.offsetY)
        })
      } else {
        this.StopDraw()
        //properties.points = line.plot().toArray()
        this.getLinePoints(line, entity.properties)
      }
    });
  }

  //drawRect(rect: Rect | Ellipse | Image | Svg | ForeignObject, properties: any) {
  drawRect(entity: HmiEntity) {
    // @ts-ignore
    let rect: Rect | Ellipse | Image | Svg | ForeignObject = entity.$element

    //let rect: Rect;
    let startX = 0;
    let startY = 0;
    let firstClick = true;

    let outline = new Rect();

    // @ts-ignore
    this.canvas.on('click.draw', (e: MouseEvent) => {
      if (firstClick) {
        firstClick = false
        startX = e.offsetX
        startY = e.offsetY
        rect.addTo(this.mainLayer).move(startX, startY)
        this.entities.push(entity)

        outline.addTo(this.editLayer).move(startX, startY).stroke({
          width: 1,
          color: '#7be',
          dasharray: "6 2",
          dashoffset: 8
        }).fill("none")
        // @ts-ignore
        outline.animate().ease('-').stroke({dashoffset: 0}).loop();

        // @ts-ignore
        this.canvas.on('mousemove.draw', (e: MouseEvent) => {
          let width = e.offsetX - startX;
          let height = e.offsetY - startY;
          if (width > 0 && height > 0) {
            rect.size(width, height)
            outline.size(width, height)
            //entity.$component.setup.call(entity.$object, {width, height})
            entity.$component.resize?.call(entity.$object)
          }
        })
      } else {
        outline.remove()
        this.StopDraw()
        this.getRectPosition(rect, entity.properties);
      }
    });
  }

  //drawCircle(circle: Circle, properties: any) {
  drawCircle(entity: HmiEntity) {
    // @ts-ignore
    let circle: Circle = entity.$element

    let startX = 0;
    let startY = 0;
    let firstClick = true;
    let radius = 0

    // @ts-ignore
    this.canvas.on('click.draw', (e: MouseEvent) => {
      if (firstClick) {
        firstClick = false
        startX = e.offsetX
        startY = e.offsetY
        circle.addTo(this.mainLayer).center(startX, startY)
        this.entities.push(entity)

        // @ts-ignore
        this.canvas.on('mousemove.draw', (e: MouseEvent) => {
          let width = e.offsetX - startX;
          let height = e.offsetY - startY;
          radius = Math.sqrt(width * width + height * height)
          circle.radius(radius)
        })
      } else {
        this.StopDraw()
        this.getCirclePosition(circle, entity.properties)
        //properties.radius = radius
      }
    });
  }

  //drawEllipse(ellipse: Ellipse, properties: any) {
  drawEllipse(entity: HmiEntity) {
    // @ts-ignore
    let ellipse: Ellipse = entity.$element

    let startX = 0;
    let startY = 0;
    let firstClick = true;

    let outline = new Rect();

    // @ts-ignore
    this.canvas.on('click.draw', (e: MouseEvent) => {
      if (firstClick) {
        firstClick = false
        startX = e.offsetX
        startY = e.offsetY
        ellipse.addTo(this.mainLayer).move(startX, startY)
        this.entities.push(entity)

        outline.addTo(this.editLayer).move(startX, startY).stroke({
          width: 1,
          color: '#7be',
          dasharray: "6 2",
          dashoffset: 8
        }).fill("none")
        // @ts-ignore
        outline.animate().ease('-').stroke({dashoffset: 0}).loop();

        // @ts-ignore
        this.canvas.on('mousemove.draw', (e: MouseEvent) => {
          let width = e.offsetX - startX
          let height = e.offsetY - startY
          if (width > 0 && height > 0) {
            ellipse.center(startX + width / 2, startY + height / 2).size(width, height)
            outline.size(width, height)
          }
        })
      } else {
        outline.remove()
        this.StopDraw()
        this.getRectPosition(ellipse, entity.properties)
      }
    });
  }

  //drawPoly(poly: Polygon | Polyline, properties: any) {
  drawPoly(entity: HmiEntity) {
    // @ts-ignore
    let poly: Polygon | Polyline = entity.$element

    let firstClick = true;

    // @ts-ignore
    this.canvas.on('click.draw', (e: MouseEvent) => {
      if (firstClick) {
        firstClick = false
        poly.addTo(this.mainLayer).plot([e.offsetX, e.offsetY, e.offsetX, e.offsetY])
        this.entities.push(entity)

        // @ts-ignore
        this.canvas.on('mousemove.draw', (e: MouseEvent) => {
          let arr = poly.array()
          let pt = arr[arr.length - 1]
          pt[0] = e.offsetX
          pt[1] = e.offsetY
          poly.plot(arr)
        })

        let that = this;
        //TODO on Esc:cancel
        document.addEventListener('keydown', function onKeydown(e) {
          if (e.key == 'Escape') {
            let arr = poly.array()
            arr.pop() //删除最后一个
            poly.plot(arr)

            //line.draw('done');
            that.StopDraw()
            //properties.points = arr.toArray()
            that.getLinePoints(poly, entity.properties)

            //off listener
            document.removeEventListener('keydown', onKeydown)
          }
        });

      } else {
        let arr = poly.array()
        arr.pop() //删除最后一个
        arr.push([e.offsetX, e.offsetY], [e.offsetX, e.offsetY])
        poly.plot(arr)
      }
    });
  }

  drawEntity(entity: HmiEntity): void {
    this.StopDraw()

    // let elem = CreateElement(container, component)
    const type = entity.$component.type || "svg"
    switch (type) {
      case "rect" :
      case "image" :
      case "text" :
      case "svg" :
      case "object":
        // @ts-ignore
        this.drawRect(entity)
        break
      case "circle" :
        // @ts-ignore
        this.drawCircle(entity)
        break
      case "ellipse" :
        // @ts-ignore
        this.drawEllipse(entity)
        break
      case "line" :
        // @ts-ignore
        this.drawLine(entity)
        break
      case "polyline" :
      case "polygon" :
        // @ts-ignore
        this.drawPoly(entity)
        break
      case "path" :
      default:
        throw new Error("不支持的控件类型：" + type)
    }
  }

  //editLine(element: Line | Polygon | Polyline, properties: any) {
  editLine(entity: HmiEntity) {
    // @ts-ignore
    let element: Line | Polygon | Polyline = entity.$element
    let points = element.array() //.toArray()
    points.forEach((p, i) => {
      let pt = this.editLayer.circle(8).fill('#7be').center(p[0], p[1]).css('cursor', 'pointer').draggable();
      pt.on("dragmove", () => {
        p[0] = pt.cx()
        p[1] = pt.cy()
        element.plot(points)
        this.getLinePoints(element, entity.properties)
      })
    })
  }

  //editRect(element: Rect | Ellipse | Text | Image | Svg | ForeignObject, properties: any) {
  editRect(entity: HmiEntity) {
    // @ts-ignore
    let element: Rect | Ellipse | Text | Image | Svg | ForeignObject = entity.$element
    let obj = element.attr()
    let border = this.editLayer.rect(obj.width, obj.height).move(obj.x, obj.y).fill('none').stroke({
      width: 1,
      color: '#7be',
      dasharray: "6 2",
      dashoffset: 0
    });
    // @ts-ignore
    border.animate().ease('-').stroke({dashoffset: 8}).loop();

    let lt = this.editLayer.rect(8, 8).fill('#7be').center(obj.x, obj.y).css('cursor', 'nw-resize').draggable();
    let lm = this.editLayer.rect(8, 8).fill('#7be').center(obj.x, obj.y + obj.height * 0.5).css('cursor', 'w-resize').draggable();
    let lb = this.editLayer.rect(8, 8).fill('#7be').center(obj.x, obj.y + obj.height).css('cursor', 'sw-resize').draggable();
    let rt = this.editLayer.rect(8, 8).fill('#7be').center(obj.x + obj.width, obj.y).css('cursor', 'ne-resize').draggable();
    let rm = this.editLayer.rect(8, 8).fill('#7be').center(obj.x + obj.width, obj.y + obj.height * 0.5).css('cursor', 'e-resize').draggable();
    let rb = this.editLayer.rect(8, 8).fill('#7be').center(obj.x + obj.width, obj.y + obj.height).css('cursor', 'se-resize').draggable();
    let t = this.editLayer.rect(8, 8).fill('#7be').center(obj.x + obj.width * 0.5, obj.y).css('cursor', 'n-resize').draggable();
    let b = this.editLayer.rect(8, 8).fill('#7be').center(obj.x + obj.width * 0.5, obj.y + obj.height).css('cursor', 's-resize').draggable();

    lt.on("dragmove", () => {
      // @ts-ignore
      if (lt.cx() > element.width() + element.x()) return
      // @ts-ignore
      if (lt.cy() > element.height() + element.y()) return
      // @ts-ignore
      element.width(element.width() - lt.cx() + element.x())
      // @ts-ignore
      element.height(element.height() - lt.cy() + element.y())
      element.x(lt.cx())
      element.y(lt.cy())
      update()
    })
    lm.on("dragmove", () => {
      // @ts-ignore
      if (lm.cx() > element.width() + element.x()) return
      // @ts-ignore
      element.width(element.width() - lm.cx() + element.x())
      element.x(lm.cx())
      update()
    })
    lb.on("dragmove", () => {
      // @ts-ignore
      if (lb.cx() > element.width() + element.x()) return
      // @ts-ignore
      if (lb.cy() < element.y()) return
      // @ts-ignore
      element.width(element.width() - lb.cx() + element.x())
      // @ts-ignore
      element.height(lb.cy() - element.y())
      element.x(lb.cx())
      update()
    })

    rt.on("dragmove", () => {
      // @ts-ignore
      if (rt.cx() < element.x()) return
      // @ts-ignore
      if (rt.cy() > element.height() + element.y()) return
      // @ts-ignore
      element.width(rt.cx() - element.x())
      // @ts-ignore
      element.height(element.height() - rt.cy() + element.y())
      element.y(rt.cy())
      update()
    })
    rm.on("dragmove", () => {
      // @ts-ignore
      if (rm.cx() < element.x()) return
      // @ts-ignore
      element.width(rm.cx() - element.x())
      update()
    })
    rb.on("dragmove", () => {
      // @ts-ignore
      if (rb.cx() < element.x()) return
      // @ts-ignore
      if (rb.cy() < element.y()) return
      // @ts-ignore
      element.width(rb.cx() - element.x())
      // @ts-ignore
      element.height(rb.cy() - element.y())
      update()
    })
    t.on("dragmove", () => {
      // @ts-ignore
      if (t.cy() > element.y() + element.height()) return
      // @ts-ignore
      element.height(element.height() - t.cy() + element.y())
      element.y(t.cy())
      update()
    })
    b.on("dragmove", () => {
      // @ts-ignore
      if (b.cy() < element.y()) return
      // @ts-ignore
      element.height(b.cy() - element.y())
      update()
    })

    let that = this

    function update() {
      that.getRectPosition(element, entity.properties)

      let obj = element.attr()
      border.size(obj.width, obj.height).move(obj.x, obj.y)
      //entity.$component.setup?.call(entity.$object, {width: obj.width, height: obj.height})
      entity.$component.resize?.call(entity.$object)

      //border.attr(obj)
      lt.center(obj.x, obj.y)
      lm.center(obj.x, obj.y + obj.height * 0.5)
      lb.center(obj.x, obj.y + obj.height)
      rt.center(obj.x + obj.width, obj.y)
      rm.center(obj.x + obj.width, obj.y + obj.height * 0.5)
      rb.center(obj.x + obj.width, obj.y + obj.height)
      t.center(obj.x + obj.width * 0.5, obj.y)
      b.center(obj.x + obj.width * 0.5, obj.y + obj.height)
    }
  }

  //editCircle(element: Circle, properties: any) {
  editCircle(entity: HmiEntity) {
    // @ts-ignore
    let element: Circle = entity.$element
    // @ts-ignore
    let x = element.cx() + element.width() / 2 // / Math.sqrt(2)
    // @ts-ignore
    let y = element.cy() // + element.width() / 2 / Math.sqrt(2)

    let pt = this.editLayer.circle(8).fill('#7be').center(x, y).css('cursor', 'pointer').draggable();
    pt.on("dragmove", () => {

      let width = pt.cx() - element.cx();
      let height = pt.cy() - element.cy();
      let radius = Math.sqrt(width * width + height * height)
      element.radius(radius)
      this.getCirclePosition(element, entity.properties)
    })
  }

  editEntity(entity: HmiEntity) {
    this.editLayer.clear()
    this.$properties = GetComponentAllProperties(entity.$component);
    this.properties = entity.properties
    const type = entity.$component.type || "svg"
    switch (type) {
      case "ellipse" :
        break
      case "rect" :
      case "image" :
      case "text" :
      case "svg" :
      case "object":
        // @ts-ignore
        this.editRect(entity)
        break
      case "circle" :
        // @ts-ignore
        this.editCircle(entity)
        break
      case "line" :
      case "polyline" :
      case "polygon" :
        // @ts-ignore
        this.editLine(entity)
        break
      case "path" :
      default:
        throw new Error("不支持的控件类型：" + type)
    }
  }

  onPropertyChange(name: string) {
    if (this.current) {
      let val = this.properties[name];
      let cmp = this.current.$component;
      let elem = this.current.$element;
      // @ts-ignore
      let index = cmp.properties.findIndex(p => p.name == name)
      if (index > -1) {
        let props = {[name]: val}
        cmp.setup.call(this.current.$object, props)
        return
      }

      switch (name) {
        case "fill":
          elem.fill(this.properties.fill)
          break;
        case "stroke":
        case "color":
          elem.stroke({color: this.properties.color, width: this.properties.stroke})
          break;
        case "x":
        case "y":
          elem.move(this.properties.x, this.properties.y)
          this.editEntity(this.current)
          break;
        case "width":
        case "height":
          elem.size(this.properties.width, this.properties.height)
          cmp.resize?.call(this.current.$object)
          this.editEntity(this.current)
          break;
        case "rotate":
          elem.transform({rotate: this.properties.rotate})
          this.editEntity(this.current)
          break;
      }
    }
    //entity.$component.setup?.call(entity.$object, {width: obj.width, height: obj.height})
  }

}
