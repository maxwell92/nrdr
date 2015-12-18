package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)



func main() {
	fmt.Println("This is a redis")
	c, err := redis.Dial("tcp", "172.21.1.28:6379")
	if err != nil {
		fmt.Println("connected to redis server failed: ", err)
	}




}
