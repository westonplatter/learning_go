package main

import "encoding/json"
import "fmt"
import "os"

type R1 struct {
	Page   int
	Fruits []string
}

// why the ticks, "`"?
// magick
// the ticks allow you to decode/UnMarshal json data directly into structs
// yep, magick
type R2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	// marshall data
	r1 := &R1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	rm1, _ := json.Marshal(r1)
	fmt.Println(rm1)
	// woo! -> [123 34 80 97 103 101 34 58 49 44 34 70 114 117 105 116 115 34 58 91 34 97 112 112 108 101 34 44 34 112 101 97 99 104 34 44 34 112 101 97 114 34 93 125]

	// unmarshall data
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// to take care of decoding into proper types
	num := dat["num"].(float64)
	fmt.Println(num)

	// access nested data
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(strs) // [a b]   Array
	fmt.Println(str1) // a       String

	// the real magic
	// using the R2 struct with the ticks, "`"
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &R2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// more fancy stuff
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
