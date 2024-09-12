# Golem

> GO Language Entity Management

This is an attempt to create a simple [ECS system](https://en.wikipedia.org/wiki/Entity_component_system) based on interfaces.

It uses [Ebiten engine](https://github.com/hajimehoshi/ebiten) as the rendering engine.

This project **IS NOT** production ready, it needs more testing and benchmarks. 

## Installation

```bash
go get github.com/t-geindre/golem
go install github.com/t-geindre/golem/cmd/golem
```

## Usage

### Layers

Golem uses a layers system. This allows to render and update entities in a specific order.

Layer are identified by `golem.LayerID` type, which is just an alias for `int`.

You can define layers as follows:

```go
package main

import "github.com/t-geindre/golem/pkg/golem"

const (
	LayerBackground = iota
	LayerPlayer
	LayerDebug
)

func main() {
	world := golem.NewWorld()
	world.AddLayers(LayerBackground, LayerPlayer, LayerDebug)
}
```

Layers are rendered and updated in the order they are added.

### Components

Components are just data structures that store data. They should not have any logic.

This can be done as follows:

```go
package component

//go:generate golem position.go
type Position struct {
	X, Y float64
}
```

The `//go:generate golem position.go` comment  will generate a `position_golem.go` file containing the required code to retrieve components from entities.

Each time a new entity is created, the following command must be run:

```bash
$ go generate ./...
```

### Entities

Entities are just a collection of components. They must all embed the Golem entity type to be handled by the engine.

```go
package entity

import (
	"github.com/t-geindre/golem/pkg/golem"
	"component"
)

type Player struct {
	golem.Entity
	*component.Position
}

func NewPlayer() *Player {
	return &Player{
		Entity:   golem.NewEntity(LayerPlayer), // This tells the engine to render and update the entity on the LayerPlayer
		Position: &component.Position{X: 100, Y: 100},
	}
}
```

The component retrieval is based on interfaces, that's why all components must be embedded as pointer.

### Systems

TODO

### Layers


TODO
