package dfperms

type Config struct {
	Group []ConfigGroup
}

type ConfigGroup struct {
	Title string
	Level int16
	Tags  []string
}

func (c Config) New() {
	for _, i := range c.Group {
		group := NewGroup(i.Title, i.Level)
		group.WithTags(i.Tags)

		Register(group)
	}
}

func DefaultConfig() Config {
	return Config{
		Group: []ConfigGroup{},
	}
}
