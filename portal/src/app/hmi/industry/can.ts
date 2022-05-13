import {HmiComponent} from "../hmi";

export let CanComponent: HmiComponent = {
  uuid: "can",
  name: "水罐",
  icon: "/assets/hmi/can.svg",
  group: "工业",
  type: "svg",

  properties: [
    {
      label: '颜色',
      name: 'color',
      type: 'color',
      default: '#ccc'
    },
    {
      label: '背景',
      name: 'back',
      type: 'color',
      default: '#666'
    },
    {
      label: '背景',
      name: 'fill',
      type: 'color',
      default: '#8BBB11'
    },
  ],

  //配置
  create(props: any) {
    // @ts-ignore
    let box = this.$element.bbox();
    let radius = Math.min(50, box.height * 0.5, box.width * 0.5)
    let stroke = radius * 0.2

    // @ts-ignore
    this.rect = this.$element.rect().size("100%", "100%").radius(radius)

    // @ts-ignore
    this.back = this.$element.rect().size(box.width - stroke * 2, box.height - stroke * stroke).radius(radius).cx(box.cx).cy(box.cy)

    // @ts-ignore
    this.cell = this.$element.rect().size(box.width - stroke * 2, box.height - stroke * stroke).radius(radius).cx(box.cx).cy(box.cy)

    // @ts-ignore
    this.clipRect = this.$element.rect().size("100%", "100%").move(0, box.cy)
    // @ts-ignore
    this.clip = this.$element.clip()
    // @ts-ignore
    this.clip.add(this.clipRect)
    // @ts-ignore
    this.cell.clipWith(this.clip)

  },

  resize() {
    // @ts-ignore
    let box = this.$element.bbox();
    let radius = Math.min(50, box.height * 0.5, box.width * 0.5)
    let stroke = radius * 0.2

    // @ts-ignore
    this.back.size(box.width - stroke * 2, box.height - stroke * 2).radius(radius).cx(box.cx).cy(box.cy)

    // @ts-ignore
    this.cell.size(box.width - stroke * 2, box.height - stroke * 2).radius(radius).cx(box.cx).cy(box.cy)

    // @ts-ignore
    this.rect.radius(radius)

    // @ts-ignore
    this.clipRect.move(0, box.cy)
  },

  //配置
  setup(props: any) {
    if (props.color) { // @ts-ignore
      this.cell.fill(props.fill)
    }
    if (props.back) { // @ts-ignore
      this.back.fill(props.back)
    }
    if (props.fill) { // @ts-ignore
      this.rect.fill(props.color)
    }
  },
}
