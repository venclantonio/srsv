package lab3

type Button struct {
	direction rune
	pressed   bool
}

type Floor struct {
	level      int
	people     []*Person
	upButton   *Button
	downButton *Button
}

func newFloor(level int, upButton *Button, downButton *Button) *Floor {
	return &Floor{
		level:      level,
		upButton:   upButton,
		downButton: downButton,
		people:     make([]*Person, 0, 10),
	}
}

func newButton(direction rune) *Button {
	return &Button{
		direction: direction,
		pressed:   false,
	}
}

func createFloors(numberOfFloors int) []*Floor {
	result := make([]*Floor, numberOfFloors)
	// result[0] = newFloor(0, buttons[0], nil)
	for i := 0; i < numberOfFloors; i++ {
		result[i] = newFloor(i, newButton(upDirection), newButton(downDirection))
	}
	// result[numberOfFloors-1] = newFloor(numberOfFloors-1, nil, buttons[1])
	return result
}

func (floor *Floor) FindPeopleToEnterLift(lift *Lift) []*Person {
	result := make([]*Person, liftCapacity)
	peopleCounter := 0

	for _, person := range floor.people {
		if lift.numberOfPeople+peopleCounter >= liftCapacity {
			break
		}
		if person.direction == lift.direction {
			result[peopleCounter] = person
			peopleCounter++
		}
	}
	floor.removePeopleFromFloor(result[0:peopleCounter])
	return result[0:peopleCounter]
}

func (floor *Floor) removePeopleFromFloor(people []*Person) {
	if len(people) == 0 {
		return
	}
	result := DefaultPeople()
	// peopleOnFloor := floor.people
	peopleCounter := 0
	for _, person := range floor.people {
		if !contains(people, person) {
			result[peopleCounter] = person
			peopleCounter++
		}
	}

	floor.disableButtons(result[0:peopleCounter])
	floor.people = result[0:peopleCounter]
}

func (floor *Floor) disableButtons(people []*Person) {
	if len(people) == 0 {
		floor.downButton.pressed = false
		floor.upButton.pressed = false
	}
	downPeople := 0
	upPeople := 0
	for _, person := range people {
		if person.direction == upDirection {
			upPeople++
		}
		if person.direction == downDirection {
			downPeople++
		}
	}
	if downPeople == 0 {
		floor.downButton.pressed = false
	}
	if upPeople == 0 {
		floor.upButton.pressed = false
	}
}

func contains(floor []*Person, person *Person) bool {
	for _, p := range floor {
		if p.name == person.name {
			return true
		}
	}
	return false
}
