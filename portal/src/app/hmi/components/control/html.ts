import {HmiComponent} from "../../hmi";

export let HtmlComponent: HmiComponent = {
  id: "html",
  name: "HTML",
  icon: "/assets/hmi/code.svg",
  group: "控件",
  drawer: "rect",

  properties: [
    {
      label: 'HTML',
      name: 'html',
      type: 'textarea',
      default: '<div>HTML</div>'
    },
  ],

  create() {
    //@ts-ignore
    this.element = this.$container.foreignObject()
    //@ts-ignore
    this.element.node.innerHTML = this.$properties.html
    //@ts-ignore
    this.element.node.firstChild.style = "width:100%; height:100%;" //TODO 代码无效。。。。
  },

  setup(props: any): void {
    //@ts-ignore
    let p = this.$properties
    if (props.hasOwnProperty("width") || props.hasOwnProperty("height")) {
      //@ts-ignore
      this.element.size(p.width, p.height)
    }
    if (props.hasOwnProperty("html"))//@ts-ignore
      this.element.node.innerHTML = p.html
  },
}
