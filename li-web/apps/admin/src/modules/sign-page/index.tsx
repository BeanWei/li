import Logo from "@/assets/logo.svg";
import { Layout } from "@arco-design/web-react";
import { SchemaComponent } from "schema-components";
import { useRoute } from "../route-switch/hooks";

export const SignPage: React.FC = (props) => {
  const route = useRoute();
  const {
    title = "Li Admin",
    footer = "Li Admin",
    logo,
    body = {},
  } = route.config;

  return (
    <Layout>
      <Layout.Header>
        <div>
          {logo ? <img src={logo} /> : <Logo />}
          <div style={{ fontSize: "20px" }}>{title}</div>
        </div>
      </Layout.Header>
      <Layout.Content>
        <div>
          <SchemaComponent schema={body} />
        </div>
      </Layout.Content>
      <Layout.Footer>{footer}</Layout.Footer>
    </Layout>
  );
};

export default SignPage;
