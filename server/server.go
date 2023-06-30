package server

import (
	"log"

	"github.com/KhoirulAziz99/mnc/api"
	"github.com/KhoirulAziz99/mnc/config"
	"github.com/KhoirulAziz99/mnc/pkg/utility"
	
)


func Run() error {
	db, err := config.InitDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := api.SetupRouter(db)
	serverAddress := utility.GetEnv("SERVER_ADDRESS")
	log.Printf("Server running in port %s", serverAddress)

	err = router.Run(serverAddress)
	if err != nil {
		return err
	}

	return nil
}