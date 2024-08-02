package builderPattern

import (
	"fmt"
	"testing"
)

func TestBuilderPattern(t *testing.T) {
	hBuilder := new(HouseBuilder)
	house := hBuilder.SetFloor(1).
		SetWindowType("xx").
		SetDoorType("yy").
		Build()

	fmt.Printf("%T %+v \n", house, house)
}
