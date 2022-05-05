import { RectNode, RectNodeModel, h } from "@logicflow/core";
import { eletype } from "../config";

class UserTaskModel extends RectNodeModel {
  static extendKey = "UserTaskModel";
  setAttributes() {
    this.width = 90;
    this.height = 42;
    this.radius = 4;
  }
}

class UserTaskView extends RectNode {
  static extendKey = "UserTaskNode";
  getShape(): any {
    const { model } = this.props;
    const { x, y, width, height, radius } = model;
    const style = model.getNodeStyle();
    return h("g", {}, [
      h("rect", {
        ...style,
        x: x - width / 2,
        y: y - height / 2,
        rx: radius,
        ry: radius,
        width,
        height,
        fill: "#E7F7FE",
        stroke: "#1890FF",
        style: {
          strokeWidth: 1,
        },
      }),
      h(
        "svg",
        {
          x: x - width / 2 + 4,
          y: y - height / 2 + 4,
          width: 14,
          height: 14,
          viewBox: "0 0 1274 1024",
        },
        h("path", {
          fill: "#1890FF",
          d: "M289.071721 313.25654c-1.350764-101.018741 12.281721-176.459944 106.60804-227.359194 19.306718-10.417257 65.041342-15.911388 86.877673-15.911388 23.923876-2.447748 0.23843-21.964244 53.016471-21.964245s126.176724 23.252587 161.259754 95.225805c35.08303 71.974241 38.052664 150.966317 38.052664 169.903622 0 18.938328 19.383466 9.468652 19.383466 37.875632 0 28.40698-8.361435 84.613096-38.10076 122.79265-29.740348 38.179554-39.432081 47.648207-58.815547 76.055187-19.383466 28.40698-29.075199 59.379389-29.075199 85.22094 0 94.689592 162.000628 101.68696 255.906368 161.484881 37.066197 23.603581 63.063291 54.19532 77.990258 91.774194 13.048177 32.848129-3.00238 70.053496-35.850509 83.101673a63.984267 63.984267 0 0 1-23.62507 4.520966l-781.415033 0.001024c-35.344996 0-63.99757-28.652574-63.997569-63.99757 0-8.118911 1.545192-16.163121 4.551666-23.704888 14.97097-37.543058 40.871873-68.107167 77.703732-91.694375 93.377714-59.798944 256.137635-66.796312 256.137635-161.484881 0-25.217335-9.691733-56.81396-29.075199-85.220941-19.383466-28.40698-29.075199-37.875632-58.149375-76.055186-29.075199-38.179554-38.766932-94.385671-38.766932-122.792651 0-28.40698 19.383466-18.938328 19.383466-37.771255z",
        })
      ),
    ]);
  }
}

const UserTask = {
  type: eletype.usertask,
  view: UserTaskView,
  model: UserTaskModel,
};

export { UserTaskModel, UserTaskView };
export default UserTask;
