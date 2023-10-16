package logic

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForId(id string) (string, bool) {
	name, ok := sds.userData[id]
	return name, ok
}

func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "John",
			"2": "Jane",
		},
	}
}