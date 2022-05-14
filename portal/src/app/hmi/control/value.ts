import {HmiComponent} from "../hmi";

export let ValueComponent: HmiComponent = {
  uuid: "value",
  name: "值",
  icon: "/assets/hmi/value.svg",
  group: "控件",

  color: true,
  stroke: true,

  create() {
  },

  setup(props: any): void {
    //elem.text(props.text || '0')
  }
}
