package ar150shiftregout

import gpio "github.com/iketsj/ar150gpio"

const numberOfOutputPinsInAShiftRegister uint8 = 8

const LOW uint8 = 0
const HIGH uint8 = 1

var numberOfOutputPinsMinusOne uint8

type Shiftregister struct {
	numberOfOutputPins uint8
	stateOfOutputPins uint64
	dataInPin gpio.Gpio
	latchPin gpio.Gpio
	clockPin gpio.Gpio
	clearPin gpio.Gpio
}

func NewShiftRegister() Shiftregister {
	return Shiftregister {}
}

func (shiftreg *Shiftregister) Initialize(dataInPinNumber uint8, latchPinNumber uint8, clockPinNumber uint8, clearPinNumber uint8, numOfShiftRegs uint8) {
	(*shiftreg).numberOfOutputPins = numOfShiftRegs * numberOfOutputPinsInAShiftRegister
	numberOfOutputPinsMinusOne = (*shiftreg).numberOfOutputPins - 1	

	(*shiftreg).stateOfOutputPins = 0

	(*shiftreg).dataInPin = gpio.NewGPIO()
	(*shiftreg).latchPin = gpio.NewGPIO()
	(*shiftreg).clockPin = gpio.NewGPIO()
	(*shiftreg).clearPin = gpio.NewGPIO()


	(*shiftreg).dataInPin.Initialize(dataInPinNumber, gpio.OUT)
	(*shiftreg).latchPin.Initialize(latchPinNumber, gpio.OUT)
	(*shiftreg).clockPin.Initialize(clockPinNumber, gpio.OUT)
	(*shiftreg).clearPin.Initialize(clearPinNumber, gpio.OUT)


	(*shiftreg).clockPin.Write(gpio.LOW)
	(*shiftreg).latchPin.Write(gpio.LOW)
	(*shiftreg).clearPin.Write(gpio.HIGH)
}


func (shiftreg *Shiftregister) Write(pinNumber uint8, pinState uint8) {
	if pinState == gpio.HIGH {
		(*shiftreg).stateOfOutputPins |= (1 << pinNumber)	
	}else {
		(*shiftreg).stateOfOutputPins &= ^(1 << pinNumber)
	}
	for i := uint8(0); i < (*shiftreg).numberOfOutputPins; i++ {
		if ((*shiftreg).stateOfOutputPins & (1 << (numberOfOutputPinsMinusOne - i))) != 0 {
			(*shiftreg).dataInPin.Write(gpio.HIGH)
		}else {
			(*shiftreg).dataInPin.Write(gpio.LOW)
		}
		(*shiftreg).clockPin.Write(gpio.HIGH)
		(*shiftreg).clockPin.Write(gpio.LOW)
	}
}

func (shiftreg *Shiftregister) Latch() {
	(*shiftreg).latchPin.Write(gpio.HIGH)
	(*shiftreg).latchPin.Write(gpio.LOW)
}

func (shiftreg *Shiftregister) ClearOutput() {
	(*shiftreg).clearPin.Write(gpio.LOW)
	(*shiftreg).clearPin.Write(gpio.HIGH)
	(*shiftreg).stateOfOutputPins = 0
}


