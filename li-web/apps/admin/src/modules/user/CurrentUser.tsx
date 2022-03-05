import { createContext, useState } from "react";
import { Button, Dropdown, Menu } from "@arco-design/web-react";
import { useTranslation } from "react-i18next";
import { useHistory } from "react-router-dom";
import { ChangePassword } from "./ChangePassword";
import { EditProfile } from "./EditProfile";
import { LanguageSettings } from "./LanguageSettings";
import { SwitchRole } from "./SwitchRole";
import { useCurrentUserContext } from "./CurrentUserProvider";
import { request } from "pro-utils";

export const DropdownVisibleContext = createContext<any>({});

export const CurrentUser = () => {
  const history = useHistory();
  const { t } = useTranslation();
  const [visible, setVisible] = useState(false);
  const { data } = useCurrentUserContext();
  return (
    <div style={{ display: "inline-block", verticalAlign: "top" }}>
      <DropdownVisibleContext.Provider value={{ visible, setVisible }}>
        <Dropdown
          popupVisible={visible}
          onVisibleChange={(visible) => {
            setVisible(visible);
          }}
          droplist={
            <Menu>
              <EditProfile />
              <ChangePassword />
              <SwitchRole />
              <LanguageSettings />
              <Menu.Item
                key="signout"
                onClick={() => {
                  request("userSignOut");
                  history.push("/signin");
                }}
              >
                {t("Sign out")}
              </Menu.Item>
            </Menu>
          }
        >
          <Button style={{ border: 0 }}>
            {data?.data?.nickname || data?.data?.email}
          </Button>
        </Dropdown>
      </DropdownVisibleContext.Provider>
    </div>
  );
};
