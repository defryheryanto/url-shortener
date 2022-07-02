package link

import (
	"github.com/defryheryanto/url-shortener/internal/link"
	linkstorage "github.com/defryheryanto/url-shortener/internal/storage/link"
	randomstring "github.com/defryheryanto/url-shortener/internal/stringhelper/random"
)

type LinkService struct {
	storage  linkstorage.IService
	idLength int
}

func NewService(storage linkstorage.IService, idLength int) *LinkService {
	return &LinkService{storage, idLength}
}

func (s *LinkService) CreateLink(url string) *link.Link {
	randomStringService := randomstring.NewService("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", s.idLength)
	newUniqueId := randomStringService.GenerateRandomString()
	existingLink := s.storage.GetLinkByID(newUniqueId)

	for existingLink != nil {
		newUniqueId = randomStringService.GenerateRandomString()
		existingLink = s.storage.GetLinkByID(newUniqueId)
	}

	newLink := &linkstorage.Link{
		Id:  newUniqueId,
		Url: url,
	}

	newLink = s.storage.Add(newLink)

	return &link.Link{
		Id:  newLink.Id,
		Url: newLink.Url,
	}
}

func (s *LinkService) GetLink(uniqueId string) *link.Link {
	existingLink := s.storage.GetLinkByID(uniqueId)
	if existingLink == nil {
		return nil
	}

	return &link.Link{
		Id:  existingLink.Id,
		Url: existingLink.Url,
	}
}
