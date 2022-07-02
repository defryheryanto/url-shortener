package link

var links []*Link

type LinkStorageService struct{}

func (s *LinkStorageService) Add(newLink *Link) *Link {
	links = append(links, newLink)
	return newLink
}

func (s *LinkStorageService) GetLinkByID(id string) *Link {
	for _, l := range links {
		if l.Id == id {
			return l
		}
	}

	return nil
}
