import {HmiComponent} from "../../hmi";

export let SliderComponent: HmiComponent = {
  id: "slider",
  name: "滑块",
  icon: "/assets/hmi/slider.svg",
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
      name: 'fill',
      type: 'color',
      default: '#fff'
    },
  ],

  events: [
    {label: '点击', name: 'click'},
    {label: '变化', name: 'change'},
  ],

  values: [
    {label: '值', name: 'value'}
  ],

  //配置
  create() {
    // @ts-ignore
    this.rect = this.$container.rect()
    // @ts-ignore
    this.cell = this.$container.circle(0)
  },

  //配置
  setup(props: any) {
    //@ts-ignore
    let p = this.$properties
    // @ts-ignore
    let radius = p.height / 4
    let radius2 = p.height / 2

    if (props.hasOwnProperty("color"))  // @ts-ignore
      this.cell.fill(p.color)
    if (props.hasOwnProperty("fill"))  // @ts-ignore
      this.rect.fill(p.fill)
    if (props.hasOwnProperty("stroke")
      || props.hasOwnProperty("width")
      || props.hasOwnProperty("height")
    ) {
      // @ts-ignore
      this.rect.radius(radius).size(p.width, p.height / 2).x(p.x).cy(p.y + p.height / 2)
      // @ts-ignore
      this.cell.radius(radius2).cx(p.x + radius + radius2).cy(p.y + p.height / 2)
    }
  },
}
