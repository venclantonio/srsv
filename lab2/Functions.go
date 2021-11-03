package lab2

import (
	"time"
)

func CreateCarSemaphore() map[string]*Semafor {
	semaphores := make(map[string]*Semafor)
	// Horizontal semaphores
	semaphores[StraightNorth] = New()
	semaphores[LeftWest] = New()

	// Vertical semaphores
	semaphores[StraightEast] = New()
	semaphores[LeftNorth] = New()

	// Special right semaphores
	semaphores[RightEast] = New()
	semaphores[RightNorth] = New()

	return semaphores
}

func CreateWalkerSemaphore() map[string]*Semafor {
	semaphores := make(map[string]*Semafor)
	semaphores[WalkerBottom] = New()
	semaphores[WalkerLeft] = New()
	semaphores[WalkerRight] = New()
	semaphores[WalkerTop] = New()
	return semaphores
}

func StartSemaphores(semaphores map[string]*Semafor) {
	go semaphores[StraightNorth].Start()
	go semaphores[LeftWest].Start()
	go semaphores[StraightEast].Start()
	go semaphores[LeftNorth].Start()
	go semaphores[RightEast].Start()
	go semaphores[RightNorth].Start()
	go semaphores[WalkerBottom].Start()
	go semaphores[WalkerLeft].Start()
	go semaphores[WalkerRight].Start()
	go semaphores[WalkerTop].Start()
}

func StartGenerator(generator *Generator, char rune) {
	go generator.Start(char)
}

func Sleep(semaphore *Semafor, walkerSemaphore *Semafor) {
	// counter := 0
	for i := WaitTime; i > 0; i-- {
		if semaphore.currentVehicle == 'A' {
			i--
		}
		if walkerSemaphore != nil {
			if walkerSemaphore.currentVehicle == 'P' {
				i--
			}
		}
		time.Sleep(1 * time.Second)
		// counter++
	}
	// fmt.Printf("spavo sam: %d sekundi\n", counter)
}
