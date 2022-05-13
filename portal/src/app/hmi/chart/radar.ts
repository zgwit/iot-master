import {HmiComponent} from "../hmi";
import * as echarts from "echarts";

export let RadarChartComponent: HmiComponent = {
  uuid: "radar-chart",
  name: "雷达图",
  icon: "/assets/hmi/chart-radar.svg",
  group: "图表",
  type: "object",

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

  create(props: any) {
    //@ts-ignore
    this.chart = echarts.init(this.$element.node)
    //@ts-ignore
    this.chart.setOption(this.options)
  },

  resize() {
    //@ts-ignore
    this.chart.resize()
  },

  setup(props: any): void {

  },

  update(values: any) {

  }
}
