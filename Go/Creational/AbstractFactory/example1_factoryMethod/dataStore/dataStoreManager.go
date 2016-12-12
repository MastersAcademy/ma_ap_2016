package dataStore

import (
	"log"
	"errors"
	"fmt"
	"strings"
)

type DataStoreManager struct {
	datastoreFactories map[string]DataStoreFactoryMethod
}

func NewDataStoreManager() DataStoreManager{
	manager := DataStoreManager{
		datastoreFactories: make(map[string]DataStoreFactoryMethod),
	}
	manager.Register("postgres", NewPostgreSQLDataStore)
	manager.Register("memory", NewMemoryDataStore)
	return manager
}

func (m DataStoreManager) Register(name string, factory DataStoreFactoryMethod) {
	if factory == nil {
		log.Panicf("Datastore factory %s does not exist.", name)
	}
	_, registered := m.datastoreFactories[name]
	if registered {
		log.Printf("Datastore factory %s already registered. Ignoring.", name)
	}
	m.datastoreFactories[name] = factory
}

func (m DataStoreManager) CreateDatastore(conf Configuration) (DataStore, error) {

	engineName := conf.Get("DATASTORE", "memory")

	engineFactory, ok := m.datastoreFactories[engineName]
	if !ok {
		//Factory not exist
		availableDatastores := make([]string, len(m.datastoreFactories))
		for k, _ := range m.datastoreFactories {
			availableDatastores = append(availableDatastores, k)
		}
		return nil, errors.New(fmt.Sprintf("Invalid Datastore name. Must be one of: %s", strings.Join(availableDatastores, ", ")))
	}

	return engineFactory(conf)
}