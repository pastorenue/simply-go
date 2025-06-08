package main

import (
	"fmt"
)


type Resource struct {
	name string
}


func Acquire(name string) *Resource {
	fmt.Println("Acquiring resource: ", name)
	return &Resource{name: name}
}

// This is like a method in languages like Python. where r *Resource
// can be synonymous to self. Use (r *Resource) if you intend to modify the pointer variable
// else use (r Resource)
func (r *Resource) Release() {
	r.name = "GO"
	fmt.Println("Releasing resource: ", r.name)
}

// This is like a static/class method.
func (r Resource) DoNothing() {
	r.name = "TYE"
	fmt.Println("This is doing nothing %v", r.name)
}

func WithResource(name string, fn func(*Resource) error) error {
	resource := Acquire(name)
	defer resource.Release()
	resource.DoNothing()

	return fn(resource)
}
