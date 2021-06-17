package variables

type Variables struct {
	Variables []*Variable
}

type Variable struct {
	Name  string
	Value string
}

func Load() *Variables {
	v := Variables{}
	v.Variables = make([]*Variable, 0)
	return &v
}

func (vs *Variables) Set(name string, value string) {
	for _, v := range vs.Variables {
		if v.Name == name {
			v.Value = value
			return
		}
	}

	v := Variable{
		Name:  name,
		Value: value,
	}
	vs.Variables = append(vs.Variables, &v)
}

func (vs *Variables) Get(name string) (string, bool) {
	for _, v := range vs.Variables {
		if v.Name == name {
			return v.Value, true
		}
	}
	return "", false
}
