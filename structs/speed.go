package speed

type Car struct {
 battery int
 batteryDrain int
 speed int 
 distance  int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed: speed,
		batteryDrain: batteryDrain,
        battery: 100,
		distance: 0,
	}
}

type Track struct {
	distance int
}
// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	// Check if there is enough battery to drive
	if car.battery < car.batteryDrain {
		return car
	}

	// Update distance
	car.distance += car.speed

	// Reduce battery
	car.battery -= car.batteryDrain

	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
    
	// Calculate the maximum distance the car can travel with its current battery level
	maxDistance := car.battery / car.batteryDrain * car.speed

	// Check if the maximum distance is greater than or equal to the track distance
	return maxDistance >= track.distance
    }
