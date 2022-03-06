import { Avatar, Divider, Dropdown, Menu } from "@arco-design/web-react";
import {
  IconLanguage,
  IconLock,
  IconPoweroff,
  IconSettings,
} from "@arco-design/web-react/icon";
import { useAdminLayoutContext } from "../../AdminLayoutProvider";
import styles from "../style/index.module.less";
import { openEditProfileDrawer } from "./EditProfile";

const ActionUser: React.FC = () => {
  const { profile = {} } = useAdminLayoutContext();

  return (
    <Dropdown
      droplist={
        <Menu
          onClickMenuItem={(key) => {
            console.log(key);
            if (key === "editprofile") {
              openEditProfileDrawer(profile);
            }
          }}
        >
          <Menu.Item key="editprofile">
            <IconSettings className={styles["dropdown-icon"]} />
            用户设置
          </Menu.Item>
          <Menu.Item key="changepwd">
            <IconLock className={styles["dropdown-icon"]} />
            修改密码
          </Menu.Item>
          <Menu.SubMenu
            key="language"
            title={
              <div>
                <IconLanguage className={styles["dropdown-icon"]} />
                切换语言
              </div>
            }
          >
            <Menu.Item key="lang-zhCN">中文</Menu.Item>
            <Menu.Item key="lang-enUS">English</Menu.Item>
          </Menu.SubMenu>
          <Divider style={{ margin: "4px 0" }} />
          <Menu.Item key="logout">
            <IconPoweroff className={styles["dropdown-icon"]} />
            退出登录
          </Menu.Item>
        </Menu>
      }
    >
      <Avatar size={32} style={{ cursor: "pointer" }}>
        <img
          alt="avatar"
          src={
            profile.avatar ||
            "https://lf1-xgcdn-tos.pstatp.com/obj/vcloud/vadmin/start.8e0e4855ee346a46ccff8ff3e24db27b.png"
          }
        />
      </Avatar>
    </Dropdown>
  );
};

export default ActionUser;
