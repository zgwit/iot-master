import {AfterViewInit, Component, HostListener, Input, OnInit, Output, EventEmitter} from '@angular/core';
import {
  Container,
  Rect,
  Svg,
  SVG,
  G
} from '@svgdotjs/svg.js';
import '@svgdotjs/svg.draggable.js'
import {GetComponent, GroupedComponents} from "../components";
import {
  GetPropertiesDefault,
  HmiComponent,
  HmiEntity,
  HmiPropertyItem,
  GetComponentAllProperties, CreateEntityObject, HmiEvent
} from "../hmi";
import {AttachmentComponent} from "../attachment/attachment.component";
import {NzModalService} from "ng-zorro-antd/modal";

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

  @Input() entities: Array<HmiEntity> = []
  current?: HmiEntity;

  plate?: HmiEntity;

  properties: any = {}
  $properties: Array<HmiPropertyItem> = []

  color = "#FFFFFF"
  fill = "#cccccc"
  stroke = 2

  width = 800
  height = 600
  scale = 1.0

  @Input() attachment = "/api/attachment/"

  @Output() save = new EventEmitter<any>();

  _hmi: any = {}
  @Input() set hmi(hmi: any) {
    this._hmi = hmi;
    this.changeSize(hmi.width, hmi.height)
    this.entities = hmi.entities
    //绘制
    this.renderEntities()
  }

  constructor(private ms: NzModalService) {
  }

  @HostListener("document:keyup", ['$event'])
  onKeyup(event: KeyboardEvent) {
    console.log("onKeyup", event)
    switch (event.key) {
      case "Delete": //删除
        this.current?.$container.remove();
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
      case "x": //剪切
        if (event.ctrlKey)
          this.onCut()
        break;
      case "c": //复制
        if (event.ctrlKey)
          this.onCopy()
        break;
      case "v":
        if (event.ctrlKey)
          this.onPaste()
        break;
      case "ArrowUp":
        if (this.current) {
          this.current.$container.dmove(0, -5)
          this.current.properties.y -= 5
          this.editEntity(this.current)
        }
        break;
      case "ArrowDown":
        if (this.current) {
          this.current.$container.dmove(0, 5)
          this.current.properties.y += 5
          this.editEntity(this.current)
        }
        break;
      case "ArrowLeft":
        if (this.current) {
          this.current.$container.dmove(-5, 0)
          this.current.properties.x -= 5
          this.editEntity(this.current)
        }
        break;
      case "ArrowRight":
        if (this.current) {
          this.current.$container.dmove(5, 0)
          this.current.properties.x += 5
          this.editEntity(this.current)
        }
        break;
    }
  }

  ngOnInit(): void {

  }

  ngAfterViewInit(): void {
    // @ts-ignore
    this.canvas = SVG().addTo('#hmi-editor-canvas').size(this.width, this.height);
    this.baseLayer = this.canvas.group();
    this.createGrid();

    this.mainLayer = this.canvas.group();
    this.editLayer = this.canvas.group();
    //绘制
    this.renderEntities()
  }

  renderEntities() {
    if (!this.canvas) return;
    this.entities.forEach(entity => {
      entity.$container = new G()
      entity.$component = GetComponent(entity.component)
      CreateEntityObject(entity)
      entity.$component.setup.call(entity.$object, entity.properties)

      this.appendEntity(entity);
    })
  }

  createGrid(): void {
    //网格线
    let gridSize = 10
    let gridColor = "#202020"
    let pattern = this.baseLayer.pattern(gridSize, gridSize, pattern => {
      pattern.line(0, 0, gridSize, 0).stroke(gridColor)
      pattern.line(0, 0, 0, gridSize).stroke(gridColor)
    })
    this.grid = this.baseLayer.rect().size("100%", "100%").fill(pattern).stroke(gridColor);
  }

  appendEntity(entity: HmiEntity) {
    // @ts-ignore
    this.mainLayer.add(entity.$container)
    //传入全局配置
    this.setupEntity(entity, entity.properties)

    this.makeEntityEditable(entity);

    //this.entities.push(entity)
  }

  setupEntity(entity: HmiEntity, props: any) {
    Object.assign(entity.properties, props)

    //旋转
    if (props.hasOwnProperty("rotate"))
      entity.$container.transform({rotate: props.rotate})
    if (props.hasOwnProperty("x"))
      entity.$container.x(props.x)
    if (props.hasOwnProperty("y"))
      entity.$container.y(props.y)

    entity.$component.setup.call(entity.$object, props)
  }

  makeEntityEditable(entity: HmiEntity) {
    let container = entity.$container;
    // @ts-ignore
    container.draggable().on('dragmove', (e: CustomEvent) => {
      //console.log("move", e)
      //this.onMove(entity, e)
    });

    let startEvent: CustomEvent;

    // @ts-ignore
    container.on('dragstart', (e: CustomEvent) => {
      //console.log("dragstart", e)
      this.editLayer.clear()
      startEvent = e
    })

    // @ts-ignore
    container.on('dragend', (e: CustomEvent) => {
      //console.log("dragend", e)
      let b = startEvent.detail.box
      let b2 = e.detail.box

      let drawer = entity.$component.drawer || "rect"
      switch (drawer) {
        case "rect" :
        case "circle" :
          let x = entity.properties.x + b2.x - b.x
          let y = entity.properties.y + b2.y - b.y
          this.setupEntity(entity, {x, y})
          break
        case "line" :
          let x1 = entity.properties.x1 + b2.x - b.x
          let y1 = entity.properties.y1 + b2.y - b.y
          let x2 = entity.properties.x2 + b2.x - b.x
          let y2 = entity.properties.y2 + b2.y - b.y
          this.setupEntity(entity, {x1, y1, x2, y2})
          break;
        case "poly" :
          let points: Array<any> = entity.properties.points
          points.forEach(pt => {
            pt[0] += b2.x - b.x
            pt[1] += b2.y - b.y
          })
          this.setupEntity(entity, {points})
          break
        default:
          throw new Error("不支持的控件类型：" + drawer)
      }

      this.editEntity(entity)
    })

    container.on('click', (e) => {
      if (this.current == entity) {
        //TODO 取消编辑
        return
      }
      this.current = entity
      //this.editLayer.clear()
      this.editEntity(entity)
    })

  }

  drawBegin(cmp: HmiComponent) {
    //清空编辑区
    this.editLayer.clear()
    //清空选择
    this.current = undefined;

    this.currentComponent = cmp;

    this.$properties = GetComponentAllProperties(cmp);
    this.properties = GetPropertiesDefault(cmp)

    //使用默认填充色 线框
    if (cmp.color) {
      this.properties.fill = this.fill
    }
    if (cmp.stroke) {
      this.properties.color = this.color
      this.properties.stroke = this.stroke
    }

    let entity: HmiEntity = {
      name: "",
      component: cmp.id,
      properties: this.properties,
      handlers: {},
      bindings: {},

      $container: new G(),
      $component: cmp,
      $object: undefined,
    }

    //TODO 改为返回值
    CreateEntityObject(entity)

    //画
    this.drawEntity(entity);
  }

  drawEnd() {
    this.editLayer.clear()
    this.canvas.off('click.draw')
    this.canvas.off('mousemove.draw')
  }

  drawLine(entity: HmiEntity) {
    let firstClick = true;

    // @ts-ignore
    this.canvas.on('click.draw', (e: MouseEvent) => {
      if (firstClick) {
        firstClick = false
        this.entities.push(entity)
        this.appendEntity(entity)
        this.setupEntity(entity, {x1: e.offsetX, y1: e.offsetY, x2: e.offsetX, y2: e.offsetY})

        // @ts-ignore
        this.canvas.on('mousemove.draw', (e: MouseEvent) => {
          this.setupEntity(entity, {x2: e.offsetX, y2: e.offsetY})
        })
      } else {
        this.drawEnd()
      }
    });
  }

  drawRect(entity: HmiEntity) {
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

        this.appendEntity(entity)
        this.entities.push(entity)
        this.setupEntity(entity, {x: e.offsetX, y: e.offsetY})

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
            outline.size(width, height)
            this.setupEntity(entity, {width, height})
          }
        })
      } else {
        this.drawEnd()
      }
    });
  }

  drawCircle(entity: HmiEntity) {
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
        this.appendEntity(entity)
        this.entities.push(entity)
        this.setupEntity(entity, {x: e.offsetX, y: e.offsetY})

        // @ts-ignore
        this.canvas.on('mousemove.draw', (e: MouseEvent) => {
          let width = e.offsetX - startX;
          let height = e.offsetY - startY;
          radius = Math.sqrt(width * width + height * height)
          this.setupEntity(entity, {radius})
        })
      } else {
        this.drawEnd()
      }
    });
  }

  drawPoly(entity: HmiEntity) {
    let firstClick = true;
    let points: any = [];

    // @ts-ignore
    this.canvas.on('click.draw', (e: MouseEvent) => {
      if (firstClick) {
        firstClick = false
        points.push([e.offsetX, e.offsetY], [e.offsetX, e.offsetY])

        this.appendEntity(entity)
        this.entities.push(entity)
        this.setupEntity(entity, {points})

        // @ts-ignore
        this.canvas.on('mousemove.draw', (e: MouseEvent) => {
          let pt = points[points.length - 1]
          pt[0] = e.offsetX
          pt[1] = e.offsetY
          this.setupEntity(entity, {points})
        })

        let that = this;
        //TODO on Esc:cancel
        document.addEventListener('keydown', function onKeydown(e) {
          if (e.key == 'Escape') {
            points.pop() //删除最后一个
            that.setupEntity(entity, {points})

            that.drawEnd()
            //off listener
            document.removeEventListener('keydown', onKeydown)
          }
        });
      } else {
        points.pop() //删除最后一个
        points.push([e.offsetX, e.offsetY], [e.offsetX, e.offsetY])
        this.setupEntity(entity, {points})
      }
    });
  }

  drawEntity(entity: HmiEntity): void {
    this.drawEnd()

    // let elem = CreateElement(container, component)
    const type = entity.$component.drawer || "rect"
    switch (type) {
      case "rect" :
        // @ts-ignore
        this.drawRect(entity)
        break
      case "circle" :
        // @ts-ignore
        this.drawCircle(entity)
        break
      case "line" :
        // @ts-ignore
        this.drawLine(entity)
        break
      case "poly" :
        // @ts-ignore
        this.drawPoly(entity)
        break
      default:
        throw new Error("不支持的控件类型：" + type)
    }
  }

  editLine(entity: HmiEntity) {
    // @ts-ignore
    let p = entity.properties
    let points: Array<any> = [[p.x1, p.y1], [p.x2, p.y2]];
    points.forEach((p, i) => {
      let pt = this.editLayer.circle(8).fill('#7be').center(p[0], p[1]).css('cursor', 'pointer').draggable();
      pt.on("dragmove", () => {
        p[0] = pt.cx()
        p[1] = pt.cy()
        this.setupEntity(entity, {
          x1: points[0][0],
          y1: points[0][1],
          x2: points[1][0],
          y2: points[1][1],
        })
      })
    })
  }

  editPoly(entity: HmiEntity) {
    // @ts-ignore
    let points: Array<any> = entity.properties.points
    points.forEach((p, i) => {
      let pt = this.editLayer.circle(8).fill('#7be').center(p[0], p[1]).css('cursor', 'pointer').draggable();
      pt.on("dragmove", () => {
        p[0] = pt.cx()
        p[1] = pt.cy()
        this.setupEntity(entity, {points})
      })
    })
  }

  editRect(entity: HmiEntity) {
    // @ts-ignore
    let obj = entity.properties
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
      if (lt.cx() > obj.width + obj.x) return
      // @ts-ignore
      if (lt.cy() > obj.height + obj.y) return
      // @ts-ignore
      obj.width = (obj.width - lt.cx() + obj.x)
      // @ts-ignore
      obj.height = (obj.height - lt.cy() + obj.y)
      obj.x = (lt.cx())
      obj.y = (lt.cy())
      update()
    })
    lm.on("dragmove", () => {
      // @ts-ignore
      if (lm.cx() > obj.width + obj.x) return
      // @ts-ignore
      obj.width = (obj.width - lm.cx() + obj.x)
      obj.x = (lm.cx())
      update()
    })
    lb.on("dragmove", () => {
      // @ts-ignore
      if (lb.cx() > obj.width + obj.x) return
      // @ts-ignore
      if (lb.cy() < obj.y) return
      // @ts-ignore
      obj.width = (obj.width - lb.cx() + obj.x)
      // @ts-ignore
      obj.height = (lb.cy() - obj.y)
      obj.x = (lb.cx())
      update()
    })

    rt.on("dragmove", () => {
      // @ts-ignore
      if (rt.cx() < obj.x) return
      // @ts-ignore
      if (rt.cy() > obj.height + obj.y) return
      // @ts-ignore
      obj.width = (rt.cx() - obj.x)
      // @ts-ignore
      obj.height = (obj.height - rt.cy() + obj.y)
      obj.y = (rt.cy())
      update()
    })
    rm.on("dragmove", () => {
      // @ts-ignore
      if (rm.cx() < obj.x) return
      // @ts-ignore
      obj.width = (rm.cx() - obj.x)
      update()
    })
    rb.on("dragmove", () => {
      // @ts-ignore
      if (rb.cx() < obj.x) return
      // @ts-ignore
      if (rb.cy() < obj.y) return
      // @ts-ignore
      obj.width = (rb.cx() - obj.x)
      // @ts-ignore
      obj.height = (rb.cy() - obj.y)
      update()
    })
    t.on("dragmove", () => {
      // @ts-ignore
      if (t.cy() > obj.y + obj.height) return
      // @ts-ignore
      obj.height = (obj.height - t.cy() + obj.y)
      obj.y = (t.cy())
      update()
    })
    b.on("dragmove", () => {
      // @ts-ignore
      if (b.cy() < obj.y) return
      // @ts-ignore
      obj.height = (b.cy() - obj.y)
      update()
    })

    let that = this

    function update() {
      border.size(obj.width, obj.height).move(obj.x, obj.y)

      that.setupEntity(entity, {
        width: obj.width,
        height: obj.height,
        x: obj.x,
        y: obj.y,
      })

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

  editCircle(entity: HmiEntity) {
    let obj = entity.properties
    let x = obj.x + obj.radius
    let y = obj.y

    let pt = this.editLayer.circle(8).fill('#7be').center(x, y).css('cursor', 'pointer').draggable();
    pt.on("dragmove", () => {

      let width = pt.cx() - obj.x;
      let height = pt.cy() - obj.y;
      let radius = Math.sqrt(width * width + height * height)
      this.setupEntity(entity, {radius})
    })
  }

  editEntity(entity: HmiEntity) {
    this.editLayer.clear()
    this.$properties = GetComponentAllProperties(entity.$component);
    this.properties = entity.properties


    const type = entity.$component.drawer || "rect"
    switch (type) {
      case "rect" :
        // @ts-ignore
        this.editRect(entity)
        break
      case "circle" :
        // @ts-ignore
        this.editCircle(entity)
        break
      case "line" :
        // @ts-ignore
        this.editLine(entity)
        break
      case "poly" :
        // @ts-ignore
        this.editPoly(entity)
        break
      default:
        throw new Error("不支持的控件类型：" + type)
    }
  }

  onPropertyChange(name: string) {
    //console.log("onPropertyChange", name)
    if (this.current) {
      let val = this.properties[name];
      console.log("onPropertyChange", name, val)
      let props = {[name]: val}
      this.setupEntity(this.current, props)
    }
  }

  moveTop() {
    let index = this.entities.findIndex(e => e == this.current)
    if (index < this.entities.length - 1) {
      let es = this.entities.splice(index, 1)
      this.entities.push(...es)
      this.current?.$container.front()
    }
  }

  moveBottom() {
    let index = this.entities.findIndex(e => e == this.current)
    if (index > 0) {
      let es = this.entities.splice(index, 1)
      this.entities.splice(0, 0, ...es)
      this.current?.$container.back()
    }
  }

  moveUp() {
    let index = this.entities.findIndex(e => e == this.current)
    if (index < this.entities.length - 1) {
      let es = this.entities.splice(index, 1)
      this.entities.splice(index + 1, 0, ...es)
      this.current?.$container.forward()
    }
  }

  moveDown() {
    let index = this.entities.findIndex(e => e == this.current)
    if (index > 0) {
      let es = this.entities.splice(index, 1)
      this.entities.splice(index - 1, 0, ...es)
      this.current?.$container.backward()
    }
  }

  onCut() {
    this.onCopy()
    this.current?.$container.remove()
    this.current = undefined
  }

  onCopy() {
    let index = this.entities.findIndex(e => e == this.current)
    if (index > 0) {
      let es = this.entities.splice(index, 1)
      this.plate = this.current
      //this.current?.$container.remove()
      //this.current = undefined
      this.editLayer.clear();
    }
  }

  onPaste() {
    if (this.plate) {
      let entity: HmiEntity = {
        name: this.plate.name,
        component: this.plate.component,
        properties: JSON.parse(JSON.stringify(this.plate.properties)),
        handlers: JSON.parse(JSON.stringify(this.plate.handlers)),
        bindings: JSON.parse(JSON.stringify(this.plate.bindings)),
        $container: new G(),
        $component: this.plate.$component,
        $object: undefined,
      }
      //向右下移动一下
      entity.properties.x += 10
      entity.properties.y += 10
      CreateEntityObject(entity)
      this.appendEntity(entity)
      this.plate = entity //切换成新的
    }

  }

  onDelete() {
    let index = this.entities.findIndex(e => e == this.current)
    if (index > 0) {
      let es = this.entities.splice(index, 1)
    }
    this.current?.$container.remove()
    this.current = undefined
    this.editLayer.clear();
  }

  onSizeChange() {
    this.canvas.size(this.width, this.height)
  }

  onScaleChange() {
    this.canvas.transform({scale: this.scale})
  }

  changeSize(width: number, height: number) {
    this.width = width
    this.height = height
    this.canvas?.size(this.width, this.height)
  }

  bindEvent(event: HmiEvent) {
    console.log('bind', event)
    if (this.current) {
      this.current.handlers[event.name] = this.current.handlers[event.name] || []
      this.current.handlers[event.name].push({})
    }
  }

  onSave() {
    let svg = this.canvas.clone();
    svg.children()[0].remove()
    svg.children()[2].remove()
    //@ts-ignore
    this._hmi.snap = svg.svg() //.flatten()
    this._hmi.entities = this.entities.map(e => {
      return {
        name: e.name,
        component: e.component,
        properties: e.properties,
        handlers: e.handlers,
        bindings: e.bindings,
      }
    })
    this.save.emit(this._hmi)
  }

  browserAttachment(prop: string) {
    const modal = this.ms.create({
      nzTitle: '选择附件',
      nzContent: AttachmentComponent,
      nzWidth: '80%',
      nzComponentParams: {
        url: "/api/hmi/" + this._hmi.id + "/attachment/"
        //TODO 此处应该判断，新建组态时，没有目录（可以先创建UUID）
      },
    });
    modal.afterClose.subscribe(res => {
      console.log('attach', res)
      if (res) {
        this.properties[prop] = res
        this.onPropertyChange(prop)
      }
    })
  }
}
