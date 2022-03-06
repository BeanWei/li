export default {
  name: "qqzzjakwkwl",
  type: "void",
  "x-component": "Menu",
  "x-component-props": {
    mode: "mix",
    defaultSelectedUid: "u1",
    onSelect: "{{ onSelect }}",
    sideMenuRefScopeKey: "sideMenuRef",
  },
  properties: {
    item3: {
      type: "void",
      title: "SubMenu u3",
      "x-data": "u3",
      "x-component": "Menu.SubMenu",
      "x-component-props": {},
      properties: {
        item6: {
          type: "void",
          title: "SubMenu u6",
          "x-data": "u6",
          "x-component": "Menu.SubMenu",
          "x-component-props": {},
          properties: {
            item7: {
              type: "void",
              title: "Menu Item u7",
              "x-data": "u7",
              "x-component": "Menu.Item",
              "x-component-props": {},
            },
            item8: {
              type: "void",
              title: "Menu Item u8",
              "x-data": "u8",
              "x-component": "Menu.Item",
              "x-component-props": {},
            },
          },
        },
        item4: {
          type: "void",
          title: "Menu Item u4",
          "x-data": "u4",
          "x-component": "Menu.Item",
          "x-component-props": {},
        },
        item5: {
          type: "void",
          title: "Menu Item u5",
          "x-data": "u5",
          "x-component": "Menu.Item",
          "x-component-props": {},
        },
      },
    },
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
    item9: {
      type: "void",
      title: "SubMenu u9",
      "x-data": "u9",
      "x-component": "Menu.SubMenu",
      "x-component-props": {},
      properties: {
        item10: {
          type: "void",
          title: "Menu Item u10",
          "x-data": "u10",
          "x-component": "Menu.Item",
          "x-component-props": {},
        },
      },
    },
  },
};
