import {HmiComponent} from "../hmi";

export let ProgressComponent: HmiComponent = {
  uuid: "progress",
  name: "开关",
  icon: "/assets/hmi/progress.svg",
  group: "控件",
  drawer: "rect",

  properties: [
    {
      label: '颜色',
      name: 'color',
      type: 'color',
      default: '#8BBB11'
    },
    {
      label: '背景',
      name: 'back',
      type: 'color',
      default: '#666'
    },
    {
      label: '边框色',
      name: 'fill',
      type: 'color',
      default: '#ccc'
    },
    {
      label: '圆角',
      name: 'radius',
      type: 'number',
      default: 20
    },
    {
      label: '边框',
      name: 'stroke',
      type: 'number',
      default: 10
    },
  ],

  //配置
  create() {
    // @ts-ignore
    this.rect = this.$container.rect()
    // @ts-ignore
    this.back = this.$container.rect()
    // @ts-ignore
    this.cell = this.$container.rect()
  },

  resize() {
    // @ts-ignore
    let box = this.$element.bbox()
    // @ts-ignore
    let radius = this.$properties.radius

    // @ts-ignore
    let stroke = this.$properties.stroke


    // @ts-ignore
    this.rect.radius(radius)

    // @ts-ignore
    this.back.radius(radius - stroke).size(box.width - stroke * 2, box.height - stroke * 2).x(stroke).cy(box.cy)

    // @ts-ignore
    this.cell.radius(radius - stroke).size(box.width * 0.6 - stroke * 2, box.height - stroke * 2).x(stroke).cy(box.cy)
  },

  //配置
  setup(props: any) {
    //@ts-ignore
    let p = this.$properties
    if (props.hasOwnProperty("color"))  // @ts-ignore
      this.cell.fill(p.color)
    if (props.hasOwnProperty("back"))  // @ts-ignore
      this.back.fill(p.back)
    if (props.hasOwnProperty("fill"))  // @ts-ignore
      this.rect.fill(p.fill)
    if (props.hasOwnProperty("radius")
      || props.hasOwnProperty("stroke")
      || props.hasOwnProperty("width")
      || props.hasOwnProperty("height")
    ) {
      // @ts-ignore
      this.rect.radius(p.radius).size(p.width, p.height)
      // @ts-ignore
      this.back.radius(p.radius - p.stroke)
        .size(p.width - p.stroke * 2, p.height - p.stroke * 2)
        .x(p.x + p.stroke).cy(p.y + p.height / 2)
      // @ts-ignore
      this.cell.radius(p.radius - p.stroke).size(p.width * 0.6 - p.stroke * 2, p.height - p.stroke * 2)
        .x(p.x + p.stroke).cy(p.y + p.height / 2)
    }
    if (props.hasOwnProperty("x") || props.hasOwnProperty("y")) {
      // @ts-ignore
      this.rect.move(p.x, p.y)
      // @ts-ignore
      this.back.x(p.x + p.stroke).cy(p.y + p.height / 2)
      // @ts-ignore
      this.cell.x(p.x + p.stroke).cy(p.y + p.height / 2)
    }
  },

  //更新数据
  update(values: any) {

  },

};
