package main

import (
	"os"

	"github.com/joho/godotenv"

	svr "github.com/whattheearl/fig/profilesvc/internal/grpc"
	db "github.com/whattheearl/fig/profilesvc/internal/mongodb"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
}

func main() {
	mongo_usr := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	mongo_pwd := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	mongo_addr := os.Getenv("MONGO_ADDR")
	mongo_db := os.Getenv("PROFILE_DB")
	mongo_col := os.Getenv("PROFILE_COLLECTION")
	db.Connect(mongo_usr, mongo_pwd, mongo_addr, mongo_db, mongo_col)
	addr := os.Getenv("PROFILE_ADDR")
	svr.Run(addr)
}
