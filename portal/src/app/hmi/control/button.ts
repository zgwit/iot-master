import {HmiComponent, HmiPropertyItem} from "../hmi";
import {fontProperties} from "../properties";

export let ButtonComponent: HmiComponent = {
  uuid: "button",
  name: '按钮',
  group: '控件',
  icon: "/assets/hmi/button.svg",
  drawer: "rect",

  properties: [
    {
      label: '文字',
      name: 'text',
      type: 'text',
      default: '按钮'
    },
    {
      label: '颜色',
      name: 'color',
      type: 'color',
      default: '#fff'
    },
    {
      label: '背景',
      name: 'back',
      type: 'color',
      default: '#8BBB11'
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
    ...fontProperties
  ],

  //配置
  create() {
    // @ts-ignore
    this.rect = this.$container.rect()
    // @ts-ignore
    this.back = this.$container.rect()
    // @ts-ignore
    this.text = this.$container.text(this.$properties.text)
  },

  //配置
  setup(props: any) {
    //@ts-ignore
    let p = this.$properties
    if (props.hasOwnProperty("text"))// @ts-ignore
      this.text.text(p.text)
    if (props.hasOwnProperty("color"))  // @ts-ignore
      this.text.fill(p.color)
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
      this.back.radius(p.radius).size(p.width - p.stroke * 2, p.height - p.stroke * 2)
        .cx(p.x + p.width / 2).cy(p.y + p.height / 2)
      // @ts-ignore
      this.text.center(p.x + p.width / 2, p.y + p.height / 2)
    }
    if (props.hasOwnProperty("x") || props.hasOwnProperty("y")) {
      // @ts-ignore
      this.rect.move(p.x, p.y)
      // @ts-ignore
      this.back.cx(p.x + p.width / 2).cy(p.y + p.height / 2)
      // @ts-ignore
      this.text.center(p.x + p.width / 2, p.y + p.height / 2)
    }
    if (props.hasOwnProperty("font"))// @ts-ignore
      this.text.font({family: p.font})
    if (props.hasOwnProperty("fontsize")) // @ts-ignore
      this.text.font({size: p.fontsize})
    if (props.hasOwnProperty("bold")) // @ts-ignore
      this.text.font({weight: p.bold ? "bold" : "normal"})
  },

  //更新数据
  update(values: any) {

  },

};
