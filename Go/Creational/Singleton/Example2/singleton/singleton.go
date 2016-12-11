package singleton

type Singleton struct {
	Number      int
	Description string
}

var instance *Singleton

func GetInstance() *Singleton {
	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}
