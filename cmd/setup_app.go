package main

import (
	"github.com/defryheryanto/url-shortener/internal/app"
	"github.com/defryheryanto/url-shortener/internal/link/v1"
	linkstorage "github.com/defryheryanto/url-shortener/internal/storage/link"
)

type storages struct {
	linkStorage linkstorage.IService
}

func buildApp() *app.Application {
	appStorage := buildStorage()
	linkService := link.NewService(appStorage.linkStorage, 8)

	return &app.Application{
		LinkService: linkService,
	}
}

func buildStorage() *storages {
	return &storages{
		linkStorage: &linkstorage.LinkStorageService{},
	}
}
