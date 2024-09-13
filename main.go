package main

import (
	"fmt"
	"sync"
)

// VehicleType represents different types of vehicles
type VehicleType int

const (
	MotorCycle VehicleType = iota
	Car
	Truck
)

// ParkingSpot represents individual parking spot.
type ParkingSpot struct {
	spotNumber   int
	vehicleType  VehicleType
	availability bool
}

type Vehicle struct {
	vehicleNumber string
	vehicleType   VehicleType
}

type ParkingLot struct {
	spots []ParkingSpot
	mu    sync.Mutex
}

func NewParkingSpot(spotNumber int, vehicleType VehicleType) ParkingSpot {
	return ParkingSpot{
		spotNumber:   spotNumber,
		vehicleType:  vehicleType,
		availability: true,
	}
}

func NewParkingLot(totalSpots int) ParkingLot {
	spots := make([]ParkingSpot, totalSpots)
	for i := 0; i < totalSpots; i++ {
		spots[i] = NewParkingSpot(i+1, Car)
	}
	return ParkingLot{
		spots: spots,
	}
}

func (p *ParkingLot) ParkVehicle(vehicle Vehicle) bool {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i := range p.spots {
		if p.spots[i].availability && p.spots[i].vehicleType == vehicle.vehicleType {
			p.spots[i].availability = false
			fmt.Printf("Parking => Vehicle %s Parked at %d\n", vehicle.vehicleNumber, p.spots[i].spotNumber)
			return true
		}
	}
	fmt.Printf("Parking Lot is full\n")
	return false
}

func (p *ParkingLot) UnparkVehicle(vehicle Vehicle) bool {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i := range p.spots {
		if !p.spots[i].availability && vehicle.vehicleType == p.spots[i].vehicleType {
			p.spots[i].availability = true
			fmt.Printf("Unparking => Vehicle %s Unparked at %d\n", vehicle.vehicleNumber, p.spots[i].spotNumber)
			return true
		}
	}
	fmt.Printf("Vehicle not found in Parking Lot.\n")
	return false
}

func main() {

	parkingLot := NewParkingLot(20)
	
	// Simulate vehicle arrivals from multiple gates using concurrency
	var wg sync.WaitGroup
	numVehicles := 10

	wg.Add(numVehicles)

	for i := 0; i < numVehicles; i++ {
		go func(vehicleIndex int) {
			defer wg.Done()

			v := Vehicle{
				vehicleNumber: fmt.Sprintf("KL %d", (vehicleIndex + 1)),
				vehicleType:   Car,
			}
			parkingLot.ParkVehicle(v)
			parkingLot.UnparkVehicle(v)
		}(i)
	}

	wg.Wait()

}
