package main

import (
	"github.com/jolycky/task-5-pbi-btpns-JoseBagusRamadhan/database"
	"github.com/jolycky/task-5-pbi-btpns-JoseBagusRamadhan/models"
	"github.com/jolycky/task-5-pbi-btpns-JoseBagusRamadhan/router"
)

func main() {
	db := database.SetupDB()
	db.AutoMigrate(&models.User{})

	r := router.SetupRoutes(db)
	r.Run()
}
