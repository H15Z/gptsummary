package actors

import "log"

type ActorInterface interface {
	IsActive() bool
	IsFinished() bool
	MessageCount() int
	AddMsg(msg ActorMsg)
	GetId() string
}

type ActorMsg struct {
	sender *Actor
	Msg    interface{}
}

type Actor struct {
	ID             string
	msgs           chan ActorMsg
	processing_msg []bool
	Finished       bool
	Supervisor     *Supervisor
}

func (a *Actor) Init(super *Supervisor, recieve_callback func(ActorMsg), workers int) {
	a.Supervisor = super
	a.msgs = make(chan ActorMsg, 1000000) // TODO make this a parameter
	a.Finished = false

	for i := 0; i < workers; i++ {
		a.processing_msg = append(a.processing_msg, false)
		go a.actorLoop(recieve_callback, i)
	}

	log.Println("[ACTOR INITIALIZED]:", a.ID, "[WORKERS]:", workers)

}

func (a *Actor) actorLoop(recieve_callback func(ActorMsg), worker_id int) {

	for m := range a.msgs {
		a.processing_msg[worker_id] = true
		recieve_callback(m)
		a.processing_msg[worker_id] = false
	}
}

func (a *Actor) ActorStop() {
	a.Finished = true
	close(a.msgs)
	log.Println("[ACTOR STOPED]:", a.ID)

}

func (a *Actor) MessageCount() int {
	return len(a.msgs)
}

func (a *Actor) IsActive() bool {
	for _, p := range a.processing_msg {
		if p {
			return true
		}
	}
	return false
}

func (a *Actor) IsFinished() bool {
	return a.Finished
}

// send message and adds sender to message
func (a *Actor) SendMsg(address ActorInterface, msg ActorMsg) {
	msg.sender = a
	address.AddMsg(msg)
}

// recive message add to queue
func (a *Actor) AddMsg(msg ActorMsg) {
	a.msgs <- msg
}

func (a *Actor) GetId() string {
	return a.ID
}
