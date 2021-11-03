package lab2

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mitchellh/colorstring"
)

type input struct {
	counter    int
	semaphores map[string]*Semafor
}

type Drawer struct {
	Inputs chan *input
}

func NewDraw() *Drawer {
	return &Drawer{Inputs: make(chan *input)}
}

func (drawer *Drawer) Start() {
	for {
		select {
		case crossroad := <-drawer.Inputs:
			drawer.draw(crossroad.semaphores)
		}
	}
}

func (drawer *Drawer) draw(semaphores map[string]*Semafor) {

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println("\t\t|\t|\t|\t|\t|\t|")
	fmt.Println("\t\t|\t|\t|\t|\t|\t|")
	fmt.Println("\t\t|   " + drawCar(semaphores[RightEast]) + "   |   " + drawCar(semaphores[StraightNorth]) + "   |   " +
		drawCar(semaphores[LeftWest]) + "   |\t|\t|")
	fmt.Println("\t     " + drawWalker(semaphores[WalkerTop]) + " " + drawHorizontalSemaphore(semaphores[WalkerTop]) + "|   ←   |   ↓   |   →   |\t|\t|" + drawHorizontalSemaphore(semaphores[WalkerTop]) + " " + drawWalker(semaphores[WalkerTop]))
	fmt.Println("\t  " + drawWalker(semaphores[WalkerLeft]) + " " + drawHorizontalSemaphore(semaphores[WalkerLeft]) + "   |" + drawVerticalSemaphore(semaphores[RightEast]) + "|" + drawVerticalSemaphore(semaphores[StraightNorth]) +
		"|" + drawVerticalSemaphore(semaphores[LeftWest]) + "|\t|\t|   " + drawHorizontalSemaphore(semaphores[WalkerRight]) + " " + drawWalker(semaphores[WalkerRight]))
	fmt.Println("----------------+\t\t\t\t\t+----------------")
	fmt.Println("\t\t\t\t\t\t\t" + drawHorizontalSemaphore(semaphores[RightNorth]) + " ↑ " + drawCar(semaphores[RightNorth]))
	fmt.Println("-----------------\t\t\t\t\t-----------------")
	fmt.Println("\t\t\t\t\t\t\t" + drawHorizontalSemaphore(semaphores[StraightEast]) + " ← " + drawCar(semaphores[StraightEast]))
	fmt.Println("-----------------\t\t\t\t\t-----------------")
	fmt.Println("\t     " + drawCar(semaphores[LeftNorth]) + " ↑ " + drawHorizontalSemaphore(semaphores[LeftNorth]) + "\t\t\t\t\t" +
		drawHorizontalSemaphore(semaphores[LeftNorth]) + " ↓ " + drawCar(semaphores[LeftNorth]))
	fmt.Println("-----------------\t\t\t\t\t-----------------")
	fmt.Println("\t     " + drawCar(semaphores[StraightEast]) + " → " + drawHorizontalSemaphore(semaphores[StraightEast]))
	fmt.Println("-----------------\t\t\t\t\t-----------------")
	fmt.Println("\t     " + drawCar(semaphores[RightNorth]) + " ↓ " + drawHorizontalSemaphore(semaphores[RightNorth]))
	fmt.Println("----------------+\t\t\t\t\t+----------------")
	fmt.Println("\t  " + drawWalker(semaphores[WalkerLeft]) + " " + drawHorizontalSemaphore(semaphores[WalkerLeft]) + "   |\t|\t|" + drawVerticalSemaphore(semaphores[LeftWest]) + "|" + drawVerticalSemaphore(semaphores[StraightNorth]) +
		"|" + drawVerticalSemaphore(semaphores[RightEast]) + "|   " + drawHorizontalSemaphore(semaphores[WalkerRight]) + " " + drawWalker(semaphores[WalkerRight]))
	fmt.Println("\t     " + drawWalker(semaphores[WalkerBottom]) + " " + drawHorizontalSemaphore(semaphores[WalkerBottom]) + "|\t|\t|   ←   |   ↑   |   →   |" + drawHorizontalSemaphore(semaphores[WalkerBottom]) + " " + drawWalker(semaphores[WalkerBottom]))
	fmt.Println("\t\t|\t|\t|   " + drawCar(semaphores[LeftWest]) + "   |   " + drawCar(semaphores[StraightNorth]) + "   |   " +
		drawCar(semaphores[RightEast]) + "   |")
	fmt.Println("\t\t|\t|\t|\t|\t|\t|")
	fmt.Println("\t\t|\t|\t|\t|\t|\t|")

}

func drawWalker(semaphore *Semafor) string {
	if semaphore.currentVehicle == 'P' {
		return colorstring.Color("[blue]P")
	}
	return " "
}

func drawCar(semaphore *Semafor) string {
	if semaphore.currentVehicle == 'A' {
		return colorstring.Color("[blue]A")
	}
	return " "
}

/*func drawCarsLeft(semafor *Semafor) string {
	if semafor.currentState == Green {
		return "\t\t  <- A A A A A A A A A A A A A A"
	}
	return "\t\t\t\t\t"
}*/

/*func drawCars(semafor *Semafor) string {
	if semafor.currentState == Green {
		return " A A A A A A A A A A A A A A A A A A A A A ->"
	}
	return ""
}*/

func drawVerticalSemaphore(semaphore *Semafor) string {
	if semaphore.currentState == Red {
		return colorstring.Color("   [red]R   ")
	} else {
		return colorstring.Color("   [green]G   ")
	}
}

func drawHorizontalSemaphore(semaphore *Semafor) string {
	if semaphore.currentState == Red {
		return colorstring.Color("[red]R")
	} else {
		return colorstring.Color("[green]G")
	}
}
