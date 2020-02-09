package motion

import (
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
)

// DetectHandler event handler whenever the motion sensor detects motion
type DetectHandler func()

// Sensor handles interfacing with the HC-SR501 PIR motion sensor
type Sensor struct {
	pin     gpio.PinIO
	handler DetectHandler
	started bool
	closed  bool
}

// Begin begin listening for motion using the motion sensor.
// Any subsequent 'Begin' calls are ignored
func (me *Sensor) Begin() {
	if !me.started {

		me.started = true

		go func() {
			for me.pin.WaitForEdge(-1) && !me.closed {
				me.handler()
			}
		}()
	}
}

// Close closes this sensor and stops firing any event listeners for detected motion
// Any subsequent calls to 'Close' after the first are ignored.
// The 'Close' method call is ignored if 'Begin' has not yet been called
func (me *Sensor) Close() {

	if me.started && !me.closed {
		me.closed = true
	}
}

// NewSensor creates and initializes a new Motion Sensor
func NewSensor(pin string, handler DetectHandler) (*Sensor, error) {
	s := &Sensor{
		pin:     gpioreg.ByName(pin),
		handler: handler,
		started: false,
		closed:  false,
	}

	if err := s.pin.In(gpio.PullDown, gpio.RisingEdge); err != nil {
		return nil, err
	}

	return s, nil
}
