import { TabsProps, Tabs as ArcoTabs } from "@arco-design/web-react";
import { RecursionField, useFieldSchema } from "@formily/react";
import { Fragment } from "react";
import { Icon } from "../__builtins__";

type ComposedTabs = React.FC<TabsProps> & {
  TabPane?: React.FC;
};

export const Tabs: ComposedTabs = (props) => {
  const fieldSchema = useFieldSchema();
  return (
    <ArcoTabs {...props}>
      {fieldSchema.mapProperties((schema, key) => {
        return (
          <ArcoTabs.TabPane
            key={key}
            title={
              schema["x-component-props"]?.icon ? (
                <span>
                  <Icon
                    type={schema["x-component-props"]?.icon}
                    style={{ marginRight: 6 }}
                  />
                  {schema.title}
                </span>
              ) : (
                schema.title
              )
            }
          >
            <RecursionField name={key} schema={schema} onlyRenderProperties />
          </ArcoTabs.TabPane>
        );
      })}
    </ArcoTabs>
  );
};

Tabs.TabPane = () => {
  return <Fragment />;
};

export default Tabs;
