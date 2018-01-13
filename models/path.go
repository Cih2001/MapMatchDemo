package models

var (
	defaultPath = &Path{
		Name: "default",
	}
)

type Point struct {
	Lat  float32
	Long float32
}

type Path struct {
	Name           string
	OriginalPoints []Point
	MatchedPoints  []Point
}

func (model *Path) AddPoint(lat, lng float32) {
	point := Point{
		Lat:  lat,
		Long: lng,
	}
	model.OriginalPoints = append(model.OriginalPoints, point)
}

func (model *Path) Find(name string) *Path {
	return defaultPath
}

func (model *Path) Match() {
	
}
