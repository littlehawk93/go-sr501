# HC-SR501 Sensor Go API

Provides a simple, easy-to-use API for interfacing with the HC-SR501 IR motion sensor on the Raspberry Pi platform using GPIO pins.

### Features

- Simple API for easily detecting and responding to motion events
- Leverages [Periph.io]'s GPIO library for low-level communication on Pi hardware
- Customizable pin selection for binding event listeners to

### Installation

```
go get github.com/littlehawk93/go-sr501
```

### Getting started

Basic working example

```
wg := sync.WaitGroup{}
wg.Add(1)

// Bind the sensor signal pin to GPIO Pin #17
sensor, err := motion.NewSensor("17", func() {

    // Print that motion was detected and unlock the wait group
    log.Println("MOTION DETECTED!")
    wg.Done()
})

if err != nil {
    log.Fatalf("Error initializing motion sensor: %s", err.Error())
}

defer sensor.Close()

sensor.Begin()

// Wait until motion has been detected before proceeding
wg.Wait()

```
