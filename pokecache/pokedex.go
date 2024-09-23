package pokecache

type Pokedex struct {
	Height          int
	Weight          int
	Hp              int
	Attack          int
	Defense         int
	Special_attack  int
	Special_defense int
	Speed           int
	Name            string
	Types           []string
}

func NewPokedex() map[string]Pokedex {
	p := map[string]Pokedex{}
	return p
}
