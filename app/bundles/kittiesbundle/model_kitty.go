package kittiesbundle

// Kitty struct
type Kitty struct {
	ID        int               `json:"id" gorm:"AUTO_INCREMENT"`
	Name      string            `json:"name"`
	Breed     string            `json:"breed"`
	BirthDate string            `json:"birthDate"`
	Errors    map[string]string `json:"-" gorm:"-"`
}

// NewKitty create a new Kitty
func NewKitty(name string, breed string, birthDate string) *Kitty {
	return &Kitty{
		Name:      name,
		Breed:     breed,
		BirthDate: birthDate,
	}
}

// Validate a Kitty
func (k *Kitty) Validate() bool {
	k.Errors = make(map[string]string)

	if k.Name == "" {
		k.Errors["name"] = "name can not be empty"
	}

	if len(k.Errors) > 0 {
		return false
	}

	return true
}
