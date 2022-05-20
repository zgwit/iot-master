import {HmiComponent} from "../../hmi";
import {fontProperties} from "../../properties";

export let InputComponent: HmiComponent = {
  id: "input",
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
    {
      label: '类型',
      name: 'type',
      type: 'select',
      options: [
        {value: 'text', label:'文本'},
        {value: 'number', label:'数值'},
        {value: 'integer', label:'整数'},
      ],
      default: 'text'
    },
    {
      label: '内容提示',
      name: 'placeholder',
      type: 'text',
      default: '输入框'
    },
    ...fontProperties
  ],

  events: [
    {name: "click", label: "点击"},
    {name: "change", label: "变化"},
  ],

  values: [
    {name: "value", label: "内容"},
  ],

  create() {
    //@ts-ignore
    let p = this.$properties
    //@ts-ignore
    this.element = this.$container.foreignObject()
    let input = document.createElement("input");
    input.setAttribute("type", p.type);
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

  setup(props: any): void {
    //@ts-ignore
    let p = this.$properties

    if (props.hasOwnProperty("color"))  // @ts-ignore
      this.input.setAttribute("style", "width:100%; height:100%;color:" + props.color + ";");
    if (props.hasOwnProperty("fill"))  // @ts-ignore
      this.input.setAttribute("style", "width:100%; height:100%;background-color:" + props.fill + ";");
    if (props.hasOwnProperty("type"))  // @ts-ignore
      this.input.setAttribute("type", p.type);
    if (props.hasOwnProperty("placeholder"))  // @ts-ignore
      this.input.setAttribute("placeholder", p.placeholder);
    if (props.hasOwnProperty("width") || props.hasOwnProperty("height")   ) // @ts-ignore
      this.element.size(p.width, p.height)
  },

  //更新数据
  update(values: any) {

  },
}
