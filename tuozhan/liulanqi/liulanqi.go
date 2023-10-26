package liulanqi

import (
	"fmt"
	"os"
	"path/filepath"

	"searchall3.5/tuozhan/liulanqi/browser"
	"searchall3.5/tuozhan/liulanqi/log"
	"searchall3.5/tuozhan/liulanqi/utils/fileutil"
)

var (
	outputDir    string
	outputFormat string
	isFullExport bool
)

func Execute(browserFlag string, profilePath string) {

	outputDir = "results"
	outputFormat = "csv"

	isFullExport = true

	browsers, err := browser.PickBrowsers(browserFlag, profilePath)

	if err != nil {
		log.Error(err)

	}

	for _, b := range browsers {
		data, err := b.BrowsingData(isFullExport)
		if err != nil {
			log.Error(err)
			continue
		}
		data.Output(outputDir, b.Name(), outputFormat)
	}

	if _, err := os.Stat(outputDir); err == nil {
		log.Notice("Current directory generation folder:results")
	} else {

	}
}

func CompressResult() error {
	if err := fileutil.CompressDir(outputDir); err != nil {
		return fmt.Errorf("Compression failed：%s", err.Error())
	}

	dir, _ := os.Getwd()
	log.Noticef("Generate compressed package: %s", filepath.Join(dir, "\\results\\results.zip"))

	return nil
}
