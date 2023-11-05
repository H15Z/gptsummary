package actors

import (
	"log"
	"time"
)

type Supervisor struct {
	Actors []ActorInterface
}

func NewSupervisor() *Supervisor {
	log.Println("[SUPERVISOR INITIALIZED]")
	return &Supervisor{}
}

func (s *Supervisor) Watch(a ActorInterface) {
	log.Println("[SUPERVISOR WATCHING]:", a.GetId())
	s.Actors = append(s.Actors, a)
}

func (s *Supervisor) Run() {
	for !s.AllFinished() {
		time.Sleep(100 * time.Millisecond)
	}

	log.Println("[SUPERVISOR ALL WORKERS FINISHED]")
}

func (s *Supervisor) AllFinished() bool {
	for _, a := range s.Actors {
		if !a.IsFinished() {
			return false
		}

		if a.IsActive() {
			return false
		}
	}

	return true
}
