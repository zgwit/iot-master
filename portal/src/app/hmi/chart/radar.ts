import {HmiComponent} from "../hmi";
import * as echarts from "echarts";

export let RadarChartComponent: HmiComponent = {
  uuid: "radar-chart",
  name: "雷达图",
  icon: "/assets/hmi/chart-radar.svg",
  group: "图表",
  drawer: "rect",

  data() {
    return {
      options: {
        tooltip: {},
        radar: {
          //shape: 'circle',
          indicator: [
            {name: 'A', max: 50},
            {name: 'B', max: 50},
            {name: 'C', max: 50},
            {name: 'D', max: 50},
            {name: 'E', max: 50},
          ]
        },
        series: [
          {
            name: "d",
            type: "radar",
            data: [
              {name: "cpu", value: [10,35,30,40,45]},
            ]
          }
        ]
      }
    }
  },

  create() {
    //@ts-ignore
    this.element = this.$container.foreignObject()
    //@ts-ignore
    this.chart = echarts.init(this.element.node)
    //@ts-ignore
    this.chart.setOption(this.options)
  },

  setup(props: any): void {
    //@ts-ignore
    let p = this.$properties
    if (props.hasOwnProperty("width") || props.hasOwnProperty("height")) {
      //@ts-ignore
      this.element.size(p.width, p.height)
      //@ts-ignore
      this.chart.resize()
    }
  },

  update(values: any) {

  }
}
