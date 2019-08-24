package service

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var red redis.Conn
func Ini()  {
	red ,_ = redis.Dial("tcp","127.0.0.1:6379")
}

func Set(key string,value string)  {
	_,err:=red.Do("SET",key,value,"EX",60*60*24)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Get(key string) interface{} {
	value,err:=red.Do("GET",key)
	if err != nil {
		fmt.Println(err.Error())
	}
	return value
}

func SetOne(key string,value string)  {
	_,err:=red.Do("SET",key,value)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetOne(key string) interface{}{
	value,err := red.Do("GET",key)
	if err != nil {
		fmt.Println(err.Error())
	}
	if value!=nil{
		red.Do("DEL",key)
	}
	return value
}