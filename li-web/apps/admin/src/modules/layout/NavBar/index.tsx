import Logo from "@/assets/logo.svg";
import ActionSystem from "./ActionSystem";
import ActionTheme from "./ActionTheme";
import ActionUser from "./ActionUser";
import styles from "./style/index.module.less";

const Navbar: React.FC = () => {
  return (
    <div className={styles.navbar}>
      <div className={styles.left}>
        <div className={styles.logo}>
          <Logo />
          <div className={styles["logo-name"]}>Li Admin</div>
        </div>
      </div>
      <ul className={styles.right}>
        <li>
          <ActionTheme />
        </li>
        <li>
          <ActionSystem />
        </li>
        <li>
          <ActionUser />
        </li>
      </ul>
    </div>
  );
};

export default Navbar;
