package linear_structure

import (
	"errors"
	"reflect"
)

type Queue struct {
	element []interface{}
	size int
}

type QueueAndStack interface {
	Push(element interface{}) error
	Pop() (element interface{},err error)
	Top() (element interface{},err error)
	IsEmpty() bool
}

func NewQueue(element ...interface{}) (*Queue,error) {
	if len(element) != 0 {
		Type := reflect.TypeOf(element[0])
		for _,e := range element[1:] {
			if reflect.TypeOf(e) != Type {
				return &Queue{size:0}, errors.New("type")
			}
		}
	}
	return &Queue{element:element,size:len(element)},nil
}

func (q *Queue) Push(element interface{}) error {
	if !q.IsEmpty() {
		if reflect.TypeOf(q.element[0]) != reflect.TypeOf(element) {
			return errors.New("type")
		}
	}
	q.element = append(q.element,element)
	q.size ++
	return nil
}

func (q *Queue) Pop() (element interface{},err error) {

	if q.IsEmpty(){
		return nil,errors.New("isEmpty")
	}
	element = q.element[0]
	q.element = q.element[1:]
	q.size --
	return element,nil
}

func (q *Queue) Top() (element interface{},err error) {
	if q.IsEmpty(){
		return nil,errors.New("isEmpty")
	}
	return q.element[0],nil
}

func (q *Queue) IsEmpty() bool {
	if q.size == 0 {
		return true
	}
	return false
}