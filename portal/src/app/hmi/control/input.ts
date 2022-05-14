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
    //@ts-ignore
    this.element = this.$container.foreignObject().move(p.x, p.y);
    let input = document.createElement("input");
    input.setAttribute("type", "text");
    input.setAttribute("placeholder", "输入框");
    input.setAttribute("style", "width:100%; height:100%;background-color:black;");
    input.onchange = (event) => {
      console.log("input change", event)
    }
    //@ts-ignore
    this.input = input
    //@ts-ignore
    this.element.node.appendChild(input)
  },

  resize() {
    //this.input.sets
  },

  setup(props: any): void {
    //@ts-ignore
    let p = this.$properties

    if (props.hasOwnProperty("color"))  // @ts-ignore
      this.input.setAttribute("style", "width:100%; height:100%;color:" + props.color + ";");
    if (props.hasOwnProperty("fill"))  // @ts-ignore
      this.input.setAttribute("style", "width:100%; height:100%;background-color:" + props.fill + ";");

    if (props.hasOwnProperty("width") || props.hasOwnProperty("height")   ) // @ts-ignore
      this.element.size(p.width, p.height)
  },

  //更新数据
  update(values: any) {

  },
}
