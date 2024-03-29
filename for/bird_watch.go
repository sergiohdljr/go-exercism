package forGo


// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	count := 0
	for i := 0; i < len(birdsPerDay); i++ {
		count = count + birdsPerDay[i]
	}
	return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	startIndex := (week - 1) * 7
	endIndex := week * 7
	totalBirds := 0
	for i := startIndex; i < endIndex; i++ {
		totalBirds += birdsPerDay[i]
	}

    return totalBirds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	
	for i := 0; i < len(birdsPerDay); i++ {
		if(i == 0){
			birdsPerDay[i] = birdsPerDay[i] + 1
		} else if (i % 2 == 0) {
			birdsPerDay[i] = birdsPerDay[i] + 1
		}
	}
	return birdsPerDay
}
