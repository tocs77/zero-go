package main

import (
	"database/sql"
	"log"
	"maillist/src/grpcapi"
	"maillist/src/jsonapi"
	"maillist/src/mdb"
	"sync"

	"github.com/alexflint/go-arg"
)

var args struct {
	DbPath   string `arg:"env:MAILLIST_DB_PATH"`
	BindJson string `arg:"env:MAILLIST_BIND_JSON"`
	BindGrpc string `arg:"env:MAILLIST_BIND_GRPC"`
}

func main() {
	arg.MustParse(&args)
	if args.DbPath == "" {
		args.DbPath = "list.db"
	}
	if args.BindJson == "" {
		args.BindJson = ":8080"
	}
	if args.BindGrpc == "" {
		args.BindGrpc = ":8081"
	}
	log.Printf("Using database %s\n", args.DbPath)
	log.Printf("Starting server on %s\n", args.BindJson)

	db, err := sql.Open("sqlite3", args.DbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	mdb.TryCreate(db)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		log.Printf("Starting jsonapi server...\n")
		jsonapi.Serve(db, args.BindJson)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		log.Printf("Starting grpcapi server...\n")
		grpcapi.Serve(db, args.BindGrpc)
		wg.Done()
	}()

	wg.Wait()
}
