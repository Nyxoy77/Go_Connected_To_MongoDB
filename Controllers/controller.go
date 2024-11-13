package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	models "github.com/Nyxoy77/mongoDB/Models"
	pass "github.com/Nyxoy77/mongoDB/Password"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "netflix"
const colName = "watchlist"

// To actually store the refernce of the database
var collection *mongo.Collection

func init() {
	// This function is similar to init function in flutter where it is used to run the
	// initial functions of set up before doing the major changes or u can say operations

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	password := pass.ReturnPassword()
	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://Nyxoy:%s@cluster0.9tcsq.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0", password)).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	collection = client.Database(dbName).Collection(colName)

}

func insertOneMovie(movie models.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The movie with name %s was successfully inserteed with the id %d \n", movie.Movie, inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	// First convert the id from the string to the primitive
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The record was successfully updated ", result.MatchedCount)
}

//Well there are two bson.M and bson.D with same functionality but which to use when ? check notes

// Delete 1 record

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deletecount, _ := collection.DeleteOne(context.Background(), filter)
	fmt.Printf("The deleteccount is :  %d", deletecount.DeletedCount)
}

func deleteAllRecords() int64 {
	// filter := bson.D{{}} // dont take the extra space nigga
	deleteResult, _ := collection.DeleteMany(context.Background(), bson.D{{}})

	// if i want to delete all the records within the collection set the filter as
	// filter := bson.D{}

	return deleteResult.DeletedCount

}

func getAllMovies() []bson.M {
	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}
	var aa []bson.M
	for cursor.Next(context.Background()) {
		var entitiy bson.M

		if err := cursor.Decode(&entitiy); err != nil {
			log.Fatal(err)
		}
		aa = append(aa, entitiy)
		// It will automatically understand the incoming data and will fill the forms accordingly
		fmt.Println(entitiy)
	}
	defer cursor.Close(context.Background())
	return aa
}

// The actual helper functions now which we are gonna use in here
// Remember that these functions are to be exported
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	getAllMovies := getAllMovies()
	json.NewEncoder(w).Encode(getAllMovies)
}

func InsertoneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var mov models.Netflix

	if r.Body == nil {
		log.Fatal("The body should not be empty")
	}
	json.NewDecoder(r.Body).Decode(&mov)
	insertOneMovie(mov)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	count := deleteAllRecords()
	json.NewEncoder(w).Encode(count)
}
