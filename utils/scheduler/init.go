package scheduler

func Initiator() {
	scheduler := NewScheduler()

	scheduler.AddJob("@every 1s", LogAllCar())

	//scheduler.Start()
}

func LogAllCar() func() {
	return func() {
		// create your logic
	}
}
