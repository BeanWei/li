import Logo from "@/assets/logo.svg";
import { Layout } from "@arco-design/web-react";
import { SchemaComponent } from "schema-components";
import { useRoute } from "../route-switch/hooks";
import styles from "./index.module.less";

export const SignPage: React.FC = (props) => {
  const route = useRoute();
  const {
    title = "Li Admin",
    subTitle,
    footer = "Li Admin",
    logo,
    body = {},
  } = route.config;

  return (
    <div className={styles.container}>
      <div className={styles.top}>
        <div className={styles.header}>
          <span className={styles.logo}>
            {logo ? <img src={logo} /> : <Logo />}
          </span>
          <span className={styles.title}>{title}</span>
        </div>
        {subTitle ? <div className={styles.desc}>{subTitle}</div> : null}
      </div>
      <div className={styles.content}>
        <SchemaComponent schema={body} />
      </div>
      <Layout.Footer className={styles.footer}>{footer}</Layout.Footer>
    </div>
  );
};

export default SignPage;
