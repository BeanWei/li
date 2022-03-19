import { Dropdown, Menu } from "@arco-design/web-react";
import {
  RecursionField,
  Schema,
  useField,
  useFieldSchema,
} from "@formily/react";
import { observer } from "@formily/reactive-react";

type ComposedDropdownMenu = React.FC<any> & {
  Item?: React.FC<any>;
  SubMenu?: React.FC<any>;
  URL?: React.FC<any>;
};

export const DropdownMenu: ComposedDropdownMenu = observer((props) => {
  const schema = useFieldSchema();
  return (
    <Dropdown
      {...props}
      // FIXME: 会影响弹窗的显示
      unmountOnExit={false}
      droplist={
        <Menu>
          <RecursionField
            schema={schema.items as Schema}
            onlyRenderProperties
          />
        </Menu>
      }
    >
      <div>
        <RecursionField schema={schema} onlyRenderProperties />
      </div>
    </Dropdown>
  );
});

DropdownMenu.Item = observer((props) => {
  const schema = useFieldSchema();
  const field = useField();
  return (
    <Menu.Item {...props} key={schema.name as string}>
      {schema.properties ? (
        <RecursionField schema={schema} onlyRenderProperties />
      ) : (
        field.title
      )}
    </Menu.Item>
  );
});

DropdownMenu.SubMenu = observer((props) => {
  const schema = useFieldSchema();
  const field = useField();
  return (
    <Menu.SubMenu {...props} key={schema.name as string} title={field.title}>
      <RecursionField schema={schema} onlyRenderProperties />
    </Menu.SubMenu>
  );
});

DropdownMenu.URL = observer((props) => {
  const schema = useFieldSchema();
  const field = useField();
  return (
    <Menu.Item
      key={schema.name as string}
      onClick={() => {
        window.open(props.href, "_blank");
      }}
    >
      {field.title}
    </Menu.Item>
  );
});
