import {HmiComponent} from "../hmi";

export let ProgressComponent: HmiComponent = {
  uuid: "progress",
  name: "开关",
  icon: "/assets/hmi/progress.svg",
  group: "控件",
  type: "svg",

  properties: [
    {
      label: '颜色',
      name: 'color',
      type: 'color',
      default: '#8BBB11'
    },
    {
      label: '阴影',
      name: 'back',
      type: 'color',
      default: '#666'
    },
    {
      label: '背景',
      name: 'fill',
      type: 'color',
      default: '#ccc'
    },
  ],

  //配置
  create(props: any) {
    // @ts-ignore
    let box = this.$element.bbox()
    let radius = Math.min(50, box.height * 0.5, box.width * 0.5)
    let stroke = radius * 0.2

    // @ts-ignore
    this.rect = this.$element.rect().size("100%", "100%")
    // @ts-ignore
    this.rect.radius(radius)

    // @ts-ignore
    this.back = this.$element.rect()
    // @ts-ignore
    this.back.radius(radius).size(box.width - stroke * 2, box.height - stroke * 2).x(stroke).cy(box.cy)

    // @ts-ignore
    this.cell = this.$element.rect()
    // @ts-ignore
    this.cell.radius(radius).size(box.width * 0.6 - stroke * 2, box.height - stroke * 2).x(stroke).cy(box.cy)
  },

  resize() {
    // @ts-ignore
    let box = this.$element.bbox()
    let radius = Math.min(30, box.height * 0.5, box.width * 0.5)
    let stroke = radius * 0.2

    // @ts-ignore
    this.rect.radius(radius)
    // @ts-ignore
    this.back.radius(radius).size(box.width - stroke * 2, box.height - stroke * 2).x(stroke).cy(box.cy)
    // @ts-ignore
    this.cell.radius(radius).size(box.width * 0.6 - stroke * 2, box.height - stroke * 2).x(stroke).cy(box.cy)
  },

  //配置
  setup(props: any) {
    if (props.color) { // @ts-ignore
      this.cell.fill(props.color)
    }
    if (props.back) { // @ts-ignore
      this.back.fill(props.back)
    }
    if (props.fill) { // @ts-ignore
      this.rect.fill(props.fill)
    }
  },
}
