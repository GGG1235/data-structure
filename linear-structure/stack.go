package linear_structure

import (
	"errors"
	"reflect"
)

type Stack struct {
	element []interface{}
	size int
}

func NewStack(element ...interface{}) (*Stack,error) {
	if len(element) != 0 {
		Type := reflect.TypeOf(element[0])
		for _,e := range element[1:] {
			if reflect.TypeOf(e) != Type {
				return &Stack{size:0}, errors.New("type")
			}
		}
	}
	return &Stack{element:element,size:len(element)},nil
}

func (s *Stack) Push(element interface{}) error {
	if !s.IsEmpty() {
		if reflect.TypeOf(s.element[0]) != reflect.TypeOf(element) {
			return errors.New("type")
		}
	}
	s.element = append(s.element,element)
	s.size ++
	return nil
}

func (s *Stack) Pop() (element interface{},err error) {

	if s.IsEmpty(){
		return nil,errors.New("isEmpty")
	}
	element = s.element[s.size-1]
	s.element = s.element[:s.size-1]
	s.size --
	return element,nil
}

func (s *Stack) Top() (element interface{},err error) {
	if s.IsEmpty(){
		return nil,errors.New("isEmpty")
	}
	return s.element[s.size-1],nil
}

func (s *Stack) IsEmpty() bool  {
	if s.size==0 {
		return true
	}
	return false
}