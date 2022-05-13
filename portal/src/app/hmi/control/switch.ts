import {HmiComponent} from "../hmi";

export let SwitchComponent: HmiComponent = {
  uuid: "switch",
  name: "开关",
  icon: "/assets/hmi/switch.svg",
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
  create(props: any) {
    // @ts-ignore
    this.rect = this.$element.rect().size("100%", "100%")
    // @ts-ignore
    let box = this.rect.bbox()
    // @ts-ignore
    this.rect.radius(box.height / 2)

    // @ts-ignore
    this.cell = this.$element.rect().size("40%", "80%")
    // @ts-ignore
    let box2 = this.cell.bbox()
    // @ts-ignore
    this.cell.radius(box2.height / 2).move(0.1*box.height, 0.1*box.height)
  },

  resize() {
    // @ts-ignore
    let box = this.rect.bbox()
    // @ts-ignore
    this.rect.radius(box.height / 2)
    // @ts-ignore
    let box2 = this.cell.bbox()
    // @ts-ignore
    this.cell.radius(box2.height / 2).move(0.1*box.height, 0.1*box.height)
  },

  //配置
  setup(props: any) {
    if (props.color) { // @ts-ignore
      this.cell.fill(props.color)
    }
    if (props.fill) { // @ts-ignore
      this.rect.fill(props.fill)
    }
  },
}
