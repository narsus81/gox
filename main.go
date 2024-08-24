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
	c := Config{port: 444, debug: true, version: "v0.0.17"}
	return &Gox{config: c}
}

func (g *Gox) Version() string {
	return g.config.version
}
