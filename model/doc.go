package model

type Doc struct {
	items    []any
	vars     map[string]Values
	commands map[string]Commands
}

func New(cap int) *Doc {
	return &Doc{
		items:    make([]any, 0, cap),
		vars:     make(map[string]Values, cap),
		commands: make(map[string]Commands, cap),
	}
}

func (d *Doc) AppendAssignment(e *Assignment) {
	d.items = append(d.items, e)
	d.vars[e.Key] = append(d.vars[e.Key], e.Value...)
}

func (d *Doc) AppendCommand(c *Command) {
	d.items = append(d.items, c)
	d.commands[c.Name] = append(d.commands[c.Name], c)
}

func (d *Doc) Vars() map[string]Values {
	return d.vars
}

func (d *Doc) Get(key string) Values {
	return d.vars[key]
}

func (d *Doc) Commands(name string) Commands {
	return d.commands[name]
}

func (d *Doc) Items() []any {
	return d.items
}
