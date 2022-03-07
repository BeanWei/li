export default {
  name: "qqzzjakwkwl",
  type: "void",
  "x-component": "Menu",
  "x-component-props": {
    defaultSelectedKeys: ["WelcomePage"],
    onClickMenuItem: "{{ onClickMenuItem }}",
    menuData: [
      {
        key: "WelcomePage",
        name: "欢迎",
      },
      {
        key: "AdminPage",
        name: "管理页",
        children: [
          {
            key: "AdminPageSub1",
            name: "一级页面",
          },
          {
            key: "AdminPageSub2",
            name: "二级页面",
          },
          {
            key: "AdminPageSub3",
            name: "三级页面",
          },
        ],
      },
      {
        name: "列表页",
        key: "ListPage",
        children: [
          {
            key: "AdminPageSub1Sub",
            name: "一级列表页面",
            children: [
              {
                key: "AdminPageSub1Sub1",
                name: "一一级列表页面",
              },
              {
                key: "AdminPageSub1Sub2",
                name: "一二级列表页面",
              },
              {
                key: "sAdminPageSub1Sub3",
                name: "一三级列表页面",
              },
            ],
          },
          {
            key: "AdminPageSub2",
            name: "二级列表页面",
          },
          {
            key: "AdminPageSub3",
            name: "三级列表页面",
          },
        ],
      },
      {
        target: "https://ant.design",
        name: "Ant Design 官网外链",
      },
    ],
  },
};
