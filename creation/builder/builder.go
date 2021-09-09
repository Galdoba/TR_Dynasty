package builder

type BuildProcess interface {
	SetWheels(int) BuildProcess
	SetDoors() BuildProcess
}

type ManufacturerDirector struct {
	builder BuildProcess
}

func (f *ManufacturerDirector) Construct(i int) {
	//Implementation
	f.builder.SetWheels(i).SetDoors()
}

func (f *ManufacturerDirector) SetBuilder(b BuildProcess) {
	//Implementation
	f.builder = b
}

type VechProduct struct {
	Wheels int
	Doors  int
}

//////////////////

type CarBuilder struct {
	v VechProduct
}

func (c *CarBuilder) SetWheels(i int) BuildProcess {
	c.v.Wheels = i
	return c
}

func (c *CarBuilder) SetDoors() BuildProcess {
	return nil
}

func (c *CarBuilder) Build() VechProduct {
	return VechProduct{}
}
