import {HmiComponent} from "../hmi";
import {LineComponent} from "./basic/line";
import {CircleComponent} from "./basic/circle";
import {PolylineComponent} from "./basic/polyline";
import {PolygonComponent} from "./basic/polygon";
import {RectComponent} from "./basic/rect";
import {EllipseComponent} from "./basic/ellipse";
import {TextComponent} from "./basic/text";
import {ImageComponent} from "./basic/image";
import {ButtonComponent} from "./control/button";
import {InputComponent} from "./control/input";
import {SwitchComponent} from "./control/switch";
import {SliderComponent} from "./control/slider";
import {ValueComponent} from "./control/value";
import {ClockComponent} from "./control/clock";
import {CameraComponent} from "./control/camera";
import {WeatherComponent} from "./control/weather";
import {AlarmComponent} from "./industry/alarm";
import {CanComponent} from "./industry/can";
import {FanComponent} from "./industry/fan";
import {LightComponent} from "./industry/light";
import {PipeComponent} from "./industry/pipe";
import {PoolComponent} from "./industry/pool";
import {PumpComponent} from "./industry/pump";
import {ValveComponent} from "./industry/valve";
import {MotorComponent} from "./industry/motor";
import {BarChartComponent} from "./chart/bar";
import {GaugeChartComponent} from "./chart/gauge";
import {LineChartComponent} from "./chart/line";
import {PieChartComponent} from "./chart/pie";
import {RadarChartComponent} from "./chart/radar";

export let GroupedComponents: Array<Group> = [];

let indexedComponents: { [name: string]: HmiComponent } = {}
let indexedGroupComponents: { [name: string]: Array<HmiComponent> } = {}

export interface Group {
  name: string,
  components: Array<HmiComponent>
}

export function GetComponent(id: string): HmiComponent {
  return indexedComponents[id];
}

export function LoadComponent(obj: HmiComponent) {
  //if (indexedComponents.hasOwnProperty(obj.uuid))
  indexedComponents[obj.uuid] = obj;

  if (!obj.group)
    obj.group = "扩展";
  let group = indexedGroupComponents[obj.group]
  if (!group) {
    group = indexedGroupComponents[obj.group] = [];
    GroupedComponents.push({name: obj.group, components: group});
  }
  group.push(obj)
}

//export function loadIntervalComponents() {
let internalComponents = [
  //基础
  LineComponent, CircleComponent, EllipseComponent, PolylineComponent, PolygonComponent, RectComponent, TextComponent, ImageComponent,
  //控件
  ButtonComponent, InputComponent, SwitchComponent, SliderComponent, ValueComponent, ClockComponent, CameraComponent, WeatherComponent,
  //工业
  AlarmComponent, CanComponent, FanComponent, LightComponent, PipeComponent, PoolComponent, PumpComponent, ValveComponent, MotorComponent,
  //图表
  BarChartComponent, GaugeChartComponent, LineChartComponent, PieChartComponent, RadarChartComponent
]
internalComponents.forEach(c => LoadComponent(c))
//}

