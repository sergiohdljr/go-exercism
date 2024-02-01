package numbers

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var sucessRatePercent = successRate / 100
    var carsPerHours = float64(productionRate) * sucessRatePercent
	
    return carsPerHours
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var productionPerHour = CalculateWorkingCarsPerHour(productionRate,successRate)
	var carPerMinute = int(productionPerHour)/ 60
	return carPerMinute
}

// CalculateCost works out the cost of producing the given number of cars.
// working...
func CalculateCost(carsCount int) uint {
	
	var cost uint

	const groupOfTeenCarsPrice = 95.000
	const individualCarsPrice = 10.000

	var groupOfTen = carsCount/10
	var remaining = carsCount % 10

	if(groupOfTen > 0 && remaining > 0){
		cost := uint( int(groupOfTen * groupOfTeenCarsPrice) + int(remaining * individualCarsPrice) )
		return cost  
	}  
	
	if(groupOfTen < 0){
		cost := individualCarsPrice * float64(remaining)
		return uint(cost)
	}

	return cost
}