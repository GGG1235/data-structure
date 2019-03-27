package linear_structure

import (
	"errors"
	"reflect"
)

type ArrayList struct {
	values []interface{}
	len int
}

func NewArrayList(values ...interface{}) (*ArrayList,error) {
	if len(values) != 0 {
		Type := reflect.TypeOf(values[0])
		for _,e := range values[1:] {
			if reflect.TypeOf(e) != Type {
				return &ArrayList{len:0}, errors.New("type")
			}
		}
	}
	return &ArrayList{values:values,len:len(values)},nil
}

func (arraylist *ArrayList) Append(value interface{}) error {
	if arraylist.len != 0 && reflect.TypeOf(arraylist.values[0]) != reflect.TypeOf(value) {
		return errors.New("type")
	}
	arraylist.values=append(arraylist.values,value)
	arraylist.len++
	return nil
}

func (arraylist *ArrayList) InsertByIndex(index int,value interface{}) error {
	if index < 0 || index >= arraylist.len {
		return errors.New("rm error")
	}
	values := make([]interface{},0)
	values = append(values,value)
	values = append(values,arraylist.values[index:]...)
	arraylist.values = append(arraylist.values[:index],values...)
	arraylist.len ++
	return nil
}

func (arraylist *ArrayList) RmByIndex(index int) error {
	if index < 0 || index >= arraylist.len {
		return errors.New("rm error")
	}
	if arraylist.len == 0 {
		return errors.New("out of range")
	}
	arraylist.len --
	arraylist.values = append(arraylist.values[:index],arraylist.values[index+1:]...)
	return nil
}

func (arraylist *ArrayList) RmByValue(value interface{}) error {
	if arraylist.len != 0 && reflect.TypeOf(arraylist.values[0]) != reflect.TypeOf(value) {
		return errors.New("type")
	}
	if arraylist.len == 0 {
		return errors.New("out of range")
	}
	for i,e := range arraylist.values {
		if e == value {
			arraylist.len --
			arraylist.values = append(arraylist.values[:i],arraylist.values[i+1:]...)
			return nil
		}
	}
	return nil
}