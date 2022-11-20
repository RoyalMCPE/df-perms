package dfperms

func NewGroup(title string, level int16) *Group {
	return &Group{
		title: title,
		level: level,
		tags:  []string{},
	}
}

// Group holds pre-defined tags and a permission level that can be applied to a session.
// Sessions may have multiple groups and group will take precidence if it has a higher level
type Group struct {
	// TODO: Implement parents
	title string
	level int16
	tags  []string
}

func (g Group) Title() string {
	return g.title
}

func (g *Group) WithTags(tags []string) {
	g.tags = append(g.tags, tags...)
}
