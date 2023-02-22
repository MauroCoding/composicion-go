package items

type item struct {
	id    uint16
	name  string
	price float32
}

func (i item) NewItem(id uint16, name string, price float32) item {
	return item{
		id, name, price,
	}
}
