package app

import "github.com/defryheryanto/url-shortener/internal/link"

type Application struct {
	LinkService link.IService
}
