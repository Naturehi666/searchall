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
		"opera": {
			name:        operaName,
			profilePath: operaProfilePath,
			storage:     operaStorageName,
			items:       item.DefaultChromium,
		},
		"opera-gx": {
			name:        operaGXName,
			profilePath: operaGXProfilePath,
			storage:     operaStorageName,
			items:       item.DefaultChromium,
		},
		"vivaldi": {
			name:        vivaldiName,
			storage:     vivaldiStorageName,
			profilePath: vivaldiProfilePath,
			items:       item.DefaultChromium,
		},
		"coccoc": {
			name:        coccocName,
			storage:     coccocStorageName,
			profilePath: coccocProfilePath,
			items:       item.DefaultChromium,
		},
		"brave": {
			name:        braveName,
			profilePath: braveProfilePath,
			storage:     braveStorageName,
			items:       item.DefaultChromium,
		},
		"yandex": {
			name:        yandexName,
			storage:     yandexStorageName,
			profilePath: yandexProfilePath,
			items:       item.DefaultYandex,
		},
		"arc": {
			name:        arcName,
			profilePath: arcProfilePath,
			storage:     arcStorageName,
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
	braveProfilePath      = homeDir + "/Library/Application Support/BraveSoftware/Brave-Browser/Default/"
	operaProfilePath      = homeDir + "/Library/Application Support/com.operasoftware.Opera/Default/"
	operaGXProfilePath    = homeDir + "/Library/Application Support/com.operasoftware.OperaGX/Default/"
	vivaldiProfilePath    = homeDir + "/Library/Application Support/Vivaldi/Default/"
	coccocProfilePath     = homeDir + "/Library/Application Support/Coccoc/Default/"
	yandexProfilePath     = homeDir + "/Library/Application Support/Yandex/YandexBrowser/Default/"
	arcProfilePath        = homeDir + "/Library/Application Support/Arc/User Data/Default"

	firefoxProfilePath = homeDir + "/Library/Application Support/Firefox/Profiles/"
)

const (
	chromeStorageName     = "Chrome"
	chromeBetaStorageName = "Chrome"
	chromiumStorageName   = "Chromium"
	edgeStorageName       = "Microsoft Edge"
	braveStorageName      = "Brave"
	operaStorageName      = "Opera"
	vivaldiStorageName    = "Vivaldi"
	coccocStorageName     = "CocCoc"
	yandexStorageName     = "Yandex"
	arcStorageName        = "Arc"
)
