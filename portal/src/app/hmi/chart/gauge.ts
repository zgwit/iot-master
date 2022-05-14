import {HmiComponent} from "../hmi";
import * as echarts from "echarts";

export let GaugeChartComponent: HmiComponent = {
  uuid: "gauge-chart",
  name: "仪表盘",
  icon: "/assets/hmi/chart-gauge.svg",
  group: "图表",
  drawer: "rect",

  data() {
    return {
      options: {
        tooltip: {},
        series: [
          {
            name: "d",
            type: "gauge",
            data: [{name: "cpu", value: 20}]
          }
        ]
      }
    }
  },

  create() {
    //@ts-ignore
    this.foreignObject = this.$container.foreignObject()
    //@ts-ignore
    this.chart = echarts.init(this.foreignObject.node)
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
