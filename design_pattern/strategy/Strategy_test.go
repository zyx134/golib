package strategy

import "testing"

func TestStrategy(t *testing.T) {
	aTob := &AToB{ABDistance: 600}

	aTob.Strategy = &BikeStrategy{Speed: 15}
	aTob.Do()

	aTob.Strategy = &BusStrategy{Speed: 90}
	aTob.Do()

	aTob.Strategy = &AirStrategy{Speed: 500}
	aTob.Do()
}
