import {HmiComponent} from "../hmi";

export let SwitchComponent: HmiComponent = {
  uuid: "switch",
  name: "开关",
  icon: "/assets/hmi/switch.svg",
  group: "控件",

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
    this.cell = this.$container.circle(0)
  },

  //配置
  setup(props: any) {
    //@ts-ignore
    let p = this.$properties
    // @ts-ignore
    let radius = p.height / 2

    if (props.hasOwnProperty("color"))  // @ts-ignore
      this.cell.fill(p.color)
    if (props.hasOwnProperty("back"))  // @ts-ignore
      this.back.fill(p.back)
    if (props.hasOwnProperty("fill"))  // @ts-ignore
      this.rect.fill(p.fill)
    if (props.hasOwnProperty("stroke")
      || props.hasOwnProperty("width")
      || props.hasOwnProperty("height")
    ) {
      // @ts-ignore
      this.rect.radius(radius).size(p.width, p.height).move(p.x, p.y)
      // @ts-ignore
      this.back.radius(radius - p.stroke).size(p.width - p.stroke * 2, p.height - p.stroke * 2)
        .cx(p.x + p.width / 2).cy(p.y + p.height / 2)
      // @ts-ignore
      this.cell.radius(radius - p.stroke * 1.5)
        .cx(p.x + radius).cy(p.y + p.height / 2)
    }
  },

  //更新数据
  update(values: any) {

  },

};
