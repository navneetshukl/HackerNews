package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/navneetshukl/hackernews-concurrency/models"
)

const (
	apiBase = "https://hacker-news.firebaseio.com/v0"
)

// Client is an API client used to interact with the Hacker News API
type Client struct {

	//unexported field

	apiBase string
}

// Making the Client zero value useful without forcing users to do something
// like `NewClient()`
func (c *Client) defaultify() {
	if c.apiBase == "" {
		c.apiBase = apiBase
	}
}

// TopItems returns the ids of roughly 450 top items in decreasing order. These
// should map directly to the top 450 things you would see on HN if you visited
// their site and kept going to the next page.
//
// TopItmes does not filter out job listings or anything else, as the type of
// each item is unknown without further API calls.

func (c *Client) TopItems() ([]int, error) {
	c.defaultify()

	resp, err := http.Get(fmt.Sprintf("%s/topstories.json", c.apiBase))

	if err != nil {
		return nil, err

	}

	defer resp.Body.Close()
	var ids []int

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&ids)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

// GetItem will return the Item defined by the provided ID.
func (c *Client) GetItem(id int) (models.Item, error) {
	c.defaultify()

	var item models.Item
	resp, err := http.Get(fmt.Sprintf("%s/item/%d.json", c.apiBase, id))
	if err != nil {
		return item, err
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&item)
	if err != nil {
		return item, err
	}
	return item, nil
}
