package handlers

import (
	"golang-template/configs"
	"golang-template/modules/common/repositories/cache"
	"golang-template/modules/examples/repositories"
)

type handler struct {
	cfg   *configs.Configs
	store repositories.ExampleRepository
	cache cache.Cached
}

func ExampleHandler(cfg *configs.Configs, store repositories.ExampleRepository, cache cache.Cached) *handler {
	return &handler{cfg: cfg, store: store, cache: cache}
}
