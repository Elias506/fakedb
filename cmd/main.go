package main

import (
	"fmt"
	"Fakedb"
)

func main() {
	db := &Fakedb.DB{}
	err := db.CreateTable(`33|id = int, amount = string`)
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Tables["33"].Write()
	err = db.Insert(`33|id=1,amount="tyu"`)
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Tables["33"].Write()
	err = db.Insert(`33|id=2,amount="tyu2"`)
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Tables["33"].Write()
	s, err := db.Select(`33|`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("___")
	for _, v := range s {
		fmt.Println(v)
	}
}
