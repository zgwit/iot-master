import {HmiComponent} from "../hmi";

export let ImageComponent: HmiComponent = {
  uuid: "image",
  name: "图像",
  icon: "/assets/hmi/image.svg",
  group: "基础组件",
  drawer: "rect",


  properties: [
    {
      label: 'URL',
      name: 'url',
      type: 'text',
    },
  ],

  create() {
    //@ts-ignore
    this.element = this.$container.image().load(this.$properties.url || "/assets/hmi/image.svg")
  },

  setup(props: any): void {
    //@ts-ignore
    let p = this.$properties
    if (props.hasOwnProperty("width") || props.hasOwnProperty("height"))//@ts-ignore
      this.element.size(p.width, p.height)
    if (props.hasOwnProperty("url"))//@ts-ignore
      this.element.load(p.url)
  }
}
