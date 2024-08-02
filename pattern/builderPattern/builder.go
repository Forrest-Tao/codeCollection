package builderPattern

type House struct {
	WindowType string
	DoorType   string
	Floor      int
}

type HouseBuilder struct {
	WindowType string
	DoorType   string
	Floor      int
}

func (b *HouseBuilder) SetWindowType(windowType string) *HouseBuilder {
	b.WindowType = windowType
	return b
}
func (b *HouseBuilder) SetDoorType(doorType string) *HouseBuilder {
	b.DoorType = doorType
	return b
}
func (b *HouseBuilder) SetFloor(floor int) *HouseBuilder {
	b.Floor = floor
	return b
}

func (b *HouseBuilder) Build() *House {
	return &House{
		WindowType: b.WindowType,
		DoorType:   b.DoorType,
		Floor:      b.Floor,
	}
}
