package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client
var db *mongo.Database

type Record struct {
	GUID string                 `json:"guid"`
	Data map[string]interface{} `json:"data"`
}

func initDB() {
	var err error
	mongoURI := fmt.Sprintf("mongodb://%s:%s@mongo:27017",
		os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
		os.Getenv("MONGO_INITDB_ROOT_PASSWORD"))

	client, err = mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db = client.Database(os.Getenv("MONGO_DB"))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func changeDBConnection(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		ConnectionString string `json:"connectionString"`
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(req.ConnectionString))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	db = client.Database(os.Getenv("MONGO_DB"))
	w.WriteHeader(http.StatusOK)
}

func createRecord(w http.ResponseWriter, r *http.Request) {
	var record Record
	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	record.GUID = fmt.Sprintf("%d", time.Now().UnixNano())

	_, err := db.Collection("records").InsertOne(context.Background(), record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func getRecords(w http.ResponseWriter, r *http.Request) {
	cursor, err := db.Collection("records").Find(context.Background(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var records []Record
	if err := cursor.All(context.Background(), &records); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
}

func deleteRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	_, err := db.Collection("records").DeleteOne(context.Background(), bson.M{"guid": guid})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func updateRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	var record Record
	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Collection("records").UpdateOne(context.Background(), bson.M{"guid": guid}, bson.M{"$set": record})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	initDB()

	r := mux.NewRouter()
	r.HandleFunc("/api/health", healthCheck).Methods("GET")
	r.HandleFunc("/api/change-db", changeDBConnection).Methods("POST")
	r.HandleFunc("/api/record", createRecord).Methods("POST")
	r.HandleFunc("/api/records", getRecords).Methods("GET")
	r.HandleFunc("/api/record/{guid}", deleteRecord).Methods("DELETE")
	r.HandleFunc("/api/record/{guid}", updateRecord).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", r))
}