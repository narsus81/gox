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
	c := Config{port: 222, debug: true, version: "v0.0.16"}
	return &Gox{config: c}
}

func (g *Gox) Version() string {
	return g.config.version
}
