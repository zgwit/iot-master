import {HmiComponent} from "../hmi";
import * as echarts from "echarts";

export let PieChartComponent: HmiComponent = {
  uuid: "pie-chart",
  name: "饼状图",
  icon: "/assets/hmi/chart-pie.svg",
  group: "图表",
  type: "object",

  data() {
    return {
      options: {
        tooltip: {},
        series: [
          {
            name: "d",
            type: "pie",
            radius: '65%',
            center: ['50%', '50%'],
            data: [
              {name: "cpu", value: 20},
              {name: "mem", value: 60}]
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
