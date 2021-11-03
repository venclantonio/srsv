package lab2

type Semafor struct {
	currentVehicle rune
	currentState   int
	stateChan      chan int
	vehicleChan    chan rune
}

func New() *Semafor {
	return &Semafor{
		currentState: Red,
		stateChan:    make(chan int),
		vehicleChan:  make(chan rune),
	}
}

func (s *Semafor) Start() {
	for {
		select {
		case nextVehicle := <-s.vehicleChan:
			s.currentVehicle = nextVehicle
		case nextState := <-s.stateChan:
			s.currentState = nextState
			if nextState == Green {
				s.currentVehicle = rune(0)
			}
		}
	}
}
