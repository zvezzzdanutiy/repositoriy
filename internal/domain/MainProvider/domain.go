package MainProvider

type Domain struct {
	anekdotUrl string
}

func New(
	anekdotUrl string,
) *Domain {
	return &Domain{
		anekdotUrl: anekdotUrl,
	}
}
