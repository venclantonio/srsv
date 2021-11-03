package lab2

import "time"

type Generator struct {
	semaphores map[string]*Semafor
}

func newGen(semaphores map[string]*Semafor) *Generator {
	return &Generator{
		semaphores: semaphores,
	}
}

func (g *Generator) Start(char rune) {
	for {
		for _, v := range g.semaphores {
			time.Sleep(10 * time.Second)
			v.vehicleChan <- char
		}
	}
}
