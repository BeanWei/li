import Logo from "@/assets/logo.svg";
import { Layout } from "@arco-design/web-react";
import { useRequest } from "pro-utils";
import { SchemaComponent } from "schema-components";
import { Loading } from "../components";
import styles from "./index.module.less";

export const SignPage: React.FC = (props) => {
  const { data = {}, loading } = useRequest("@getSignView");

  if (loading) {
    return <Loading />;
  }

  const {
    title = "Li Admin",
    description,
    copyright = "Li Admin",
    logo,
    body = {},
  } = data;

  return (
    <div className={styles.container}>
      <div className={styles.top}>
        <div className={styles.header}>
          <span className={styles.logo}>
            {logo ? <img src={logo} /> : <Logo />}
          </span>
          <span className={styles.title}>{title}</span>
        </div>
        {description ? <div className={styles.desc}>{description}</div> : null}
      </div>
      <div className={styles.content}>
        <SchemaComponent schema={body} />
      </div>
      <Layout.Footer className={styles.footer}>{copyright}</Layout.Footer>
    </div>
  );
};

export default SignPage;
