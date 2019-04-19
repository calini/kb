package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/pretty"
	"io/ioutil"
	"kb/machine"
	"kb/store"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := gin.Default()

	router.Static("/static/js", "./static/js")
	router.Static("/static/css", "./static/css")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")
	router.StaticFile("/", "./static/index.html")

	db := store.NewCacheStore()

	// cache received status
	router.POST("/report", func(c *gin.Context) {
		report := make(machine.LogReport)
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Errorf("Could not read request body: %s", err)
		}
		err = json.Unmarshal(body, &report)
		if err != nil {
			log.Errorf("Could not parse received JSON to cache: %s", err)
		}

		fmt.Printf("Read:\n%s", pretty.Pretty(body))
		db.SaveSnapshot(machine.ConvertReportToSnapshot(report))
	})

	// API
	api := router.Group("/api")
	{
		api.GET("/status", func(c *gin.Context) {
			c.JSON(200, db.GetSnapshot())
		})

		api.GET("/status/machines/:machine", func(c *gin.Context) {
			machine := c.Param("machine")
			c.JSON(200, db.GetSnapshotForMachine(machine))
		})

		api.GET("/status/labs/:lab", func(c *gin.Context) {
			lab := c.Param("lab")
			c.JSON(200, db.GetSnapshotForLab(lab))
		})
	}

	log.Fatal(router.Run(":8080")) // listen and serve on 0.0.0.0:8080
}
