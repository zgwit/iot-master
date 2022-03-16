import {HmiComponent} from "../../hmi";
import {fontProperties} from "../properties";

const template = `
    <rect height="40" width="80"></rect>
  `;

export let ButtonComponent: HmiComponent = {
  uuid: "button",
  name: '按钮',
  group: '控件',
  icon: "/assets/hmi/components/button.svg",
  //template,

  properties: [...fontProperties],

  //配置
  init(props: any) {
    // @ts-ignore
    this.$element.svg(template)
  },

  //配置
  setup(props: any) {

    console.log(props)
  },

  //更新数据
  update(values: any) {

  },

};
