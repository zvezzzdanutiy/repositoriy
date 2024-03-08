package api

type Domain struct {
	anekdotProvider AnekdotProvider
}

func New(
	anekdotProvider AnekdotProvider,
) *Domain {
	return &Domain{
		anekdotProvider: anekdotProvider,
	}
}
