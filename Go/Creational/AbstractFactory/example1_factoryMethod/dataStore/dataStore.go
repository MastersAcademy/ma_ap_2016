package dataStore

import "errors"

var UserNotFoundError = errors.New("User not found")

type DataStore interface {
	Name() string
	FindUserNameById(id int64) (string, error)
}

type DataStoreFactoryMethod func(conf Configuration) (DataStore, error)

type Configuration struct {
	data map[string]string
}

func (conf Configuration)Get(key string, defValue string) string {
	v, ok := conf.data[key]
	if ok {
		return  v
	} else {
		return defValue
	}
}

func (conf Configuration) Set(key string, value string) {
	conf.data[key] = value
}

func NewConfiguration() Configuration {
	return Configuration{
		data: make(map[string]string),
	}
}
