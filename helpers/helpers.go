package helpers

import (
	"net/url"
	"strings"

	"github.com/navneetshukl/hackernews-concurrency/models"
)



func ParseItem(data models.Item) models.ModifiedItem {
	ret := models.ModifiedItem{
		Item: data,
	}

	url, err := url.Parse(ret.URL)
	if err == nil {
		ret.Host = strings.TrimPrefix(url.Hostname(), "www.")
	}
	return ret
}

func IsStoryLink(item models.ModifiedItem) bool {
	return item.Type == "story" && item.URL != ""
}
