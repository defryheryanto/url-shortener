package link

type IService interface {
	CreateLink(url string) *Link
}
