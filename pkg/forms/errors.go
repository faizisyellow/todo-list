package forms

// The name of the form field will be used as the key in this map.
type errors map[string][]string

// To add error messages for a given field to the errors map
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// To retrieve the first error message for a given field from the errors map.
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}

	return es[0]
}
