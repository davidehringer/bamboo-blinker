package main

import (
	"fmt"
	"github.com/davidehringer/bamboo-blinker/bunny"
	"github.com/davidehringer/goblync"
	"os"
	"os/signal"
	"syscall"
	"time"
	"strconv"
)

func main() {
	numArgs := len(os.Args)
	if numArgs != 2 && numArgs != 3 {
		fmt.Println("Usage: bamboo-blinker URL [INTERVAL_SECONDS]")
		os.Exit(1)
	}
	url := os.Args[1]
	buildBunny := bunny.NewBunny(url)

	interval := 10
	if numArgs == 3 {
		value, err := strconv.ParseInt(os.Args[2], 10, 32)
		if err != nil {
			fmt.Println("INTERVAL_SECONDS must be an integer")
			os.Exit(1)
		}
		interval = int(value)
	}

	light := blync.NewBlyncLight()
	light.SetColor(blync.Green)

	// clean shutdown of light
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		light.Close()
		os.Exit(1)
	}()

	monitor := NewMonitor(light)
	for {
		if buildBunny.Healthy() {
			monitor.SetHealthy()
		} else {
			monitor.SetUnhealthy()
		}
		time.Sleep(time.Second * time.Duration(interval))
	}
}

type monitor struct {
	healthy bool
	light blync.BlyncLight
}

func NewMonitor(light blync.BlyncLight) (m monitor) {
	m.healthy = true;
	m.light = light
	return
}

func (m *monitor) SetHealthy() {
	if !m.healthy {
		m.healthy = true
		m.light.StopPlay()
		m.light.SetBlinkRate(blync.BlinkOff)
		for i := 0; i < 256; i++ {
			m.light.SetColor([3]byte{255 - byte(i), byte(i), 0x00})
			time.Sleep(13 * time.Millisecond)
		}
		m.light.Play(28)
	}
}

func (m *monitor) SetUnhealthy() {
	if m.healthy {
		m.healthy = false
		for i := 0; i < 256; i++ {
			m.light.SetColor([3]byte{byte(i), 255 - byte(i), 0x00})
			time.Sleep(13 * time.Millisecond)
		}
		m.light.SetBlinkRate(blync.BlinkMedium)
		m.light.Play(52)
		// We using a never ending sound because it was one of the only ones that 
		// had some sound of urgency to it.  But we don't want it to keep playing
		time.Sleep(time.Second * 15)
		m.light.StopPlay()
	}
}
