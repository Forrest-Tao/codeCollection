package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type Person struct {
	Name    string  `json:"name"`          // 指定JSON中的字段名
	Age     int     `json:"age,omitempty"` // 如果 Age 是 0，则忽略该字段
	Email   string  `json:"-"`             // 忽略字段 Email
	Address Address `json:"address"`       // 嵌套结构体 Address
}

func main() {
	// 创建一个 Person 实例
	p := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com", // 这个字段不会被序列化
		Address: Address{
			City:    "New York",
			Country: "USA",
		},
	}

	// 序列化为 JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error serializing to JSON:", err)
		return
	}
	fmt.Println("Serialized JSON:")
	fmt.Println(string(jsonData)) // 输出序列化后的 JSON 字符串

	// JSON 字符串（用于反序列化）
	jsonString := `{"name":"Jane Doe","age":25,"address":{"city":"Los Angeles","country":"USA"}}`

	// 反序列化为 Go 结构体
	var newPerson Person
	err = json.Unmarshal([]byte(jsonString), &newPerson)
	if err != nil {
		fmt.Println("Error deserializing JSON:", err)
		return
	}

	fmt.Println("\nDeserialized struct:")
	fmt.Printf("%+v\n", newPerson) // 输出反序列化后的结构体
}
