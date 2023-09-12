package consumer

type Vendor struct {
	Id   int `ddb:"key=hash"`
	Name string
}
