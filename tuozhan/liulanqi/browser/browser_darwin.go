//go:build darwin

package browser

import (
	"searchall3.5/tuozhan/liulanqi/item"
)

var (
	chromiumList = map[string]struct {
		name        string
		storage     string
		profilePath string
		items       []item.Item
	}{
		"chrome": {
			name:        chromeName,
			storage:     chromeStorageName,
			profilePath: chromeProfilePath,
			items:       item.DefaultChromium,
		},
		"edge": {
			name:        edgeName,
			storage:     edgeStorageName,
			profilePath: edgeProfilePath,
			items:       item.DefaultChromium,
		},
		"chromium": {
			name:        chromiumName,
			storage:     chromiumStorageName,
			profilePath: chromiumProfilePath,
			items:       item.DefaultChromium,
		},
		"chrome-beta": {
			name:        chromeBetaName,
			storage:     chromeBetaStorageName,
			profilePath: chromeBetaProfilePath,
			items:       item.DefaultChromium,
		},
	}
	firefoxList = map[string]struct {
		name        string
		storage     string
		profilePath string
		items       []item.Item
	}{
		"firefox": {
			name:        firefoxName,
			profilePath: firefoxProfilePath,
			items:       item.DefaultFirefox,
		},
	}
)

var (
	chromeProfilePath     = homeDir + "/Library/Application Support/Google/Chrome/Default/"
	chromeBetaProfilePath = homeDir + "/Library/Application Support/Google/Chrome Beta/Default/"
	chromiumProfilePath   = homeDir + "/Library/Application Support/Chromium/Default/"
	edgeProfilePath       = homeDir + "/Library/Application Support/Microsoft Edge/Default/"
	arcProfilePath        = homeDir + "/Library/Application Support/Arc/User Data/Default"

	firefoxProfilePath = homeDir + "/Library/Application Support/Firefox/Profiles/"
)

const (
	chromeStorageName     = "Chrome"
	chromeBetaStorageName = "Chrome"
	chromiumStorageName   = "Chromium"
	edgeStorageName       = "Microsoft Edge"
	arcStorageName        = "Arc"
)
