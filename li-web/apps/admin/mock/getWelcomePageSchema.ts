export default {
  type: "object",
  properties: {
    row1: {
      type: "void",
      "x-component": "Grid.Row",
      "x-component-props": {
        gutter: 20,
      },
      properties: {
        col1: {
          type: "void",
          "x-component": "Grid.Col",
          "x-component-props": {
            span: 8,
          },
          properties: {
            card1: {
              type: "void",
              "x-component": "Card",
              "x-component-props": {
                title: "Li Card 1",
              },
              "x-content": "Li Card Content 1",
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
              "x-component": "Card",
              "x-component-props": {
                title: "Li Card 2",
              },
              "x-content": "Li Card Content 2",
            },
          },
        },
        col2_1: {
          type: "void",
          "x-component": "Grid.Col",
          "x-component-props": {
            span: 8,
          },
          properties: {
            card1: {
              type: "void",
              "x-component": "Card",
              "x-component-props": {
                title: "Li Card 2-1",
              },
              "x-content": "Li Card Content 2-1",
            },
          },
        },
      },
    },
    row2: {
      type: "void",
      "x-component": "Grid.Row",
      "x-component-props": {
        gutter: 20,
      },
      properties: {
        col1: {
          type: "void",
          "x-component": "Grid.Col",
          "x-component-props": {
            span: 16,
          },
          properties: {
            card1: {
              type: "void",
              "x-component": "Card",
              "x-component-props": {
                title: "Li Card 3",
              },
              "x-content": "Li Card Content 3",
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
              "x-component": "Card",
              "x-component-props": {
                title: "Li Card 4",
              },
              "x-content": "Li Card Content 4",
            },
          },
        },
      },
    },
  },
};
