package gox

type Gox struct {
	config Config
}

type Config struct {
	port  int
	debug bool
}

func InitGox() *Gox {
	c := Config{port: 888, debug: true}
	return &Gox{config: c}
}
