package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/caseymrm/menuet"
)

func initMenu() {
	for {
		app := menuet.App()
		app.SetMenuState(&menuet.MenuState{
			Title: "Tunnels",
		})
		app.Label = "com.github.ryanshaw.gotunnels"
		app.Children = menuItems
		app.MenuChanged()
		time.Sleep(time.Second)
	}
}

func menuItems() []menuet.MenuItem {
	keys := make([]string, 0, len(tunnelStatus))
	for key := range tunnelStatus {
		keys = append(keys, key)
	}

	// Sort the keys
	sort.Strings(keys)

	items := []menuet.MenuItem{}
	for _, key := range keys {
		status := tunnelStatus[key]
		var text string
		if status.Active {
			text = fmt.Sprintf("%s Active", key)
		} else {
			text = fmt.Sprintf("%s Inactive", key)
		}
		items = append(items, menuet.MenuItem{
			Text: text,
		})
	}
	return items
}
