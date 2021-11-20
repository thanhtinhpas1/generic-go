package main

type BaseModel interface {
	GetId() string
	SetId(id string)
}

type Trip struct {
	Id string
}

func (trip *Trip) GetId() string {
	return trip.Id
}

func (trip *Trip) SetId(id string) {
	trip.Id = id
}

type NotModel struct {
}

func TryingToPassGeneric(model BaseModel) {
	// do something
}

func main() {
	notModel := &NotModel{}

	// panic because not a model struct
	TryingToPassGeneric(notModel)

	// can pass to method because it has implemented generic method
	model := &Trip{Id: "Id"}
	TryingToPassGeneric(model)
}
