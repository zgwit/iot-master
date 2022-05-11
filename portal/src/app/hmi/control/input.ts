import {HmiComponent} from "../hmi";

export let InputComponent: HmiComponent = {
  uuid: "input",
  name: "输入框",
  icon: "/assets/hmi/input.svg",
  group: "控件",
  type: "object",

  color: true,
  stroke: true,

  setup(props: any): void {

  }
}
