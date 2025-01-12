package main

import (
	"log"
	"time"

	"um6p.ma/project-04/httpHandlers"
	service "um6p.ma/project-04/services"
	"um6p.ma/project-04/store"
)

func main() {

	store := store.Store{}
	store.Filepath = "./store.json"
	store.Schedule = 5 * time.Second // customize based on needs
	err := store.Load()
	if err != nil {
		log.Println(err)
		return
	}
	store.StartSchedule()
	log.Println("Storage Running and Scheduling is running")
	service := service.Service{}
	service.ReportDuration = 10 * time.Second // customize based on needs
	service.Init(&store)
	log.Println("Services are Running")

	server := httpHandlers.Server{Port: "8080"}
	err = server.StartServer(&service)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Server Running on")

}
