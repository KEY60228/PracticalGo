package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Response struct {
	Type      string          `json:"type"`
	Timestamp int             `json:"timestamp"`
	Payload   json.RawMessage `json:"payload"`
}

type Message struct {
	ID        string  `json:"id"`
	UserID    string  `json:"user_id"`
	Message   string  `json:"message"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Sensor struct {
	ID        string `json:"id"`
	DeviceID  string `json:"device_id`
	Result    string `json:"result"`
	ProductID string `json:"product_id"`
}

func main() {
	f, err := os.Open("message_a.json")
	// f, err := os.Open("message_b.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var r Response
	if err := json.NewDecoder(f).Decode(&r); err != nil {
		log.Fatal(err)
	}

	switch r.Type {
	case "message":
		var m Message
		_ = json.Unmarshal(r.Payload, &m)
		fmt.Printf("Message: %+v\n", m)
	case "sensor":
		var s Sensor
		_ = json.Unmarshal(r.Payload, &s)
		fmt.Printf("Sensor: %+v\n", s)
	}
}
