package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func GetConn() redis.Conn {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", "127.0.0.1", 6379))

	if err != nil {
		log.Println("connect redis error", err)
		return nil
	}

	return conn
}

func main() {
	conn := GetConn()
	if conn == nil {
		log.Println("conn is nil")
	}
	defer conn.Close()

	conn.Send("SET", "fruit", "apple")
	conn.Send("SET", "animal", "dog")
	conn.Send("GET", "fruit")
	conn.Flush()
	for i := 0; i < 3; i++ {
		result, err := conn.Receive()
		if err != nil {
			log.Println("pipeline error")
		}
		fmt.Println(redis.String(result, err))
	}
}
