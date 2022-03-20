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
              },
              properties: {
                actions: {
                  type: "void",
                  "x-component": "List.Action",
                  properties: {
                    refresh: {
                      type: "void",
                      "x-component": "List.Action.Refresh",
                      "x-component-props": {
                        position: "left",
                      },
                    },
                    add: {
                      type: "void",
                      title: "新建用户",
                      "x-component": "List.Action.RecordEditDrawer",
                      "x-component-props": {
                        forSubmit: "updateUser",
                        type: "primary",
                      },
                      properties: {
                        nickname: {
                          type: "string",
                          title: "昵称",
                          "x-decorator": "FormItem",
                          "x-component": "Input",
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
                userlisttable: {
                  type: "array",
                  "x-component": "List.Table",
                  "x-component-props": {
                    rowKey: "id",
                    rowSelection: {
                      type: "checkbox",
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
                            "x-component": "Input",
                            "x-read-pretty": true,
                          },
                        },
                      },
                      column3: {
                        type: "void",
                        "x-component": "List.Table.Column",
                        "x-component-props": {
                          title: "操作",
                          dataIndex: "__action",
                        },
                        properties: {
                          recordactions: {
                            type: "void",
                            "x-component": "Space",
                            properties: {
                              edit: {
                                type: "void",
                                title: "编辑",
                                "x-component": "List.Action.RecordEditDrawer",
                                "x-component-props": {
                                  forInit: "getUser",
                                  forSubmit: "updateUser",
                                },
                                properties: {
                                  nickname: {
                                    type: "string",
                                    title: "昵称",
                                    "x-decorator": "FormItem",
                                    "x-component": "Input",
                                  },
                                },
                              },
                              delete: {
                                type: "void",
                                title: "删除",
                                "x-component": "List.Action.RecordDelete",
                                "x-component-props": {
                                  status: "danger",
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
              "x-component": "CardItem",
              "x-component-props": {
                title: "Li Card 2",
              },
              "x-content": "Li Card Content 2",
            },
          },
        },
      },
    },
  },
};