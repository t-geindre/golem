# Golem examples

## [Scenes](scenes)

Demonstrates a very simple scene management system.

Click left mouse button to switch between scenes.

It can be fairly improved by adding `setup()` and `teardown()` methods
to the scenes to make sure all resources are released when the scene is switched.

The rendering system in each scene take care of the scene position, so moving the scene will
also move all its contents. Can be useful from creating a transition effect between scenes.

### Run

```bash
$ go run github.com/t-geindre/golem/examples/scenes@latest
```

### Todo
 - Add a transition system

## [Shmup](shmup)

A very simple shoot them up.

Demonstrates how to create a simple game, handle input and collision detection.

### Run

```bash
$ go run github.com/t-geindre/golem/examples/shmup@latest
```

### Todo
 - Display score
 - Display HP
 - Add a background
 - Add more enemies

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