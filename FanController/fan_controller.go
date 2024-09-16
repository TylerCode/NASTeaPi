package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

const (
	sensePin       = rpio.Pin(17)
	pwmPin         = rpio.Pin(27)
	minDutyCycle   = 20
	maxDutyCycle   = 100
	pwmFrequency   = 25000 // 25 kHz
	updateInterval = 5 * time.Second
	minTemp        = 45.0 // Celsius
	maxTemp        = 75.0 // Celsius
	tempFile       = "/sys/class/thermal/thermal_zone0/temp"
)

func main() {
	if err := rpio.Open(); err != nil {
		log.Fatalf("Failed to open rpio: %v", err)
	}
	defer rpio.Close()

	pwmPin.Mode(rpio.Pwm)
	pwmPin.Freq(pwmFrequency)
	sensePin.Mode(rpio.Input)

	log.Println("Fan controller started")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(updateInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			temp, err := getCPUTemperature()
			if err != nil {
				log.Printf("Error reading CPU temperature: %v", err)
				continue
			}

			dutyCycle := calculateDutyCycle(temp)
			setPWM(dutyCycle)

			log.Printf("CPU Temp: %.1f°C, Fan Speed: %d%%", temp, dutyCycle)

		case <-stop:
			log.Println("Shutting down fan controller")
			setPWM(0)
			return
		}
	}
}

func getCPUTemperature() (float64, error) {
	data, err := os.ReadFile(tempFile)
	if err != nil {
		return 0, err
	}

	var temp float64
	_, err = fmt.Sscanf(string(data), "%f", &temp)
	if err != nil {
		return 0, err
	}

	return temp / 1000.0, nil
}

func calculateDutyCycle(temp float64) int {
	if temp <= minTemp {
		return minDutyCycle
	}
	if temp >= maxTemp {
		return maxDutyCycle
	}

	tempRange := maxTemp - minTemp
	dutyCycleRange := maxDutyCycle - minDutyCycle
	tempRatio := (temp - minTemp) / tempRange

	return minDutyCycle + int(tempRatio*float64(dutyCycleRange))
}

func setPWM(dutyCycle int) {
	pwmPin.DutyCycle(uint32(dutyCycle), 100)
}
