package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MastersAcademy/ma_ap_2016/Go/Creational/AbstractFactory/example1_factoryMethod/dataStore"
)

//postgres "dbname=tft user=postgres password=working host=localhost sslmode=disable" system_user id username

func main(){

	manager := dataStore.NewDataStoreManager()

	params := []string{"DATASTORE", "DATASTORE_POSTGRES_DSN", "DB_TABLE", "DB_ID_FIELD", "DB_USERNAME_FIELD"}
	args := os.Args[1:]
	if len(args) == 0 {
		log.Printf("You must privide datasource name\n")
	}
	if args[0] == "postgres" && len(args) < 5 {
		log.Printf("Not enogh args for postres data source")
	}
	conf := dataStore.NewConfiguration()
	for k, v := range args{
		conf.Set(params[k], v)
	}
	store, err := manager.CreateDatastore(conf)
	if err != nil{
		log.Printf("Error: %v", err)
		return
	}
	name, err := store.FindUserNameById(1)
	if err != nil{
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Username: %s\n", name)
}