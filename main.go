package gox

type Gox struct {
	config Config
}

type Config struct {
	port int
}

func InitGox() *Gox {
	c := Config{port: 888}
	return &Gox{config: c}
}
