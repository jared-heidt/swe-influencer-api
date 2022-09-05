package main

import (
	"github.com/jared-heidt/swe-influencer-api/api/controllers"
	"github.com/jared-heidt/swe-influencer-api/api/models"
)

func main() {
	models.Init()
	controllers.Start()
}
