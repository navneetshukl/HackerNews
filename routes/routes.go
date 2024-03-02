package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/hackernews-concurrency/client"
	"github.com/navneetshukl/hackernews-concurrency/helpers"
	"github.com/navneetshukl/hackernews-concurrency/models"
)

const (
	numStories = 30
)

func GetNews(c *gin.Context) {

	start := time.Now()

	var client client.Client
	ids, err := client.TopItems()
	if err != nil {
		log.Println("Error in getting top items from hacker news ", err)
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Some Error occured.Please Retry Again!"))
		return

	}

	stories := []models.ModifiedItem{}

	for _, id := range ids {
		hnItem, err := client.GetItem(id)
		if err != nil {
			continue
		}
		item := helpers.ParseItem(hnItem)

		if helpers.IsStoryLink(item) {
			stories = append(stories, item)

			if len(stories) >= numStories {
				break
			}
		}
	}

	tmplData := models.TemplateData{
		Stories: stories,
		Time:    time.Now().Sub(start),
	}

	c.HTML(http.StatusOK, "index.page.tmpl", tmplData)

}
