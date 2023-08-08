package main

import ( "fmt"
	"encoding/json"
	"encoding/xml"
)

type Container struct {
	Field1 string `xml:"id,attr"`
	Field2 string `xml:"parent1>paren2"`
	XField string
}

func main () {
	f := Container {"whatever", "second whatever", "s"}

	if in_bytes, err := json.Marshal(f); err == nil {
		fmt.Println(string(in_bytes))
		f2 := Container{}
		x := []byte("{\"XField\": \"test\"}")
		fmt.Println(string(x))
		json.Unmarshal(x, &f2)
		fmt.Println(f2)
	} else {
		fmt.Println(err)
	}

	if in_bytes, err := xml.Marshal(f); err == nil {
		fmt.Println(string(in_bytes))
		var f2 Container
		xml.Unmarshal(in_bytes, &f2)
		fmt.Println(f2)
	}
}