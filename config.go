package dfperms

// Config contains all the groups to be registered from disk
type Config struct {
	Group []ConfigGroup
}

// A proxy struct to load a group
type ConfigGroup struct {
	Title string
	Level int16
	Tags  []string
}

// New applies the settings and groups provided by the config
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
