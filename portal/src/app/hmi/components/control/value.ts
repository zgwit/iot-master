import {HmiComponent} from "../../hmi";
import {fontProperties} from "../../properties";

export let ValueComponent: HmiComponent = {
  id: "value",
  name: "数值",
  group: '控件',
  icon: "/assets/hmi/value.svg",
  drawer: "rect",

  properties: [
    {
      label: '前缀',
      name: 'prefix',
      type: 'text',
      default: ''
    },
    {
      label: '后缀',
      name: 'suffix',
      type: 'text',
      default: ''
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
      default: 0
    },
    {
      label: '边框',
      name: 'stroke',
      type: 'number',
      default: 5
    },
    ...fontProperties
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
    this.back = this.$container.rect()
    // @ts-ignore
    this.text = this.$container.text('123')
  },

  //配置
  setup(props: any) {
    //@ts-ignore
    let p = this.$properties
    if (props.hasOwnProperty("prefix") || props.hasOwnProperty("suffix"))// @ts-ignore
      this.text.text(p.prefix + '123' + p.suffix)
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
      this.rect.radius(p.radius).size(p.width, p.height).move(p.x, p.y)
      // @ts-ignore
      this.back.radius(p.radius).size(p.width - p.stroke * 2, p.height - p.stroke * 2)
        .cx(p.x + p.width / 2).cy(p.y + p.height / 2)
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
    if (values.hasOwnProperty("value"))// @ts-ignore
      this.text.text(p.prefix + values.value + p.suffix)

  },

};
