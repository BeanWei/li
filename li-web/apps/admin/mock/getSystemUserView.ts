const friends = {
  type: "array",
  title: "朋友",
  "x-decorator": "FormItem",
  "x-component": "RecordSelect",
  "x-component-props": {
    multiple: true,
    fieldNames: {
      title: "nickname",
      avatar: "avatar",
      value: "id",
    },
    searchConfig: {
      operation: "listUser",
    },
  },
};

export default {
  type: "object",
  properties: {
    row1: {
      type: "void",
      "x-component": "Grid.Row",
      "x-component-props": {
        gutter: 16,
      },
      properties: {
        col1: {
          type: "void",
          "x-component": "Grid.Col",
          "x-component-props": {
            span: 16,
          },
          properties: {
            userlist: {
              type: "void",
              title: "用户列表",
              "x-decorator": "CardItem",
              "x-component": "List",
              "x-component-props": {
                forInit: "listUser",
                selection: {
                  multiple: true,
                },
                filter: true,
              },
              properties: {
                userlisttable: {
                  type: "array",
                  "x-component": "List.Table",
                  "x-component-props": {
                    rowKey: "id",
                  },
                  properties: {
                    actions: {
                      type: "void",
                      "x-component": "List.Action",
                      properties: {
                        reloadandreset: {
                          type: "void",
                          title: "刷新并重置",
                          "x-component": "List.Action.Reload",
                          "x-component-props": {
                            position: "left",
                            data: {
                              sorter: {},
                              filter: {},
                              page: 1,
                            },
                          },
                        },
                        reload: {
                          type: "void",
                          title: "刷新",
                          "x-component": "List.Action.Reload",
                          "x-component-props": {
                            position: "left",
                          },
                        },
                        add: {
                          type: "void",
                          title: "新建用户",
                          "x-component": "List.Action.RecordFormDrawer",
                          "x-component-props": {
                            type: "primary",
                          },
                          items: {
                            type: "object",
                            properties: {
                              nickname: {
                                type: "string",
                                title: "昵称",
                                "x-decorator": "FormItem",
                                "x-component": "Input",
                              },
                              money: {
                                type: "string",
                                title: "收入",
                                "x-decorator": "FormItem",
                                "x-component": "Money",
                              },
                              friends,
                            },
                          },
                          properties: {
                            cancel: {
                              type: "void",
                              "x-component": "Action.FormDrawer.Cancel",
                            },
                            submit: {
                              type: "void",
                              "x-component": "Action.FormDrawer.Submit",
                              "x-component-props": {
                                forSubmit: "createUser",
                              },
                            },
                          },
                        },
                        bulkdelete: {
                          type: "void",
                          title: "批量删除",
                          "x-component": "List.Action.RowSelection",
                          "x-component-props": {
                            forSubmit: "deleteManyUser",
                            afterReload: true,
                            confirmProps: {
                              title: "确认删除？",
                            },
                            status: "danger",
                          },
                        },
                      },
                    },
                  },
                  items: {
                    type: "object",
                    properties: {
                      column0: {
                        type: "void",
                        "x-component": "List.Table.Column",
                        "x-component-props": {
                          title: "编号",
                          dataIndex: "id",
                        },
                        properties: {
                          id: {
                            type: "string",
                            "x-component": "Input",
                            "x-read-pretty": true,
                          },
                        },
                      },
                      column1: {
                        type: "void",
                        "x-component": "List.Table.Column",
                        "x-component-props": {
                          title: "昵称",
                          dataIndex: "nickname",
                          filterable: true,
                        },
                        properties: {
                          nickname: {
                            type: "string",
                            "x-component": "Input",
                            "x-read-pretty": true,
                          },
                        },
                      },
                      column2: {
                        type: "void",
                        "x-component": "List.Table.Column",
                        "x-component-props": {
                          title: "邮箱",
                          dataIndex: "email",
                        },
                        properties: {
                          email: {
                            type: "string",
                            title: "NO-USED",
                            "x-decorator": "FormItem",
                            "x-component": "Input",
                            "x-read-pretty": true,
                          },
                        },
                      },
                      column3: {
                        type: "void",
                        "x-component": "List.Table.Column",
                        "x-component-props": {
                          title: "爱好",
                          dataIndex: "likes",
                          filterable: true,
                        },
                        properties: {
                          likes: {
                            type: "string",
                            "x-component": "Select",
                            enum: ["篮球", "LOL", "兵乓球"],
                          },
                        },
                      },
                      column4: {
                        type: "void",
                        "x-component": "List.Table.Column",
                        "x-component-props": {
                          title: "出生日期",
                          dataIndex: "birth_day",
                          filterable: true,
                        },
                        properties: {
                          birth_day: {
                            type: "string",
                            "x-component": "DatePicker",
                          },
                        },
                      },
                      column5: {
                        type: "void",
                        "x-component": "List.Table.Column",
                        "x-component-props": {
                          title: "是否管理员",
                          dataIndex: "is_admin",
                          filterable: true,
                        },
                        properties: {
                          is_admin: {
                            type: "bool",
                            "x-component": "Checkbox",
                          },
                        },
                      },
                      column6: {
                        type: "void",
                        "x-component": "List.Table.Column",
                        "x-component-props": {
                          title: "头像",
                          dataIndex: "avatar",
                        },
                        properties: {
                          avatar: {
                            type: "string",
                            "x-component": "Upload.Avatar",
                          },
                        },
                      },
                      column7: {
                        type: "void",
                        "x-component": "List.Table.Column",
                        "x-component-props": {
                          title: "朋友",
                          dataIndex: "friends",
                          filterable: true,
                        },
                        properties: {
                          friends,
                        },
                      },
                      column8: {
                        type: "void",
                        "x-component": "List.Table.Column",
                        "x-component-props": {
                          title: "操作",
                          dataIndex: "__action",
                          width: 150,
                        },
                        properties: {
                          recordactions: {
                            type: "void",
                            "x-component": "Space",
                            "x-component-props": {
                              split: "divider",
                              size: 0,
                            },
                            properties: {
                              read: {
                                type: "void",
                                "x-component": "List.Action.RecordFormDrawer",
                                "x-component-props": {
                                  forInit: "getUser",
                                  icon: "IconEye",
                                  type: "text",
                                  drawerProps: {
                                    title: "查看",
                                  },
                                },
                                items: {
                                  type: "object",
                                  properties: {
                                    nickname: {
                                      type: "string",
                                      required: true,
                                      title: "昵称",
                                      "x-decorator": "FormItem",
                                      "x-component": "Input",
                                    },
                                  },
                                },
                                properties: {
                                  cancel: {
                                    type: "void",
                                    "x-component": "Action.FormDrawer.Cancel",
                                  },
                                },
                              },
                              update: {
                                type: "void",
                                "x-component": "List.Action.RecordFormDrawer",
                                "x-component-props": {
                                  forInit: "getUser",
                                  icon: "IconEdit",
                                  type: "text",
                                  drawerProps: {
                                    title: "编辑",
                                  },
                                },
                                items: {
                                  type: "object",
                                  properties: {
                                    nickname: {
                                      type: "string",
                                      required: true,
                                      title: "昵称",
                                      "x-decorator": "FormItem",
                                      "x-component": "Input",
                                    },
                                    friends,
                                  },
                                },
                                properties: {
                                  cancel: {
                                    type: "void",
                                    "x-component": "Action.FormDrawer.Cancel",
                                  },
                                  submit: {
                                    type: "void",
                                    "x-component": "Action.FormDrawer.Submit",
                                    "x-component-props": {
                                      forSubmit: "updateUser",
                                    },
                                  },
                                },
                              },
                              delete: {
                                type: "void",
                                "x-component": "List.Action.RecordDelete",
                                "x-component-props": {
                                  status: "danger",
                                  icon: "IconDelete",
                                  type: "text",
                                  forSubmit: "deleteUser",
                                },
                              },
                            },
                          },
                        },
                      },
                    },
                  },
                },
              },
            },
          },
        },
        col2: {
          type: "void",
          "x-component": "Grid.Col",
          "x-component-props": {
            span: 8,
          },
          properties: {
            card1: {
              type: "void",
              "x-component": "Chart",
              items: {
                type: "void",
                properties: {
                  chart1: {
                    type: "void",
                    title: "UV",
                    "x-decorator": "ChartItem",
                    "x-decorator-props": {
                      request: {
                        operation: "getUv",
                      },
                    },
                    "x-component": "Chart.Line",
                    "x-component-props": {
                      meta: [
                        { id: "ds", name: "日期", isDim: true },
                        { id: "uv", name: "访客数" },
                      ],
                    },
                  },
                },
              },
            },
          },
        },
      },
    },
  },
};
