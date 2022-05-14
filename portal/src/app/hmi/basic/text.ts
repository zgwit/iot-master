import {HmiComponent} from "../hmi";
import {fontProperties} from "../properties";

export let TextComponent: HmiComponent = {
  uuid: "text",
  name: "文本",
  icon: "/assets/hmi/text.svg",
  group: "基础组件",
  drawer: "rect",

  color: true,
  stroke: true,

  properties: [...fontProperties],

  create() {
  },

  setup(properties: any): void {

  }
}
