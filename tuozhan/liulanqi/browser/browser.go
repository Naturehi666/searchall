package browser

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"searchall3.5/tuozhan/liulanqi/browingdata"
	"searchall3.5/tuozhan/liulanqi/browser/chromium"
	"searchall3.5/tuozhan/liulanqi/browser/firefox"
	"searchall3.5/tuozhan/liulanqi/log"
	"searchall3.5/tuozhan/liulanqi/utils/fileutil"
	"searchall3.5/tuozhan/liulanqi/utils/typeutil"
)

type Browser interface {
	// Name is browser's name
	Name() string
	// BrowsingData returns all browsing data in the browser.
	BrowsingData(isFullExport bool) (*browingdata.Data, error)
}

// PickBrowsers returns a list of browsers that match the name and profile.
func PickBrowsers(name, profile string) ([]Browser, error) {
	var browsers []Browser
	clist := pickChromium(name, profile)
	for _, b := range clist {
		if b != nil {
			browsers = append(browsers, b)

		}
	}
	flist := pickFirefox(name, profile)
	for _, b := range flist {
		if b != nil {
			browsers = append(browsers, b)

		}
	}
	return browsers, nil
}

func pickChromium(name, profile string) []Browser {
	var browsers []Browser
	name = strings.ToLower(name)
	if name == "all" {
		for _, v := range chromiumList {
			if !fileutil.IsDirExists(filepath.Clean(v.profilePath)) {
				log.Noticef("find browser %s failed, profile folder does not exist", v.name)
				continue
			}
			multiChromium, err := chromium.New(v.name, v.storage, v.profilePath, v.items)
			if err != nil {
				log.Errorf("new chromium error: %v", err)
				continue
			}
			for _, b := range multiChromium {

				log.Noticef("find browser %s success", b.Name())
				browsers = append(browsers, b)
			}
		}
	}

	if c, ok := chromiumList[name]; ok {

		if profile == "" {
			profile = c.profilePath
		}

		defer func() {
			if r := recover(); r != nil {

				os.Exit(1) // 正常退出程序，状态码为 1
			}
		}()

		if !fileutil.IsDirExists(filepath.Clean(profile)) {

			log.Fatalf("find browser %s failed, profile folder does not exist", name)

		}
		chromiumList, err := chromium.New(c.name, c.storage, profile, c.items)
		if err != nil {
			log.Fatalf("new chromium error: %s", err)
		}
		for _, b := range chromiumList {
			log.Noticef("find browser %s success", b.Name())

			browsers = append(browsers, b)
		}
	}

	return browsers
}

func pickFirefox(name, profile string) []Browser {
	var browsers []Browser
	name = strings.ToLower(name)
	if name == "all" || name == "firefox" {
		for _, v := range firefoxList {
			if profile == "" {
				profile = v.profilePath
			} else {
				profile = fileutil.ParentDir(profile)
			}

			if !fileutil.IsDirExists(filepath.Clean(profile)) {
				log.Noticef("find browser firefox %s failed, profile folder does not exist", v.name)
				continue
			}

			if multiFirefox, err := firefox.New(profile, v.items); err == nil {
				for _, b := range multiFirefox {
					log.Noticef("find browser firefox %s success", b.Name())
					browsers = append(browsers, b)
				}
			} else {
				log.Error(err)
			}
		}

		return browsers
	}

	return nil
}

func ListBrowsers() []string {
	var l []string
	l = append(l, typeutil.Keys(chromiumList)...)
	l = append(l, typeutil.Keys(firefoxList)...)
	sort.Strings(l)
	return l
}

func Names() string {
	return strings.Join(ListBrowsers(), "|")
}
