package main

import (
	"database/sql"
	"log"
	"maillist/src/jsonapi"
	"maillist/src/mdb"
	"sync"

	"github.com/alexflint/go-arg"
)

var args struct {
	DbPath   string `arg:"env:MAILLIST_DB_PATH"`
	BindJson string `arg:"env:MAILLIST_BIND_JSON"`
}

func main() {
	arg.MustParse(&args)
	if args.DbPath == "" {
		args.DbPath = "list.db"
	}
	if args.BindJson == "" {
		args.BindJson = ":8080"
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
	wg.Wait()
}
