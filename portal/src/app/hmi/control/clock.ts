import {HmiComponent} from "../hmi";

export let ClockComponent: HmiComponent = {
  uuid: "clock",
  name: "时钟",
  icon: "/assets/hmi/clock.svg",
  group: "控件",

  drawer:"rect",

  color: true,
  stroke: true,

  create() {
  },

  setup(props: any): void {

  }
}
