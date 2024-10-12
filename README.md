# Golem

> GO Language Entity Management

This is an attempt to create a simple [ECS system](https://en.wikipedia.org/wiki/Entity_component_system) based on
interfaces and some code generation.

It uses [Ebiten engine](https://github.com/hajimehoshi/ebiten) as the rendering engine.

This project is still a work in progress and **IS NOT** production ready, it needs more testing and benchmarks. 

## Installation

```bash
$ go get github.com/t-geindre/golem
$ go install github.com/t-geindre/golem/cmd/golem
```

## Usage

### Layers

Golem uses a layers system. This allows to render and update entities in a specific order.

Layers are identified by a `golem.LayerID` type, which is just an alias for `uint8`.

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

//go:generate golem component Position
type Position struct {
	X, Y float64
}
```

The `///go:generate golem component Position` comment  will generate a `position_golem.go` file containing the required code to retrieve components from entities.

Each time a new component is created, the following command must be run:

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

The component retrieval is based on interfaces, that's why all components must be embedded as 
pointer.

Components of an entity are defined at compilation time, so you can't add or remove
components at runtime. Though you can still set an entity's component to `nil`, which will
make your system act as if the component was not present.

### Systems

TODO

## Utils

TODO

## Examples

See the [examples](./examples) directory.

