package datastore

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Alice",
			"2": "Bob",
			"3": "Chris",
		},
	}
}

type DataStore interface {
	UserNameForID(userID string) (string, bool)
}
