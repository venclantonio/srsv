package lab3

import (
	"fmt"
	"time"
)

type Controller struct {
	draw *Drawer
}

func NewController() {
	floors := createFloors(numberOfFloors)
	lift := NewLift()
	generator := newGen(floors)
	generator.startGenerator()
	fmt.Println(int('a'))
	// lift.currentFloor = 2
	// floors[0].upButton.pressed = true
	// floors[1].downButton.pressed = true
	// floors[1].upButton.pressed = true
	// floors[2].downButton.pressed = true

	/*
		floor1 := []*Person{
			&Person{
				name:             "A",
				sourceFloor:      1,
				destinationFloor: 3,
				direction:        upDirection,
			},
		}

		floor2 := []*Person{
			&Person{
				name:             "B",
				sourceFloor:      2,
				destinationFloor: 1,
				direction:        downDirection,
			},
			&Person{
				name:             "D",
				sourceFloor:      2,
				destinationFloor: 3,
				direction:        upDirection,
			},
		}
		floor3 := []*Person{
			{
				name:             "C",
				sourceFloor:      3,
				destinationFloor: 1,
				direction:        downDirection,
			},
		}
		floors[0].people = floor1
		floors[1].people = floor2
		floors[2].people = floor3
	*/

	drawer := NewDraw()

	go drawer.Start()

	go func() {
		for {

			drawer.toDraw <- &Input{
				floors: floors,
				lift:   *lift,
			}
			time.Sleep(1000 * time.Millisecond)

			found, floorNumber, direction := lift.findNewPerson(floors)
			if found {
				if floorNumber == lift.currentFloor {
					lift.doThing(floors, *drawer)
				}
				if floorNumber > lift.currentFloor {
					for i := lift.currentFloor; i < floorNumber; i++ {
						lift.direction = direction
						lift.currentFloor += 1
						lift.doThing(floors, *drawer)
					}
				}
				if floorNumber < lift.currentFloor {
					for i := lift.currentFloor; i > floorNumber; i-- {
						lift.direction = direction
						lift.currentFloor -= 1
						lift.doThing(floors, *drawer)
					}
				}
			}
			for lift.numberOfPeople > 0 || lift.doors == openDoors {
				lift.doThing(floors, *drawer)
			}
			fmt.Println("Gotov s rundom")
		}
	}()
}
