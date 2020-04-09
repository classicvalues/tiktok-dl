package workflows

import (
	"regexp"

	client "../client"
	log "../utils/log"
)

// CanUseDownloadShareLink - Check's if DownloadShareLink can be used
func CanUseDownloadShareLink(url string) bool {
	match, _ := regexp.MatchString("vm.tiktok.com\\/.+", url)
	return match
}

// DownloadShareLink - Download item by share link
func DownloadShareLink(url string) {
	log.Logf("Resolving share link: %s\n", url)

	finalURL, err := client.GetRedirectUrl(url)
	if err != nil {
		OnWorkflowFail(err, url)
		return
	}

	StartWorkflowByParameter(finalURL)
}
