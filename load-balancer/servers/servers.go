package servers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type ServerList struct{
	Ports []int
}

func (s*ServerList) Populate(amount int){
	if amount >= 10{
		log.Fatal("Port Amount cant exceed 10")
	}
}
