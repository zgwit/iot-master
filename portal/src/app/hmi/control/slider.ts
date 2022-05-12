import {HmiComponent} from "../hmi";

export let SliderComponent: HmiComponent = {
  uuid: "slider",
  name: "滑块",
  icon: "/assets/hmi/slider.svg",
  group: "控件",

  properties: [
    {
      label: '颜色',
      name: 'color',
      type: 'color',
      default: '#000'
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
    this.rect = this.$element.rect().size("100%", "10%")
    // @ts-ignore
    let box = this.rect.bbox()
    // @ts-ignore
    this.rect.radius(box.height * 0.5).move(0, box.height * 4.5)
    // @ts-ignore
    this.cell = this.$element.circle(box.height * 2).cx(box.height * 3).cy(box.cy)
  },

  resize() {
    // @ts-ignore
    let box = this.rect.bbox()
    // @ts-ignore
    this.rect.radius(box.height * 0.5).move(0, box.height * 4.5)
    // @ts-ignore
    this.cell.rx(box.height * 2).cx(box.height * 3).cy(box.cy)
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
