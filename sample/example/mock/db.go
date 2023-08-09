package mock

type DB interface {
	Get(key string) (int, error)
}

func GetFromDB(db DB, key string) int {
	if value, err := db.Get(key); err == nil {
		return value
	}

	return -1
}

/*
$ go get github.com/golang/mock/gomock
$ go install github.com/golang/mock/mockgen@v1.6.0
$ mockgen -source=db.go -destination=db_mock.go -package=mock
*/
