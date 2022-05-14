import {HmiComponent} from "../hmi";

export let LightComponent: HmiComponent = {
  uuid: "light",
  name: "指示灯",
  icon: "/assets/hmi/light.svg",
  group: "工业",

  color: true,

  create() {
  },

  setup(props: any): void {

  }
}
