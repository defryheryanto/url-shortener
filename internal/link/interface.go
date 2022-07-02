package link

type IService interface {
	CreateLink(url string) *Link
	GetLink(uniqueId string) *Link
}
