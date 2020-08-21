package controllers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"reactgomongo/models"
)

func AddTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	SetHeaders(w)

	var todo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)

	if todo.Title == ""{
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}
	if todo.Description == ""{
		http.Error(w, "Description is required", http.StatusBadRequest)
		return
	}

	todo.Completed = false

	insertTodo, err := todoscol.InsertOne(context.TODO(), todo)

	if err != nil {
		log.Fatal(err)
	}
	todo.ID = insertTodo.InsertedID.(primitive.ObjectID)
	json.NewEncoder(w).Encode(todo)

}

func GetAllTodos(w http.ResponseWriter, r *http.Request){
	SetHeaders(w)
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	// Filter only for todos that are incomplete
	filter := bson.M{"completed": false}
	cur, err := todoscol.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	results := make([]primitive.M, 0)
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		// fmt.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	json.NewEncoder(w).Encode(results)
}


func GetTodo(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	var todo models.Todo

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}

	err := todoscol.FindOne(context.TODO(), filter).Decode(&todo)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(todo)
}


func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	// Simulating putting response into struct
	// Would normally get updates from here instead of setting completed to true
	var todo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}
	update := bson.D{
		{"$set", bson.D{
			{"completed", true},
		}}}
	_, err := todoscol.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	todo.ID = id

	json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request){
	SetHeaders(w)
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	_, err := todoscol.DeleteOne(context.TODO(), bson.M{"_id": id})

	if err != nil {
		log.Fatal(err)
	}

	type Response struct {
		Message string `json:"message" bson:"message"`
	}
	response := Response{
		Message: "Successfully Deleted",
	}
	json.NewEncoder(w).Encode(response)
}
