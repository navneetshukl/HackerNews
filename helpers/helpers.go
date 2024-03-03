package helpers

import (
	"log"
	"net/url"
	"sort"
	"strings"

	"github.com/navneetshukl/hackernews-concurrency/client"
	"github.com/navneetshukl/hackernews-concurrency/models"
)

const (
	numStories = 30
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

// IsStoryLink function check if the given item type is "story"
func IsStoryLink(item models.ModifiedItem) bool {
	return item.Type == "story" && item.URL != ""
}

// GetStories function is a goroutine for getting the stories
func GetStories() ([]models.ModifiedItem, error) {
	var client client.Client
	ids, err := client.TopItems()
	if err != nil {
		log.Println("Error in getting top items from hacker news ", err)

		return nil, err

	}
	var stories []models.ModifiedItem

	type result struct {
		idx  int
		item models.ModifiedItem
		err  error
	}

	resultCh := make(chan result)
	for index := 0; index < numStories; index++ {

		go func(index, id int) {

			hnItem, err := client.GetItem(id)
			if err != nil {
				resultCh <- result{idx: index, err: err, item: ParseItem(hnItem)}
			} else {
				resultCh <- result{idx: index, item: ParseItem(hnItem), err: nil}
			}

		}(index, ids[index])
	}

	var results []result

	for i := 0; i < numStories; i++ {
		results = append(results, <-resultCh)
	}
	close(resultCh)
	sort.Slice(results, func(i, j int) bool {
		return results[i].idx < results[j].idx
	})

	for _, res := range results {
		if res.err != nil {
			continue
		}
		if IsStoryLink(res.item) {
			stories = append(stories, res.item)
		}

	}

	return stories, nil

}
