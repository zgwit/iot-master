import {
  Circle,
  Container,
  ElementAlias,
  Ellipse,
  ForeignObject, Image,
  Line, Path,
  Polygon,
  Polyline,
  Rect,
  Svg,
  Text,
} from "@svgdotjs/svg.js";
import {HmiComponent} from "../hmi";

export function CreateElement(container: Container, component: HmiComponent): ElementAlias {
  let element: ElementAlias
  const type = component.type || "svg"
  switch (type) {
    case "rect" :
      element = new Rect() // container.rect();
      break;
    case "circle" :
      element = new Circle() // container.circle();
      break;
    case "ellipse" :
      element = new Ellipse() // container.ellipse();
      break;
    case "line" :
      element = new Line() // container.line();
      break;
    case "polyline" :
      element = new Polyline() // container.polyline();
      break;
    case "polygon" :
      element = new Polygon() // container.polygon();
      break;
    case "image" :
      element = new Image().load("/assets/hmi/components/image.svg") // container.image();
      break;
    case "path" :
      element = new Path() // container.path();
      break;
    case "text" :
      element = new Text().text("文本") // container.text("文本");
      break;
    case "svg" :
      element = new Svg() // container.nested();
      break;
    case "object":
      element = new ForeignObject() // container.foreignObject(0, 0);
      break;
    default:
      throw new Error("不支持的控件类型：" + type)
  }
  return element;
}
