import {HmiComponent} from "../hmi";

export let WeatherComponent: HmiComponent = {
  uuid: "weather",
  name: "天气",
  icon: "/assets/hmi/weather.svg",
  group: "控件",
  type: "object",

  color: true,
  stroke: true,

  setup(props: any): void {

  }
}
