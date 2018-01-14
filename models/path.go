package models

import (
	"errors"
	"math"

	"github.com/Cih2001/mapmatch"
)

var (
	defaultPath = &Path{
		Name:    "default",
		matcher: new(mapmatch.FastMapMatcher),
	}
)

type Point struct {
	Lat  float64
	Long float64
}

type Path struct {
	Name           string
	matcher        *mapmatch.FastMapMatcher
	OriginalPoints []Point
	MatchedPoints  []Point
}

func (model *Path) AddPoint(lat, lng float64) {
	point := Point{
		Lat:  lat,
		Long: lng,
	}
	model.OriginalPoints = append(model.OriginalPoints, point)
	model.MatchedPoints = append(model.MatchedPoints, point)
}

func (model *Path) Find(name string) *Path {
	return defaultPath
}

func (model *Path) MatchLastPoint() (int, error) {

	pointsCount := len(model.OriginalPoints)
	if pointsCount == 0 {
		return 0, errors.New("not enough points")
	}
	lat, lng := model.OriginalPoints[pointsCount-1].Lat, model.OriginalPoints[pointsCount-1].Long
	points, err := model.matcher.MatchPoint(lat, lng)
	if err != nil {
		return 0, err
	}

	changeIndex := math.MaxInt32
	for _, p := range points {
		if p.Index < changeIndex {
			changeIndex = p.Index
		}
		model.MatchedPoints[p.Index].Lat = p.Latitude
		model.MatchedPoints[p.Index].Long = p.Longitude
	}

	return changeIndex, nil
}
