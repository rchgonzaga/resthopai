package kittiesbundle

// KittiesMapper define the base contract for mapper
type KittiesMapper interface {
	FindAll() ([]Kitty, error)
	Insert(*Kitty) error
	Delete(int) error
}
