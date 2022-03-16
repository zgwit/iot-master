import {HmiComponent} from "../../hmi";

export let ValueComponent: HmiComponent = {
  uuid: "value",
  name: "值",
  icon: "/assets/hmi/components/value.svg",
  group: "控件",
  type: "text",

  setup(props: any): void {
    //elem.text(props.text || '0')
  }
}
