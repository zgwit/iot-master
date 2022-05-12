import {HmiComponent, HmiPropertyItem} from "../hmi";
import {fontProperties} from "../properties";

const template = `
    <rect height="100%" width="100%"></rect>
  `;

export let ButtonComponent: HmiComponent = {
  uuid: "button",
  name: '按钮',
  group: '控件',
  icon: "/assets/hmi/button.svg",
  //template,

  properties: [
    {
      label: '颜色',
      name: 'color',
      type: 'color',
      default: '#fff'
    },
    {
      label: '背景',
      name: 'fill',
      type: 'color',
      default: '#8BBB11'
    },
    ...fontProperties
  ],

  //配置
  create(props: any) {
    // @ts-ignore
    //this.$element.svg(template)
    this.rect = this.$element.rect().size("100%", "100%").radius(10)

    // @ts-ignore
    this.text = this.$element.text("按钮")

    // @ts-ignore
    let box = this.rect.bbox()
    // @ts-ignore
    this.text.center(box.cx, box.cy)
  },

  resize() {
    // @ts-ignore
    let box = this.rect.bbox()
    // @ts-ignore
    this.text.center(box.cx, box.cy)
  },

  //配置
  setup(props: any) {
    if (props.color) { // @ts-ignore
      this.text.fill(props.color)
    }
    if (props.fill) { // @ts-ignore
      this.rect.fill(props.fill)
    }
    if (props.hasOwnProperty("font")) {// @ts-ignore
      this.text.font({family: props.font})
    }
    if (props.hasOwnProperty("fontsize")) {// @ts-ignore
      this.text.font({size: props.fontsize})
    }
    if (props.hasOwnProperty("bold")) {// @ts-ignore
      this.text.font({weight: props.bold ? "bold" : "normal"})
    }

  },

  //更新数据
  update(values: any) {

  },

};
