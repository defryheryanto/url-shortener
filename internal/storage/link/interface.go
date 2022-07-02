package link

type IService interface {
	Add(newLink *Link) *Link
	GetLinkByID(id string) *Link
}
