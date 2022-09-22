package main

import (
	"fmt"
)

func main() {
	kach := &OutdoorSession{}

	kachWarmup := &WarmUp{workoutSession: kach}

	kachWarmupPullup := &Pullup{workoutSession: kachWarmup}
	kachWarmup2Pullups := &Pullup{workoutSession: kachWarmupPullup}

	fmt.Printf("You are at %s\nBurnt: %v calories.\n",
		kachWarmup2Pullups.getSessionType(), kachWarmup2Pullups.getBurntCalories())
}
