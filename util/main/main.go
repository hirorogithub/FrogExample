package main

import (
	util "../"
	"log"
)

type testS struct {
	ID      int
	Context string
}

func main() {

	dst := testS{}

	src := map[string]interface{}{
		"ID": 1, "Conte1xt": "2",
	}

	log.Println(util.MapToStruct(src, &dst))
	log.Println(dst)

}
