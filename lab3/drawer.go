package lab3

import "fmt"

type Input struct {
	floors []*Floor
	lift   Lift
}

type Drawer struct {
	toDraw chan *Input
}

type Tuple struct {
	source      int
	destination int
}

var alivePeople map[string]Tuple

func NewDraw() *Drawer {
	return &Drawer{toDraw: make(chan *Input)}
}

func drawLiftDirection(lift Lift) string {
	result := ""
	result += string(lift.direction) + " " + string(lift.doors)
	return result
}

func drawFloors(floors []*Floor, lift Lift) {
	for i := 0; i <= numberOfFloors-1; i++ {
		fmt.Printf("%d:"+drawPeopleOnFloor(floors[numberOfFloors-i-1])+"|"+drawLift(lift, numberOfFloors-i)+"|  "+drawLeavePeople(lift, numberOfFloors-i)+"\n", numberOfFloors-i)
	}

}

func drawLeavePeople(lift Lift, i int) string {
	result := ""
	for _, person := range lift.peopleLeft {
		if person.destinationFloor == i {
			result += person.name
		}
	}
	return result
}

func drawLift(lift Lift, currentFloor int) string {
	if len(lift.people) == 0 {
		fmt.Print()
	}
	if currentFloor != lift.currentFloor {
		return "\t\t"
	}
	result := ""
	for i := 0; i < (10-lift.numberOfPeople+1)/2; i++ {
		result += " "
	}
	result += "["

	for _, person := range lift.people {
		if person.destinationFloor == currentFloor && lift.doors == openDoors {
			// continue
		}
		result += person.name
	}

	result += "]"
	for i := 0; i < (10-lift.numberOfPeople)/2; i++ {
		result += " "
	}
	return result
}

func drawPeopleOnFloor(floor *Floor) string {
	result := ""
	people := floor.people
	for _, person := range people {
		result += person.name
	}
	for i := len(result); i < 9; i++ {
		result += " "
	}
	return result
}

func (drawer *Drawer) Start() {
	for {
		select {
		case state := <-drawer.toDraw:
			drawer.draw(state.floors, state.lift)
		}
	}
}

func (drawer *Drawer) draw(floors []*Floor, lift Lift) {
	fmt.Println("\t      lift1")
	fmt.Println("Smjer/vrata     " + drawLiftDirection(lift))
	fmt.Println("==========================Izasli")
	drawFloors(floors, lift)
	fmt.Println("==========================")
	drawPeople(floors, lift)
}

func drawPeople(floors []*Floor, lift Lift) string {
	alivePeople = make(map[string]Tuple)
	result := ""
	for _, floor := range floors {
		for _, person := range floor.people {
			result += person.name
			alivePeople[person.name] = Tuple{
				source:      person.sourceFloor,
				destination: person.destinationFloor,
			}
		}
	}

	for _, person := range lift.people {
		result += person.name
		alivePeople[person.name] = Tuple{
			source:      person.sourceFloor,
			destination: person.destinationFloor,
		}
	}
	fmt.Println("Putnici:  " + result)
	fmt.Println("     od:  " + drawSourceFloor(alivePeople))
	fmt.Println("     do:  " + drawDestinationFloor(alivePeople))
	return result
}

func drawDestinationFloor(people map[string]Tuple) string {
	result := ""
	for _, tuple := range people {
		result += fmt.Sprintf("%d", tuple.destination)
	}
	return result
}

func drawSourceFloor(people map[string]Tuple) string {
	result := ""
	for _, tuple := range people {
		result += fmt.Sprintf("%d", tuple.source)
	}
	return result
}
