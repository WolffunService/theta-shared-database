package main

import "github.com/WolffunGame/theta-shared-database/database/mongodb"

func main()  {
	_ ,_ ,_ = mongodb.SetDefaultConfig(nil,nil)
}

