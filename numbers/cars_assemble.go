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
func CalculateCost(carsCount int) uint {
	const groupOfTeenCarsPrice = 95000
	const individualCarsPrice = 10000
    var groupOfTen = carsCount / 10 
    var carsRemaning = carsCount % 10
    
	var carsCost uint

    if(carsCount >= 10){
        carsCost := (float64(groupOfTen) * groupOfTeenCarsPrice) + (float64(carsRemaning) * individualCarsPrice)
        return  uint(carsCost) 
    } 
    
	carsCost = uint(individualCarsPrice * carsCount)
	  
    
    return carsCost
	
}