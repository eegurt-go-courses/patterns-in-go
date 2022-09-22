package main

type Exercise interface {
	getBurntCalories() int
}

type WarmUp struct {
	workoutSession IWorkoutSession
}

func (e *WarmUp) getBurntCalories() int {
	burntCalories := e.workoutSession.getBurntCalories()
	return burntCalories + 2
}

func (e *WarmUp) getSessionType() string {
	return e.workoutSession.getSessionType() + "\nYou have warmed up"
}

type Squat struct {
	workoutSession IWorkoutSession
}

func (e *Squat) getBurntCalories() int {
	burntCalories := e.workoutSession.getBurntCalories()
	return burntCalories + 4
}

func (e *Squat) getSessionType() string {
	return e.workoutSession.getSessionType() + "\nYou have done squats"
}

type Pullup struct {
	workoutSession IWorkoutSession
}

func (e *Pullup) getBurntCalories() int {
	burntCalories := e.workoutSession.getBurntCalories()
	return burntCalories + 8
}

func (e *Pullup) getSessionType() string {
	return e.workoutSession.getSessionType() + "\nYou have done pull ups"
}

type Deadlift struct {
	workoutSession IWorkoutSession
}

func (e *Deadlift) getBurntCalories() int {
	burntCalories := e.workoutSession.getBurntCalories()
	return burntCalories + 12
}

func (e *Deadlift) getSessionType() string {
	return e.workoutSession.getSessionType() + "\nYou have done deadlifts"
}
