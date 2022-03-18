package view

import (
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/com"
)

func NavItems() []view.Node {
	return []view.Node{
		com.ThemeSwitch("navThemeSwitch"),
	}
}
