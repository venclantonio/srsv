package lab2

type Controller struct {
	semaphores map[string]*Semafor
	draw       *Drawer
}

func NewController() {
	walkerSemaphores := CreateWalkerSemaphore()
	carSemaphores := CreateCarSemaphore()

	// create generators
	walkerGenerator := newGen(walkerSemaphores)
	carGenerator := newGen(carSemaphores)

	semaphores := walkerSemaphores
	for k, v := range carSemaphores {
		semaphores[k] = v
	}

	StartSemaphores(semaphores)

	StartGenerator(walkerGenerator, 'P')
	StartGenerator(carGenerator, 'A')

	drawer := NewDraw()
	go drawer.Start()

	counter := 0
	go func() {

		for {

			SpecialRightEast(semaphores)
			drawer.Inputs <- &input{
				counter:    counter,
				semaphores: semaphores,
			}
			counter++
			Sleep(semaphores[StraightEast], semaphores[WalkerLeft])

			NorthSouth(semaphores)
			drawer.Inputs <- &input{
				counter:    counter,
				semaphores: semaphores,
			}
			counter++
			Sleep(semaphores[LeftWest], nil)

			NorthSouthLeft(semaphores)
			drawer.Inputs <- &input{
				counter:    counter,
				semaphores: semaphores,
			}
			counter++
			Sleep(semaphores[RightNorth], nil)

			SpecialRightSouth(semaphores)
			drawer.Inputs <- &input{
				counter:    counter,
				semaphores: semaphores,
			}
			counter++
			Sleep(semaphores[StraightEast], semaphores[WalkerBottom])

			WestEast(semaphores)
			drawer.Inputs <- &input{
				counter:    counter,
				semaphores: semaphores,
			}
			counter++
			Sleep(semaphores[LeftWest], nil)

			WestEastLeft(semaphores)
			drawer.Inputs <- &input{
				counter:    counter,
				semaphores: semaphores,
			}
			counter++
			Sleep(semaphores[RightEast], nil)
		}

	}()
}

func SpecialRightEast(semaphores map[string]*Semafor) {
	semaphores[RightEast].stateChan <- Green
	setRedLights(semaphores, RightEast)
}

func SpecialRightSouth(semaphores map[string]*Semafor) {
	semaphores[RightNorth].stateChan <- Green
	setRedLights(semaphores, RightNorth)
}

func WestEastLeft(semaphores map[string]*Semafor) {
	semaphores[LeftNorth].stateChan <- Green
	setRedLights(semaphores, "LeftNorth")
}

func setRedLights(semaphores map[string]*Semafor, s string) {
	for k, v := range semaphores {
		if k != s {
			v.stateChan <- Red
		}
	}

}

func WestEast(semaphores map[string]*Semafor) {
	semaphores[StraightEast].stateChan <- Green
	setRedLights(semaphores, "StraightEast")
	semaphores[WalkerTop].stateChan <- Green
	semaphores[WalkerBottom].stateChan <- Green
}

func NorthSouthLeft(semaphores map[string]*Semafor) {
	semaphores[LeftWest].stateChan <- Green
	setRedLights(semaphores, "LeftWest")
}
func NorthSouth(semaphores map[string]*Semafor) {
	semaphores[StraightNorth].stateChan <- Green
	setRedLights(semaphores, "StraightNorth")
	semaphores[WalkerLeft].stateChan <- Green
	semaphores[WalkerRight].stateChan <- Green

}
