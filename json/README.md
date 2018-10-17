JSON
----

JavaScript Object Notation (JSON) is a standard notation for sending and receiving structured information. JSON is not the only such notation. XML (§7.14), ASN.1, and Google’s Protocol Buffers serve similar purposes and each has its niche, but because of its simplicity, readability, and universal support, JSON is the most widely used.

Go has excellent support for encoding and decoding these formats, provided by the standard library packages encoding/json, encoding/xml, encoding/asn1, and so on, and these packages all have similar APIs. This section gives a brief overview of the most important parts of the encoding/json package.

JSON is an encoding of JavaScript values—strings, numbers, booleans, arrays, and objects—as Unicode text. It’s an efficient yet readable representation for the basic data types of Chapter 3 and the composite types of this chapter—arrays, slices, structs, and maps.

The basic JSON types are numbers (in decimal or scientific notation), booleans (true or false), and strings, which are sequences of Unicode code points enclosed in double quotes, with backslash escapes using a similar notation to Go, though JSON’s \Uhhhh numeric escapes denote UTF-16 codes, not runes.

These basic types may be combined recursively using JSON arrays and objects. A JSON array is an ordered sequence of values, written as a comma-separated list enclosed in square brackets; JSON arrays are used to encode Go arrays and slices. A JSON object is a mapping from strings to values, written as a sequence of name:value pairs separated by commas and surrounded by braces; JSON objects are used to encode Go maps (with string keys) and structs. For example:

```
boolean            true
number             -273.15
string             "She said \"Hello, Image\""
array              ["gold", "silver", "bronze"]
object             {"year": 1980,
                    "event": "archery",
                    "medals": ["gold", "silver", "bronze"]}
````
Consider an application that gathers movie reviews and offers recommendations. Its Movie data type and a typical list of values are declared below. (The string literals after the Year and Color field declarations are field tags; we’ll explain them in a moment.)

````
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
    // ...
}
````
Data structures like this are an excellent fit for JSON, and it’s easy to convert in both directions. Converting a Go data structure like movies to JSON is called marshaling. Marshaling is done by json.Marshal:

````
data, err := json.Marshal(movies)
if err != nil {
    log.Fatalf("JSON marshaling failed: %s", err)
}
fmt.Printf("%s\n", data)
````
Marshal produces a byte slice containing a very long string with no extraneous white space; we’ve folded the lines so it fits:

````
[{"Title":"Casablanca","released":1942,"Actors":["Humphrey Bogart","Ingr
id Bergman"]},{"Title":"Cool Hand Luke","released":1967,"color":true,"Ac
tors":["Paul Newman"]},{"Title":"Bullitt","released":1968,"color":true,"
Actors":["Steve McQueen","Jacqueline Bisset"]}]
````
This compact representation contains all the information but it’s hard to read. For human consumption, a variant called json.MarshalIndent produces neatly indented output. Two additional arguments define a prefix for each line of output and a string for each level of indentation:

````
data, err := json.MarshalIndent(movies, "", "    ")
if err != nil {
    log.Fatalf("JSON marshaling failed: %s", err)
}
fmt.Printf("%s\n", data)
````
The code above prints
````
[
    {
        "Title": "Casablanca",
        "released": 1942,
        "Actors": [
            "Humphrey Bogart",
            "Ingrid Bergman"
        ]
    },
    {
        "Title": "Cool Hand Luke",
        "released": 1967,
        "color": true,
        "Actors": [
            "Paul Newman"
        ]
    },
    {
        "Title": "Bullitt",
        "released": 1968,
        "color": true,
        "Actors": [
            "Steve McQueen",
            "Jacqueline Bisset"
        ]
    }
]
````
Marshaling uses the Go struct field names as the field names for the JSON objects . Only exported fields are marshaled, which is why we chose capitalized names for all the Go field names.

You may have noticed that the name of the Year field changed to released in the output, and Color changed to color. That’s because of the field tags. A field tag is a string of metadata associated at compile time with the field of a struct:

````
Year  int  `json:"released"`
Color bool `json:"color,omitempty"`
````
A field tag may be any literal string, but it is conventionally interpreted as a space-separated list of key:"value" pairs; since they contain double quotation marks, field tags are usually written with raw string literals. The json key controls the behavior of the encoding/json package, and other encoding/... packages follow this convention. The first part of the json field tag specifies an alternative JSON name for the Go field. Field tags are often used to specify an idiomatic JSON name like total_count for a Go field named TotalCount. The tag for Color has an additional option, omitempty, which indicates that no JSON output should be produced if the field has the zero value for its type (false, here) or is otherwise empty. Sure enough, the JSON output for Casablanca, a black-and-white movie, has no color field.

The inverse operation to marshaling, decoding JSON and populating a Go data structure, is called unmarshaling, and it is done by json.Unmarshal. The code below unmarshals the JSON movie data into a slice of structs whose only field is Title. By defining suitable Go data structures in this way, we can select which parts of the JSON input to decode and which to discard. When Unmarshal returns, it has filled in the slice with the Title information; other names in the JSON are ignored.

````
var titles []struct{ Title string }
if err := json.Unmarshal(data, &titles); err != nil {
    log.Fatalf("JSON unmarshaling failed: %s", err)
}
fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"
````
Many web services provide a JSON interface—make a request with HTTP and back comes the desired information in JSON format.