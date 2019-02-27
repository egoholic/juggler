package feed

type Feed struct {
	name string
	sig  string
}

func New(name, sig string) *Feed {
	return &Feed{name, sig}
}
