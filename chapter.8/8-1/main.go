package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

type Record struct {
	ProcessID string `json:"process_id"`
	DeletedAt JSTime `json:"deleted_at"`
}

type JSTime time.Time

func (t JSTime) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)
	if tt.IsZero() {
		return []byte("null"), nil
	}
	v := strconv.Itoa(int(tt.UnixMilli()))
	return []byte(v), nil
}

func (t *JSTime) UnmarshalJSON(data []byte) error {
	var jsonNumber json.Number
	err := json.Unmarshal(data, &jsonNumber)
	if err != nil {
		return err
	}
	unix, err := jsonNumber.Int64()
	if err != nil {
		return err
	}
	*t = JSTime(time.Unix(0, unix))
	return nil
}

func main() {
	r := &Record{
		ProcessID: "0001",
		DeletedAt: JSTime(time.Now()),
	}
	fmt.Printf("%+v\n", r)
	b, _ := json.Marshal(r)
	fmt.Println(string(b))

	var r2 *Record
	if err := json.Unmarshal([]byte(string(b)), &r2); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", time.Time(r2.DeletedAt).Format(time.RFC3339Nano))
}
