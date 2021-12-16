package lab3

import (
	"fmt"
	"time"
)

type Lift struct {
	people         []*Person
	direction      rune
	doors          rune
	numberOfPeople int
	currentFloor   int
	peopleLeft     []*Person
}

func NewLift() *Lift {
	return &Lift{
		people:         make([]*Person, 0, liftCapacity),
		direction:      upDirection,
		doors:          openDoors,
		numberOfPeople: 0,
		currentFloor:   1,
	}
}

func (lift *Lift) AddPeopleToLift(people []*Person) {
	peopleInLift := lift.people[0:max(lift.numberOfPeople, 0)]
	peopleToAdd := make([]*Person, len(people))
	for i, person := range people {
		peopleToAdd[i] = person
	}
	lift.people = append(peopleInLift, peopleToAdd...)
	lift.numberOfPeople += len(people)
}

func (lift *Lift) MoveLift(floors []*Floor) {

	for lift.numberOfPeople > 0 {
		floorsToStop := lift.getFloorsToStop()
		for i, b := range floorsToStop {
			fmt.Printf("trenutno sam na %d katu\n", i+1)
			peopleToLeaveFloor := lift.boardPeople(floors[i])
			lift.AddPeopleToLift(peopleToLeaveFloor)
			floors[i].removePeopleFromFloor(peopleToLeaveFloor)
			NewDraw()
			if b {
				fmt.Printf("Stao sam na %d. katu jer je netko morao izaci\n", i+1)
				lift.PeopleLeave(i + 1)
			}
			lift.currentFloor = i + 1

		}
	}
}

func (lift *Lift) oneStep(floors []*Floor) {
	// open lift doors
	// people leave
	lift.removePeople(floors[lift.currentFloor-1])
	// people come in
	peopleToBoard := lift.boardPeople(floors[lift.currentFloor-1])
	if len(peopleToBoard) == 0 && lift.numberOfPeople == 0 {
		lift.switchDirection()
	}
	peopleToBoard = lift.boardPeople(floors[lift.currentFloor-1])
	lift.AddPeopleToLift(peopleToBoard)
	floors[lift.currentFloor-1].removePeopleFromFloor(peopleToBoard)

	// draw
	// NewDraw(floors,*lift)
	// lift moves in direction
	// lift.moveInDirection()
}

func (lift *Lift) getFloorsToStop() []bool {
	result := make([]bool, numberOfFloors)
	for _, person := range lift.people {
		result[person.destinationFloor-1] = true
	}
	return result
}

func (lift *Lift) PeopleLeave(floorNumber int) {
	result := make([]*Person, liftCapacity)
	counter := 0
	for _, person := range lift.people {
		if !(person.destinationFloor == floorNumber) {
			result[counter] = person
			counter++
		}
	}
	lift.people = result[0:counter]
	lift.numberOfPeople = counter
}

func (lift *Lift) boardPeople(floor *Floor) []*Person {
	result := make([]*Person, liftCapacity)
	counter := 0
	for _, person := range floor.people {
		if person.direction == lift.direction {
			result[counter] = person
			counter++
			if counter == liftCapacity {
				break
			}
		}
	}
	return result[0:counter]
}

func (lift *Lift) removePeople(floor *Floor) {
	peopleInLift := make([]*Person, liftCapacity)
	peopleCounter := 0
	peopleRemoved := make([]*Person, liftCapacity)
	peopleRemovedCounter := 0
	for _, person := range lift.people {
		if person.destinationFloor == floor.level+1 {
			peopleRemoved[peopleRemovedCounter] = person
			peopleRemovedCounter++
		} else {
			if person.name != "" {
				peopleInLift[peopleCounter] = person
				peopleCounter++
			}
		}
	}
	lift.numberOfPeople = peopleCounter
	lift.people = peopleInLift[0:peopleCounter]
	lift.peopleLeft = peopleRemoved[0:peopleRemovedCounter]
}

func (lift *Lift) moveInDirection() {
	if lift.numberOfPeople == 0 {
		return
	}
	lift.peopleLeft = make([]*Person, 0, 6)
	liftDirection := string(lift.direction)
	for _, person := range lift.people {
		if string(person.direction) == liftDirection {
			if lift.direction == upDirection {
				lift.currentFloor += 1
			} else {
				lift.currentFloor -= 1
			}
			return
		}
	}
	lift.switchDirection()
}

func (lift *Lift) switchDirection() {
	if lift.direction == upDirection {
		lift.direction = downDirection
		// lift.currentFloor -= 1
	} else {
		lift.direction = upDirection
		// lift.currentFloor +=1
	}
}

func (lift *Lift) openDoors() {
	lift.doors = openDoors
}

func (lift *Lift) closeDoors() {
	lift.doors = closedDoors
	lift.peopleLeft = make([]*Person, 0, 6)
}

func (lift *Lift) findNewPerson(floors []*Floor) (bool, int, rune) {
	for _, floor := range floors {
		if floor.downButton.pressed {
			floor.downButton.pressed = false
			return true, floor.level + 1, downDirection
		}
		if floor.upButton.pressed {
			floor.upButton.pressed = false
			return true, floor.level + 1, upDirection
		}

	}
	return false, 0, rune(0)
}

func (lift *Lift) needToStop(floor *Floor) bool {
	if lift.direction == upDirection && floor.upButton.pressed {
		return true
	}
	if lift.direction == downDirection && floor.downButton.pressed {
		return true
	}
	for _, person := range lift.people {
		if person.destinationFloor == floor.level+1 {
			return true
		}
	}
	for _, person := range floor.people {
		if person.sourceFloor == floor.level+1 && person.direction == lift.direction {
			return true
		}
	}
	return false
}
func DefaultPeople() []*Person {
	result := make([]*Person, liftCapacity)
	for i := 0; i < liftCapacity; i++ {
		result[i] = &Person{name: ""}
	}
	return result
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func (lift *Lift) doThing(floors []*Floor, drawer Drawer) {
	if lift.needToStop(floors[lift.currentFloor-1]) {
		time.Sleep(1000 * time.Millisecond)
		lift.openDoors()
		drawer.toDraw <- &Input{
			floors: floors,
			lift:   *lift,
		}
		time.Sleep(1000 * time.Millisecond)
		lift.oneStep(floors)
		drawer.toDraw <- &Input{
			floors: floors,
			lift:   *lift,
		}
		time.Sleep(1000 * time.Millisecond)
		lift.closeDoors()
		drawer.toDraw <- &Input{
			floors: floors,
			lift:   *lift,
		}
	}

	time.Sleep(1000 * time.Millisecond)
	lift.moveInDirection()
	drawer.toDraw <- &Input{
		floors: floors,
		lift:   *lift,
	}

}
