package container

type Bootstrapper interface{
	Initialize()
}

func Initialize(){
	return nil
}