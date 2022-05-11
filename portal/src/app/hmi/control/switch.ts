import {HmiComponent} from "../hmi";

export let SwitchComponent: HmiComponent = {
  uuid: "switch",
  name: "开关",
  icon: "/assets/hmi/switch.svg",
  group: "控件",
  type: "object",

  color: true,
  stroke: true,

  setup(props: any): void {

  }
}
