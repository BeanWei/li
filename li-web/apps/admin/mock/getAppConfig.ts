export default {
  title: "Li Demo",
  copyright: "Powered by ❤️璃❤️",
  home: "Welcome",
  navitems: [
    {
      type: "void",
      name: "langswitch",
      "x-component": "LangSwitch",
    },
    {
      type: "void",
      name: "themeswitch",
      "x-component": "ThemeSwitch",
    },
    {
      type: "void",
      name: "currentuser",
      "x-component": "DropdownMenu",
      "x-component-props": {
        triggerProps: {
          popupStyle: {
            marginRight: "28px",
          },
        },
      },
      items: {
        type: "void",
        properties: {
          editprofile: {
            type: "void",
            "x-component": "DropdownMenu.Item",
            properties: {
              action: {
                title: "用户设置",
                type: "void",
                "x-component": "Action.FormDrawer",
                "x-component-props": {
                  isMenuItem: true,
                  initialValues: {
                    nickname: "{{global.currentUser.nickname}}",
                  },
                  forSubmit: "updateProfile",
                },
                properties: {
                  nickname: {
                    type: "string",
                    title: "昵称",
                    "x-decorator": "FormItem",
                    "x-component": "Input",
                  },
                  avatar: {
                    type: "string",
                    title: "头像",
                    "x-decorator": "FormItem",
                    "x-component": "Upload.Avatar",
                  },
                },
              },
            },
          },
          changepwd: {
            type: "void",
            title: "修改密码",
            "x-component": "DropdownMenu.Item",
          },
          divider: {
            type: "void",
            "x-component": "Divider",
            "x-component-props": {
              style: {
                margin: "4px 0",
              },
            },
          },
          signout: {
            type: "void",
            title: "退出登录",
            "x-component": "DropdownMenu.Item",
          },
        },
      },
      properties: {
        curUser: {
          type: "void",
          "x-component": "Avatar",
          "x-component-props": {
            size: 32,
            style: {
              cursor: "pointer",
            },
            src: "{{global.currentUser.avatar}}",
          },
        },
      },
    },
  ],
  menus: [
    {
      key: "Welcome",
      title: "欢迎",
      icon: "IconHome",
      path: "Welcome",
    },
    {
      key: "System",
      path: "System",
      title: "系统管理",
      icon: "IconSettings",
      children: [
        {
          key: "SystemUser",
          title: "用户管理",
          path: "System.SystemUser",
        },
      ],
    },
    {
      title: "列表页",
      key: "Listpage",
      path: "Listpage",
      icon: "IconList",
      children: [
        {
          key: "AdminPageSub1Sub",
          path: "Listpage.AdminPageSub1Sub",
          title: "一级列表页面",
          children: [
            {
              key: "AdminPageSub1Sub1",
              path: "Listpage.AdminPageSub1Sub.AdminPageSub1Sub1",
              title: "一一级列表页面",
            },
            {
              key: "AdminPageSub1Sub2",
              path: "Listpage.AdminPageSub1Sub.AdminPageSub1Sub2",
              title: "一二级列表页面",
            },
            {
              key: "sAdminPageSub1Sub3",
              path: "Listpage.AdminPageSub1Sub.sAdminPageSub1Sub3",
              title: "一三级列表页面",
            },
          ],
        },
        {
          key: "AdminPageSub2",
          title: "二级列表页面",
          path: "Listpage.AdminPageSub2",
        },
        {
          key: "AdminPageSub3",
          title: "三级列表页面",
          path: "Listpage.AdminPageSub3",
        },
      ],
    },
    {
      target: "https://github.com/BeanWei",
      key: "qqq",
      path: "qqq",
      title: "Github",
      icon: "IconGithub",
    },
  ],
  binding: {
    signform: {
      type: "object",
      "x-component": "Form",
      properties: {
        email: {
          type: "string",
          required: true,
          "x-component": "Input",
          "x-validator": "email",
          "x-decorator": "FormItem",
          "x-component-props": { placeholder: "邮箱", style: {} },
        },
        password: {
          type: "string",
          "x-component": "Password",
          required: true,
          "x-decorator": "FormItem",
          "x-component-props": { placeholder: "密码", style: {} },
        },
        submit: {
          type: "void",
          "x-content": "登录",
          "x-component": "Submit",
          "x-decorator": "FormItem",
          "x-component-props": {
            long: true,
            type: "primary",
            style: { width: "100%" },
            forSubmit: "userSignIn",
            forSubmitSuccessTo: "/admin/Welcome",
          },
        },
      },
    },
  },
};
