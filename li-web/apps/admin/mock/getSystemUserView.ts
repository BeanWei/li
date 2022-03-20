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
