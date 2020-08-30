package builder

import "testing"
import "strconv"

func TestBuilderPattern(t *testing.T) {
	director := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	director.SetBuilder(carBuilder)
	director.Construct()
	car := carBuilder.Build()
	if car.Wheels != 4 {
		t.Errorf("Something went wrong : " + strconv.Itoa(car.Wheels) + " wheels found")
	}

	bikeBuilder := &BikeBuilder{}
	director.SetBuilder(bikeBuilder)
	director.Construct()
	bike := bikeBuilder.Build()
	if bike.Wheels != 2 {
		t.Errorf("Something went wrong : " + strconv.Itoa(bike.Wheels) + " wheels found")
	}
}
