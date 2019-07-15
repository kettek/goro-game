package entity

type BasicFighter struct {
	hp, maxHp int
	defense   int
	power     int
}

func (f *BasicFighter) Hp() int {
	return f.hp
}

func (f *BasicFighter) MaxHP() int {
	return f.maxHp
}
func (f *BasicFighter) Defense() int {
	return f.defense
}
func (f *BasicFighter) Power() int {
	return f.power
}
