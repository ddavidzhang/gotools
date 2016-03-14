package main

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	argsLen := len(os.Args)
	if argsLen != 3 {
		fmt.Println("command dbhost enterprisename")
	} else {

		fmt.Printf("host: %v ea: %v \n", os.Args[1], os.Args[2])
		session, err := mgo.Dial(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer session.Close()
		dbs, err := session.DatabaseNames()
		if err != nil {
		}
		for _, dbName := range dbs {
			// fmt.Println(dbName)
			c := session.DB(dbName).C("C_MetaSession")
			n, err := c.Find(bson.M{"EA": os.Args[2]}).Count()
			if err == nil && n != 0 {
				fmt.Println(dbName)
			}
		}
	}
}
