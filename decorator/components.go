package main

type IWorkoutSession interface {
	getBurntCalories() int
	getSessionType() string
}

type HomeSession struct {
}

func (*HomeSession) getBurntCalories() int {
	return 0
}

func (*HomeSession) getSessionType() string {
	return "home"
}

type OutdoorSession struct {
}

func (*OutdoorSession) getBurntCalories() int {
	return 5
}

func (*OutdoorSession) getSessionType() string {
	return "outdoors"
}

type GymSession struct {
}

func (*GymSession) getBurntCalories() int {
	return 25
}

func (*GymSession) getSessionType() string {
	return "gym"
}
