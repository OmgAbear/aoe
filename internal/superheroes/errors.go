package superheroes

type InvalidSuperpowerError struct {
	error
}

func (err InvalidSuperpowerError) Error() string {
	return "the only allowed superpowers are strength, speed, flight, invulnerability and healing"
}
