package context

type SubStore struct {
	Data string
}

func (o *SubStore) Fetch() string {
	return o.Data
}

//func TestServer(t *testing.T) {
//	data := "heelo world "
//}
