package actors

func StartStream(count int, threads int) {
	super := NewSupervisor()

	a := NewLoaderActor(super)
	a.SendMsg(a, ActorMsg{
		Msg: LoaderMsg{
			Count:   count,
			Threads: threads,
		},
	})

	super.Watch(a)
	super.Run() //Keep running untill all actors finish
}
