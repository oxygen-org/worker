package task


type Callback struct{

}

func (task *Callback)Do()*Callback{
	return task
}