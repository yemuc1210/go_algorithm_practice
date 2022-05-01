package structs

import "errors"


type Queue struct{
	maxSize int
	array []interface{}
	front int    //队首下标
	rear int	//队尾下标
	curSize int
}

//增删查改四个基本操作
func (q *Queue)AddQueue(val interface{}) (error error){
	//判满
	if q.curSize == q.maxSize {
		return errors.New("queue full")
	}

	q.rear ++
	q.curSize ++
	q.array = append(q.array, val)
	return 
}


func (q *Queue) GetQueue()(val interface{}, err error) {
	//判kong
	if q.curSize == 0{
		return -1,errors.New("queue empty")
	}

	q.front ++
	q.curSize --
	return q.array[q.front],nil
}

func (q *Queue) Len() int{
	return q.curSize
}

func (q *Queue) IsEmpty()bool{
	return q.curSize == 0
}

