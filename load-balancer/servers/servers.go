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
	for x := 0; x < amount;x++{
		s.Ports = append(s.Ports, x)
	}
}

func (s*ServerList) Pop() int{
	port := s.Ports[0]
	s.Ports = s.Ports[1:]
	return port
}
