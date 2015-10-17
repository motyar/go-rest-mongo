// Example of REST API using MondoDb database
package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// The main structure for database
type Panda struct {
	Id   bson.ObjectId `json:"id" bson:"_id,omitempty"` // Mongodb _id
	Name string        `json:"name"`                    // Name of Panda
}

func main() {
    // Home page route
	http.HandleFunc("/", simpleHandler)
    // Route for API /pandas
	http.HandleFunc("/pandas/", apiHandler)

	bind := fmt.Sprintf("%s:%s", 127.0.0.1,8080)
	fmt.Printf("listening on %s...", bind)
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		panic(err)
	}
}

func apiHandler(res http.ResponseWriter, req *http.Request) {

    // Mongo db configurations
	session, err := mgo.Dial("mongodb://admin:password@127.5.241.2:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

    // Db and collection name configs
	c := session.DB("farm").C("pandas")

	var result []Panda
    // Get id form URL
	id := strings.Replace(req.URL.Path, "/pandas/", "", -1)

	//set mime type to JSON, Its JSON REST API
	res.Header().Set("Content-type", "application/json")

    // Handle the methods and behave accordingly
	switch req.Method {
	case "GET":
        // If no id passed in url, show them all out Pandas
		if id != "" {
			err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&result)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err = c.Find(nil).All(&result)
		}

	case "POST":
        // Read POST body from request
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

        // Convert body json to struct data
		var panda Panda
		err = json.Unmarshal(body, &panda)
		if err != nil {
			panic(err)
		}

        // We need a new mongodb _id to insert record, We are doing this becuase mongodb doesnt return last inserted record info
		i := bson.NewObjectId()
		panda.Id = i

        // Insert panda into farm.pandas
		err = c.Insert(panda)
		if err != nil {
			log.Fatal(err)
		}

        // Get details about just inserted row
		err = c.Find(bson.M{"_id": i}).All(&result)
		if err != nil {
			log.Fatal(err)
		}
	case "PUT":
        // Read POST body
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

        // Convert request json to struct
		var panda Panda
		err = json.Unmarshal(body, &panda)
		if err != nil {
			panic(err)
		}
        // We need a new mongoDb _id
		i := bson.ObjectIdHex(id)
		panda.Id = i

        // Update
		err = c.Update(bson.M{"_id": i}, panda)
		if err != nil {
			log.Fatal(err)
		}

        // Get info about just inserted document
		err = c.Find(bson.M{"_id": i}).All(&result)
		if err != nil {
			log.Fatal(err)
		}
	case "DELETE":
        // When a panda leaves :(, Delete from database
		err = c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
		if err != nil {
			log.Fatal(err)
		}
	default:
	}

    // Convert result struct to JSON
	json, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send the result JSON to the client.
	fmt.Fprintf(res, "%v", string(json))
}
func simpleHandler(res http.ResponseWriter, req *http.Request) {
	//set mime type to HTML
	res.Header().Set("Content-type", "text/html")
    // Guide them
	fmt.Fprintf(res, "Sir you are at wrong place!<br />Pandas are at <a href='/pandas/'>/pandas</a>")
}
