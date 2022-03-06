export default {
  name: "qqzzjakwkwl",
  type: "void",
  "x-component": "Menu",
  "x-component-props": {
    defaultSelectedKeys: "{{ defaultSelectedKeys }}",
    onClickMenuItem: "{{ onClickMenuItem }}",
  },
  properties: {
    item1: {
      type: "void",
      title: "Menu Item u1",
      "x-data": "u1",
      "x-component": "Menu.Item",
      "x-component-props": {},
    },
    item2: {
      type: "void",
      title: "Menu Item u2",
      "x-data": "u2",
      "x-component": "Menu.Item",
      "x-component-props": {},
    },
    item3: {
      type: "void",
      title: "SubMenu u3",
      "x-data": "u3",
      "x-component": "Menu.SubMenu",
      "x-component-props": {},
      properties: {
        item31: {
          type: "void",
          title: "SubMenu u31",
          "x-data": "u31",
          "x-component": "Menu.SubMenu",
          "x-component-props": {},
          properties: {
            item311: {
              type: "void",
              title: "Menu Item u311",
              "x-data": "u311",
              "x-component": "Menu.Item",
              "x-component-props": {},
            },
            item312: {
              type: "void",
              title: "Menu Item u312",
              "x-data": "u312",
              "x-component": "Menu.Item",
              "x-component-props": {},
            },
          },
        },
        item32: {
          type: "void",
          title: "Menu Item u32",
          "x-data": "u32",
          "x-component": "Menu.Item",
          "x-component-props": {},
        },
        item33: {
          type: "void",
          title: "Menu Item u33",
          "x-data": "u33",
          "x-component": "Menu.Item",
          "x-component-props": {},
        },
      },
    },
    item4: {
      type: "void",
      title: "SubMenu u4",
      "x-data": "u4",
      "x-component": "Menu.SubMenu",
      "x-component-props": {},
      properties: {
        item41: {
          type: "void",
          title: "Menu Item u41",
          "x-data": "u41",
          "x-component": "Menu.Item",
          "x-component-props": {},
        },
      },
    },
  },
};
