package main

import (
	"log"

	"github.com/adam69iiu/backend/internal/server"
)



func main(){
 if err :=  server.Run(); err != nil{
	 log.Fatalf(" running server failed: %v",err)
 }
}
