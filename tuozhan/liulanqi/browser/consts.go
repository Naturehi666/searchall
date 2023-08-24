package browser

import (
	"os"
)

// home dir path for all platforms
var homeDir, _ = os.UserHomeDir()

const (
	chromeName     = "Chrome"
	chromeBetaName = "Chrome Beta"
	chromiumName   = "Chromium"
	edgeName       = "Microsoft Edge"
	firefoxName    = "Firefox"
	speed360Name   = "360speed"
	qqBrowserName  = "QQ"
	sogouName      = "Sogou"
)
