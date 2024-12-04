package jsonapi

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"log"
	"maillist/src/mdb"
	"net/http"
)

func setJsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func fromJson[T any](body io.Reader, target T) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	json.Unmarshal(buf.Bytes(), &target)
}

func returnJson[T any](w http.ResponseWriter, withData func() (T, error)) {
	setJsonHeader(w)
	data, serverErr := withData()

	if serverErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		serverErrJson, err := json.Marshal(&serverErr)
		if err != nil {
			log.Println(err)
			return
		}
		w.Write(serverErrJson)
		return
	}

	jsonData, err := json.Marshal(&data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func returnErr(w http.ResponseWriter, err error, code int) {
	returnJson(w, func() (interface{}, error) {
		errorMessage := struct {
			Err string
		}{
			Err: err.Error(),
		}
		w.WriteHeader(code)
		return errorMessage, nil
	})
}

func CreateEmail(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		entry := mdb.EmailEntry{}
		fromJson(r.Body, &entry)
		if err := mdb.CreateEmail(db, entry.Email); err != nil {
			returnErr(w, err, http.StatusInternalServerError)
			return
		}
		returnJson(w, func() (interface{}, error) {
			log.Printf("Created email %v\n", entry.Email)
			return mdb.GetEmail(db, entry.Email)
		})
	})
}
func GetEmail(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		entry := mdb.EmailEntry{}
		fromJson(r.Body, &entry)
		returnJson(w, func() (interface{}, error) {
			log.Printf("Getemail %v\n", entry.Email)
			return mdb.GetEmail(db, entry.Email)
		})
	})
}

func UpdateEmail(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		entry := mdb.EmailEntry{}
		fromJson(r.Body, &entry)
		if err := mdb.UpdateEmail(db, &entry); err != nil {
			returnErr(w, err, http.StatusInternalServerError)
			return
		}
		returnJson(w, func() (interface{}, error) {
			log.Printf("Updated email %v\n", entry.Email)
			return mdb.GetEmail(db, entry.Email)
		})
	})
}
func DeleteEmail(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		entry := mdb.EmailEntry{}
		fromJson(r.Body, &entry)
		if err := mdb.DeleteEmail(db, entry.Email); err != nil {
			returnErr(w, err, http.StatusInternalServerError)
			return
		}
		returnJson(w, func() (interface{}, error) {
			log.Printf("Deleted email %v\n", entry.Email)
			return mdb.GetEmail(db, entry.Email)
		})
	})
}

func GetEmailBatch(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		queryOptions := mdb.GetEmailBatchQueryParams{}
		fromJson(r.Body, &queryOptions)
		if queryOptions.Count <= 0 || queryOptions.Page <= 0 {
			returnErr(w, errors.New("count and page must be greater than 0"), http.StatusBadRequest)
			return
		}
		returnJson(w, func() (interface{}, error) {
			log.Printf("Get email batch %v\n", queryOptions)
			return mdb.GetEmailBatch(db, &queryOptions)
		})
	})
}

func Serve(db *sql.DB, bind string) {
	http.Handle("/email/create", CreateEmail(db))
	http.Handle("/email/get", GetEmail(db))
	http.Handle("/email/get_batch", GetEmailBatch(db))
	http.Handle("/email/update", UpdateEmail(db))
	http.Handle("/email/delete", DeleteEmail(db))

	err := http.ListenAndServe(bind, nil)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
