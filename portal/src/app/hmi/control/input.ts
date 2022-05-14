import {HmiComponent} from "../hmi";
import {borderProperties, colorProperties, fontProperties} from "../properties";

export let InputComponent: HmiComponent = {
  uuid: "input",
  name: "输入框",
  icon: "/assets/hmi/input.svg",
  group: "控件",
  drawer: "rect",

  properties: [
    {
      label: '颜色',
      name: 'color',
      type: 'color',
      default: '#fff'
    },
    {
      label: '背景',
      name: 'fill',
      type: 'color',
      default: '#8BBB11'
    },
    ...fontProperties
  ],

  create() {
    let input = document.createElement("input");
    input.setAttribute("type","text");
    input.setAttribute("placeholder","输入框");
    input.setAttribute("style","width:100%; height:100%;background-color:black;");
    input.onchange = (event) => {
      console.log("input change", event)
    }
    //@ts-ignore
    this.input = input
    //@ts-ignore
    this.$element.node.appendChild(input)
  },

  resize() {
    //this.input.sets
  },

  setup(props: any): void {
    if (props.color) { // @ts-ignore
      this.input.setAttribute("style","width:100%; height:100%;color:"+props.color+";");
    }
    if (props.fill) { // @ts-ignore
      this.input.setAttribute("style","width:100%; height:100%;background-color:"+props.fill+";");
    }

  }
}
