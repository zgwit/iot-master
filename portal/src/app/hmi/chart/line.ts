import {HmiComponent} from "../hmi";
import * as echarts from "echarts";

export let LineChartComponent: HmiComponent = {
  uuid: "line-chart",
  name: "折线图",
  icon: "/assets/hmi/chart-line.svg",
  group: "图表",
  drawer: "rect",

  data() {
    return {
      options: {
        tooltip: {},
        xAxis: {data: ['a', 'b', 'c', 'e']},
        yAxis: {},
        series: [
          {
            name: "d",
            type: "line",
            data: [1, 2, 3, 4]
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
