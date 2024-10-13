# Golem examples

## [Camera](camera)

Demonstrate a simple camera system with:

 - Culling
 - Zooming

### Run

```bash
$ go run github.com/t-geindre/golem/examples/camera@latest
```

## [Nodes](nodes)

Demonstrate how a nodal system can be implemented with Golem.
Each entity is a node that can have children nodes. 
Geometry and transformations are inherited from parent nodes.

### Run

```bash
$ go run github.com/t-geindre/golem/examples/node@latest
```

## [Scenes](scenes)

Demonstrate a scene management system with transitions and lifecycle.

Golem automatically take care of updating and rendering worlds embedded in an entity.

### Run

```bash
$ go run github.com/t-geindre/golem/examples/scenes@latest
```

## [Shmup](shmup)

A very simple shoot them up.

Demonstrates how to create a simple game, handle input and collision detection.

### Run

```bash
$ go run github.com/t-geindre/golem/examples/shmup@latest
```

## [Squares](squares)

A benchmark that renders a lot of squares.

 - Left mouse button to add more squares
 - Right mouse button to remove squares
 - Mouse wheel to increase/decrease the square add/remove rate
 
Mainly inspired by the [Mizu Bunnymark example](https://github.com/sedyh/mizu/tree/main/examples/bunnymark).

### Run

```bash
$ go run github.com/t-geindre/golem/examples/squares@latest
```