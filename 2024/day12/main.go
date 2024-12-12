package main

import (
	"fmt"

	c "github.com/daanjo3/adventofcode/2024/common"
)

func main() {
	c.AdventCommand("day12",
		calculateFenceCosts,
		c.PlaceholderFunc,
	)
}

type Plot struct {
	Point        c.Point
	Fences       map[c.Ordinal]int
	PerimiterVal int
	Region       int
	Type         rune
}

func spreadRegion(plots [][]*Plot, curPlot *Plot, regionNum int) {
	c.ScanNeighbors(plots, curPlot.Point, func(nplot **Plot, npos c.Point, _ c.Ordinal) {
		if (*nplot).Region != -1 {
			return
		}
		if curPlot.Type == (*nplot).Type {
			(*nplot).Region = regionNum
			spreadRegion(plots, *nplot, regionNum)
		}
	})
}

func checkNeighborFence(plots [][]*Plot, curPlot *Plot, ord c.Ordinal, fenceId int) {
	c.ScanNeighbors(plots, curPlot.Point, func(nplot **Plot, npos c.Point, _ c.Ordinal) {
		if curPlot.Type != (*nplot).Type {
			return
		}
		if val, ok := (*nplot).Fences[ord]; !ok || val != -1 {
			return
		}

		(*nplot).Fences[ord] = fenceId
		checkNeighborFence(plots, *nplot, ord, fenceId)
	})
}

func calculateFenceCosts(inputfile string) {
	plots := [][]*Plot{}
	regionSize := map[int][]*Plot{}

	matrix := c.ReadRuneMatrix(inputfile)
	for y, row := range matrix {
		plotsY := []*Plot{}
		for x, char := range row {
			point := c.Point{X: x, Y: y}
			fences := map[c.Ordinal]int{}
			c.ScanNeighbors(matrix, point, func(nval *rune, npos c.Point, ord c.Ordinal) {
				if *nval != char {
					fences[ord] = -1
				}
			})
			if point.X == 0 {
				fences[c.WEST] = -1
			}
			if point.X == len(matrix[y])-1 {
				fences[c.EAST] = -1
			}
			if point.Y == 0 {
				fences[c.NORTH] = -1
			}
			if point.Y == len(matrix)-1 {
				fences[c.SOUTH] = -1
			}
			plotsY = append(plotsY, &Plot{
				Point:        point,
				Fences:       fences,
				PerimiterVal: len(fences),
				Region:       -1,
				Type:         char,
			})
		}
		plots = append(plots, plotsY)
	}

	regionNum := 0

	// Divide plots in regions
	for _, row := range plots {
		for _, plot := range row {
			if plot.Region != -1 {
				continue
			}
			plot.Region = regionNum
			spreadRegion(plots, plot, regionNum)
			regionNum++
		}
	}

	// Calculate size per region
	for _, row := range plots {
		for _, plot := range row {
			if _, ok := regionSize[plot.Region]; !ok {
				regionSize[plot.Region] = []*Plot{}
			}
			regionSize[plot.Region] = append(regionSize[plot.Region], plot)
		}
	}

	// Calculate costs: perimiter * regionSize
	totalCost := 0
	for _, row := range plots {
		for _, plot := range row {
			totalCost += plot.PerimiterVal * len(regionSize[plot.Region])
		}
	}

	// for _, row := range plots {
	// 	for _, plot := range row {
	// 		// Means not yet visited
	// 		if c.SomeMapKey(plot.Fences) != -1 {
	// 			for ord, fenceId := range plot.Fences {

	// 			}
	// 		}
	// 	}
	// }

	// Loop over de matrix
	// Bij de eerste cel van een unieke regio
	// Loop over de perimiter values
	// Plot.Map[Perimiter] => FenceId
	// Loop over neighbors met zelfde perimiter
	// Als nog geen fenceId: geef fenceId

	fmt.Printf("This overly complicated algorithm sets the total fence cost to %v.\n", totalCost)
}

// 0 => 1
// 1 => 2024
// 2024 => 20 24
// 20 => 2 0
// 24 => 2 4
// 2 => 4048 => 40 48 => 4 0 4 8
// 0 new cycle
// 2 => 4048 => 40 48 => 4 0 4 8
// 4 => 8096 => 80 => 96 => 8 0 9 6
