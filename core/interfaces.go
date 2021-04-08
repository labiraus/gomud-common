package core

import (
	"context"
)

//Entity is a User or any other interactable entity
type Entity interface {
	ID() uint
	Name() string
	Send(string) <-chan struct{}
	Init(context.Context) Entity
	Home() uint
	EnterString() string
	LeaveString() string
	SetLocation(Location)
	Location() Location
}

//Action is an action that can be performed on an entity
type Action interface {
	Act() <-chan struct{}
}

//Location is a place that a User can be
type Location interface {
	ID() uint
	Name() string
	Description() string
	GetOccupants() <-chan map[uint]Entity
	Move(context.Context, Entity, Direction) <-chan bool
	Act(context.Context, Action) <-chan struct{}
	Exits() string
}
