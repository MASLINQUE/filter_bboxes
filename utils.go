package main

import "math"

func Filter(bbox []float64, thresh map[string]float64) bool {

	diag := func(x float64, y float64) float64 {
		return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	}
	if bbox[2] > thresh["w"] {
		if bbox[3] > thresh["h"] {
			if diag(bbox[2], bbox[3]) > thresh["d"] {
				return true
			}
		}
	}
	return false

}
