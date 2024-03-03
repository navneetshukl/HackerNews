package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/hackernews-concurrency/helpers"
	"github.com/navneetshukl/hackernews-concurrency/models"
)

func GetNews(c *gin.Context) {

	start := time.Now()
	stories, err := helpers.GetStories()

	if err != nil {
		log.Println("Error in getting the Stories ", err)
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("some Error Occured.Please retry"))
		return
	}

	tmplData := models.TemplateData{
		Stories: stories,
		Time:    time.Now().Sub(start),
	}

	c.HTML(http.StatusOK, "index.page.tmpl", tmplData)

}
