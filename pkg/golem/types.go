package golem

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

// LayerID is the identifier of a layer in the world.
type LayerID uint8

// World is the main interface of Golem.
// It represents the game world and contains all systems and entities classified by layers.
type World interface {
	// Clear reset the whole world to its initial state
	// this will remove all layers, entities, systems
	// and reset the clock.
	Clear()

	// AddLayers adds new layers to the world.
	// If the layer already exists, it will be ignored.
	// Returns the next available LayerID
	AddLayers(layers ...LayerID) LayerID

	// RemoveLayer removes a layer from the world.
	// All entities in the layer will be removed.
	RemoveLayer(layer LayerID)

	// AddEntity adds an Entity to the world.
	// The Entity will be added to the layer it belongs to.
	// If the Entity is already added to another world, it will panic.
	AddEntity(e Entity)

	// AddEntities adds multiple entities to the world.
	// See AddEntity
	AddEntities(e ...Entity)

	// RemoveEntity removes an entity from the world.
	// Entity removal is delayed until the next Flush call.
	// The Entity will be removed from the layer it belongs to.
	// If the Entity is not in the world, it will be ignored.
	RemoveEntity(e Entity)

	// GetEntities returns all entities in the world belonging to the given LayerID.
	GetEntities(layer LayerID) []Entity

	// SetParentEntity sets the parent entity of the world.
	// This is used when the world is a child of another entity.
	SetParentEntity(e Entity)

	// GetParentEntity returns the parent entity of the world.
	// This is used when the world is a child of another entity.
	GetParentEntity() Entity

	GetLayers() []LayerID

	// Size returns the total number of entities in the world.
	// Including the entities of the child worlds.
	Size() int

	// Flush executes all delayed functions, such as adding/removing entities.
	Flush()

	// AddSystem adds a system to the world.
	// The system will be added to the world as a Drawer, Updater, DrawerOnce, or UpdaterOnce.
	// If none of the interfaces are implemented, it will be ignored.
	AddSystem(s system)

	// AddSystems adds multiple Systems to the world.
	// See AddSystem
	AddSystems(s ...system)

	// RemoveSystem removes a system from the world.
	RemoveSystem(s system)

	// Draw will go through all layers
	// Calls all DrawerOnce systems once
	// Calls all Drawer systems for each world Entity
	// Calls are performed in the order of the layers
	Draw(screen *ebiten.Image)

	// Update will go through all layers
	// Calls all UpdaterOnce systems once
	// Calls all Updater systems for each world Entity
	// Calls are performed in the order of the layers
	Update()

	// Freeze will stop the world from updating.
	// Child world are not frozen but won't update if the parent world is frozen.
	// A frozen World will still be drawn.
	Freeze()

	// Unfreeze will allow the world to update.
	Unfreeze()

	// IsFrozen returns the current state of the world.
	IsFrozen() bool
}

// Entity represents an object in the game world.
// Any entity added to a world must be an Entity.
type Entity interface {
	// GetLayer returns the LayerID the entity belongs to
	GetLayer() LayerID

	// setIndex sets the index of the entity in the world
	// this is used for fast world removal
	setIndex(int)

	// index returns the index of the entity in the world
	index() int

	// hasWorld is called when the entity is added to a world
	// returns the current state and invert it
	hasWorld() bool
}

// Clock measure time in a world.
type Clock interface {
	// Now returns the current time in the world.
	// Time elapsed since the world was created, or since it was cleared.
	Now() time.Time

	// Since returns the time elapsed since the given time.
	Since(time.Time) time.Duration

	// Tick advances the clock by one tick.
	// The duration of a tick is defined by ebiten.TPS().
	Tick()

	Elapsed() time.Duration
}

// system internal purpose to allow anything to be a system.
type system interface {
}

// Drawer is a system that draws entities.
type Drawer interface {
	// Draw is called for each entity in the world.
	Draw(e Entity, screen *ebiten.Image, w World)
}

// Updater is a system that updates entities.
type Updater interface {
	// Update is called for each entity in the world.
	Update(e Entity, w World, c Clock)
}

// DrawerOnce is a system that draws once per frame.
type DrawerOnce interface {
	// DrawOnce is called once per frame, before iterating over the entities.
	// This is useful for drawing backgrounds or other static elements.
	// The call is made when the world is starting to draw the layer the system belongs to.
	DrawOnce(screen *ebiten.Image, w World)

	// GetLayer returns the LayerID the system belongs to.
	GetLayer() LayerID
}

// UpdaterOnce is a system that updates once per frame.
type UpdaterOnce interface {
	// UpdateOnce is called once per update iteration, before iterating over the entities.
	UpdateOnce(w World, c Clock)
}
