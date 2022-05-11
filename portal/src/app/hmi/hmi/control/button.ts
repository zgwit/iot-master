import {HmiComponent} from "../hmi";
import {fontProperties} from "../../editor/properties";

const template = `
    <rect height="100%" width="100%"></rect>
  `;

export let ButtonComponent: HmiComponent = {
  uuid: "button",
  name: '按钮',
  group: '控件',
  icon: "/assets/hmi/button.svg",
  //template,

  color: true,
  stroke: true,

  properties: [...fontProperties],

  //配置
  create(props: any) {
    // @ts-ignore
    //this.$element.svg(template)
    this.rect = this.$element.rect().size("100%", "100%")

    // @ts-ignore
    this.text = this.$element.text("按钮")

    // @ts-ignore
    let box = this.rect.bbox()
    // @ts-ignore
    this.text.center(box.cx, box.cy)
  },

  //配置
  setup(props: any) {
    //console.log(props)
    if (props.width || props.height) {
      // @ts-ignore
      let box = this.rect.bbox()
      // @ts-ignore
      this.text.center(box.cx, box.cy)
    }

    if (props.stroke) { // @ts-ignore
      this.text.fill(props.stroke)
    }
    if (props.color) { // @ts-ignore
      this.rect.fill(props.color)
    }
  },

  //更新数据
  update(values: any) {

  },

};
