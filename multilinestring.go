package geogoth

// MultiLineString ...
type MultiLineString struct {
	Coords [][][]float64
}

// NewMultiLineString creates MultiLineString
func NewMultiLineString(coords [][][]float64) MultiLineString {
	return MultiLineString{
		Coords: coords,
	}
}

// Coordinates returns array of longitude, latitude of the MultiLineString
func (m MultiLineString) Coordinates() interface{} {
	return m.Coords // longitude (Y), latitude (X)
}

// GetCoordinates returns array of longitude, latitude of MultiLineString
// coordnum - index of coordinate arr
func (m MultiLineString) GetCoordinates(setnum, coordnum int) (float64, float64) {
	coords := (m.Coordinates()).([][][]float64)

	lon := coords[setnum][coordnum][0]
	lat := coords[setnum][coordnum][1]

	return lon, lat // longitude (Y), latitude (X)
}

// Type returns type of the MultiLineString (MultiLineString)
func (m MultiLineString) Type() string {
	return "MultiLineString"
}

// Length returns length of the MultiLineString
func (m MultiLineString) Length() float64 {
	return MultiLineStringLength(m)
}

// DistanceTo returns distance between two geo objects
func (m MultiLineString) DistanceTo(f Feature) float64 {

	var distance float64

	switch f.Type() {
	case "Point":
		point := f.(*Point)
		distance = DistancePointMultiLineString(point, &m)

	case "MultiPoint":
		mpoint := f.(*MultiPoint)
		distance = DistanceMultiPointMultiLinestring(mpoint, &m)

	case "LineString":
		lstr := f.(*LineString)
		distance = DistanceLineStringMultiLineString(lstr, &m)

	case "MultiLineString":
		mlstr := f.(*MultiLineString)
		distance = DistanceMultiLineStringMultiLineString(&m, mlstr)

	case "Polygon":
		pol := f.(*Polygon)
		distance = DistanceMultiLineStringPolygon(&m, pol)

	case "MultiPolygon":
		mpol := f.(*MultiPolygon)
		distance = DistanceMultiLineStringMultiPolygon(&m, mpol)

	}

	return distance
}

// IntersectsWith returns true if geoObject intersects with Feature
func (m MultiLineString) IntersectsWith(f Feature) bool {
	var intersection bool

	switch f.Type() {
	case "Point":
		// point := f.(*Point)

	case "MultiPoint":
		// mpoint := f.(*MultiPoint)

	case "LineString":
		// lstr := f.(*LineString)

	case "MultiLineString":
		// mlinestr := f.(*MultiLineString)

	case "Polygon":
		// polygon := f.(*Polygon)

	case "MultiPolygon":
		// mpolyg := f.(*MultiPolygon)
	}

	return intersection

}
