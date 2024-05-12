package main

import "fmt"

// Subtype polymorphism

type CommonMechanism struct {
	// Dummy
}

func (CommonMechanism) DoWork() {
	// Dummy
}

type Mechanism struct {
	// Composition
	CommonMechanism
}

func (Mechanism) DoWork() {}

// Subtype polymorphism

type MechanismInterface interface {
	DoWork()
}

// Parametric polymorphism

func DoWork[T any](dummy T) {
	fmt.Printf("%v do work...", dummy)
}
