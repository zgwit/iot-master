import {HmiComponent} from "../../hmi";

export let CameraComponent: HmiComponent = {
  id: "camera",
  name: "摄像头",
  icon: "/assets/hmi/camera.svg",
  group: "控件",
  drawer: "rect",

  properties: [
    {
      label: 'URL',
      name: 'url',
      type: 'text',
      default: ''
    },
  ],

  create() {
    //@ts-ignore
    this.video = document.createElement("video")
    //@ts-ignore
    this.element = this.$container.foreignObject()
    //@ts-ignore
    this.element.node.appendChild(this.video)

    const v = document.createElement("video")
    v.autoplay = true
  },

  setup(props: any): void {
    //@ts-ignore
    let p = this.$properties
    if (props.hasOwnProperty("url")) {
      //@ts-ignore
      this.video.autoplay = true
      //@ts-ignore
      this.video.src = p.url
    }
  }
}
