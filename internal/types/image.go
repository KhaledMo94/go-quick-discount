package types

type Image struct {
	Image *string
}

func (i Image) URL(baseURL string) string {
	if i.Image == nil {
		return ""
	}
	return baseURL + "/storage/" + *i.Image
}