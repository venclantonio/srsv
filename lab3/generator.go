package lab3

import (
	"math/rand"
	"time"
)

type Generator struct {
	floors []*Floor
}

func newGen(floors []*Floor) *Generator {
	return &Generator{
		floors: floors,
	}
}

func (g *Generator) start() {
	breakPoint := 0
	for {
		randomPerson := genRandomPerson()
		g.addPersonToFloor(randomPerson)
		time.Sleep(highDensity * time.Second)
		breakPoint++
		if breakPoint == 20 {
			break
		}

	}
}

func (g *Generator) addPersonToFloor(person *Person) {
	people := g.floors[person.sourceFloor-1].people
	result := make([]*Person, len(people)+1)
	for i := 0; i < len(people); i++ {
		result[i] = people[i]
	}
	result[len(people)] = person
	g.floors[person.sourceFloor-1].people = result
	if person.direction == upDirection {
		g.floors[person.sourceFloor-1].upButton.pressed = true
	}
	if person.direction == downDirection {
		g.floors[person.sourceFloor-1].downButton.pressed = true
	}
}

func (g *Generator) startGenerator() {
	go g.start()
}

func genRandomPerson() *Person {
	var randomChar rune
	randomInt := rand.Intn(2)
	if randomInt == 1 {
		randomChar = 'A' + rune(rand.Intn(26))
	} else {
		randomChar = 'a' + rune(rand.Intn(26))
	}
	randomSource := 1 + rand.Intn(numberOfFloors)
	randomDestination := 1 + rand.Intn(numberOfFloors)
	for randomDestination == randomSource {
		randomDestination = 1 + rand.Intn(numberOfFloors)
	}
	randomChar += 1
	var direction rune
	if randomSource > randomDestination {
		direction = downDirection
	} else {
		direction = upDirection
	}

	return &Person{
		name:             string(randomChar),
		sourceFloor:      randomSource,
		destinationFloor: randomDestination,
		direction:        direction,
	}
}
