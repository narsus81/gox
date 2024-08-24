package gox

type Gox struct {
	config Config
}

type Config struct {
	port    int
	debug   bool
	version string
}

func InitGox() *Gox {
	c := Config{port: 888, debug: true, version: "0.0.1"}
	return &Gox{config: c}
}

func (g *Gox) Version() string {
	return g.config.version
}
