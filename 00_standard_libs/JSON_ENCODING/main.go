package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	//Create a map of key/value pairs.
	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "Programmer"
	c["content"] = map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.555.5555",
	}

	//Mashall the map into a json string
	data, err := json.MarshalIndent(c, "", " ") //1
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(data)) //2

}

/*
MarshalIndent function from json package eases the difficulty with pretty formatting,
especially when you publish a JSON document from a Go map or struct.

1. arguments are data, prefix and indent respectively.

2. data here is basiclaly []byte. So you have to convert them back to string.



WHAT ABOUT JSON response from Golang application?
You can set your content-type header

```
func SomeHandler(w http.ResponseWriter, r *http.Request) {
    data := SomeStruct{}
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(data)
}
```

*/
