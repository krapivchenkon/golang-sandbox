package main

import (
	"encoding/json"
	"fmt"
	//"reflect"
	"time"
)

// start with a string representation of our JSON data
var input = `
{
    "created_at": "Thu May 31 00:00:01 +0000 2012",
    "str_tst": "str_slice",
    "id": 123
}
`

type msg struct {
	Created Timestamp `json:"created_at"`
	Id      int       `json:"id"`
	Desc    CustMarsh `json:"str_tst"`
}

type Timestamp time.Time

// func (t *Timestamp) Format

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	v, err := time.Parse(time.RubyDate, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	*t = Timestamp(v)
	return nil
}

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	v := time.Time(*t).Format(time.RubyDate)

	fmt.Printf("ruby date:%v\n", v)
	fmt.Println([]byte(v))
	// don't forget to wrap value with quotes
	return []byte(fmt.Sprintf("\"%v\"", v)), nil
}

type CustMarsh string

func (t *CustMarsh) UnmarshalJSON(b []byte) error {

	*t = CustMarsh(string(b[1:4]))
	return nil
}

func main() {
	// our target will be of type map[string]interface{}, which is a pretty generic type
	// that will give us a hashtable whose keys are strings, and whose values are of
	// type interface{}
	var val msg //map[string]Timestamp

	if err := json.Unmarshal([]byte(input), &val); err != nil {
		panic(err)
	}

	fmt.Println(val)
	//for k, v := range val {
	//  fmt.Println(k, reflect.TypeOf(v))
	//}
	fmt.Println(time.Time(val.Created))

	fmt.Println(time.Time(Timestamp(time.Now())).Format(time.RubyDate))

	// marshal sample
	res1D := &msg{
		Id:      432,
		Created: Timestamp(time.Now()),
		Desc:    "sdf",
	}

	res1B, err := json.Marshal(res1D)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res1B))
}
