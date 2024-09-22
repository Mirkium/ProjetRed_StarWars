package jeux

type Lightsaber struct {
	Name        string
	DamageBonus int
	PvBonus     int
	Type        int
	Color       Cristal
	Quantity    int
	Price       int
	CoteForce   int
	Description string
}

type SecondaryLightsaber struct {
	Name        string
	DamageBonus int
	PvBonus     int
	Type        int
	Color       Cristal
	Quantity    int
	Price       int
	CoteForce   int
	Description string
}

type Weapon struct {
	Name        string
	DamageBonus int
	PvBonus     int
	Type        int
	Quantity    int
	Price       int
	CoteForce   int
	Description string
}

type Kit struct {
	Name        string
	Heal        int
	Quantity    int
	Price       int
	Description string
}

type Item struct {
	Name     string
	Quantity int
	Price    int
	Description string
}
