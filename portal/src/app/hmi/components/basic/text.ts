import {HmiComponent} from "../../hmi";
import {fontProperties} from "../../properties";

export let TextComponent: HmiComponent = {
  id: "text",
  name: "文本",
  icon: "/assets/hmi/text.svg",
  group: "基础组件",
  drawer: "rect",

  color: true,

  properties: [...fontProperties],

  create() {
    //@ts-ignore
    this.element = this.$container.text(this.$properties.text || "文本")
  },

  setup(props: any): void {
    //@ts-ignore
    let p = this.$properties
    if (props.hasOwnProperty("fill"))//@ts-ignore
      this.element.fill(p.fill)
    if (props.hasOwnProperty("width") || props.hasOwnProperty("height"))//@ts-ignore
      this.element.size(p.width, p.height)
  }
}
