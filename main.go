package main

import "wolffundb/database/mongodb"

func main()  {
	_ ,_ ,_ = mongodb.SetDefaultConfig(nil,nil)
}

