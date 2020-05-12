package storygraph

// Weenie contains story actor data.
type Weenie struct {
	Name   string
	Tags   map[string]string
	Status string
}

// NewWeenie constructs a new Weenie.
func NewWeenie(name, status string, tags map[string]string) *Weenie {
	return &Weenie{
		Name:   name,
		Tags:   tags,
		Status: status,
	}
}

// IsEqualTo compares two weenies for equality.
func (w *Weenie) IsEqualTo(other *Weenie) bool {
	if len(w.Tags) != len(other.Tags) {
		return false
	}

	for k, v := range w.Tags {
		if w, ok := other.Tags[k]; !ok || v != w {
			return false
		}
	}

	return w.Name == other.Name && w.Status == other.Status
}
