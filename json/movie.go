package json

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
Marshaling uses the Go struct field names as the field names for the JSON objects.
Only exported fields are marshaled, which is why we chose capitalized names for
all the Go field names.
*/
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

/*
Converting a Go data structure like movies to JSON is called marshaling.
Marshaling is done by json.Marshal. Marshal produces a byte slice containing
a very long string with no extraneous white space
*/
func Marshal() []byte {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	return data
}

func MarshalIndent() []byte {
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	return data
}

/*
The inverse operation to marshaling, decoding JSON and populating a Go data structure,
is called unmarshaling, and it is done by json.Unmarshal. The code below unmarshals the
JSON movie data into a slice of structs whose only field is Title. By defining suitable
Go data structures in this way, we can select which parts of the JSON input to decode
and which to discard. When Unmarshal returns, it has filled in the slice with the Title
information; other names in the JSON are ignored.
*/
func Unmarshal() {
	var titles []struct{ Title string }
	data := Marshal()
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"
}
