package main

import (
	"github.com/jared-heidt/swe-influencer-api/controllers"
	"github.com/jared-heidt/swe-influencer-api/models"
)

func main() {
	models.Init()
	controllers.Start()
}
