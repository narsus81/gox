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
	c := Config{port: 777, debug: true, version: "v0.0.22"}
	return &Gox{config: c}
}

func (g *Gox) Version() string {
	return g.config.version
}

func (g *Gox) Port() int {
	return g.config.port
}
