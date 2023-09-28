package main

import (
	"log"
)

func main() {
	
	explorer := NewVademIndex("https://www.vademecum.com/Diego-Andres-Marmota")
	explorer = NewLogin(explorer)

	api := NewApiVademecum(explorer)
	log.Fatal(api.Iniciate(":8000"))

}

