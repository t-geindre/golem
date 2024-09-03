package golem

import "github.com/hajimehoshi/ebiten/v2"

type System interface {
}

type Drawer interface {
	Draw(e Entity, screen *ebiten.Image, w World)
}

type Updater interface {
	Update(e Entity, w World)
}

type DrawerOnce interface {
	Draw(screen *ebiten.Image, w World)
	GetLayer() LayerID
}

type UpdaterOnce interface {
	Update(w World)
}
