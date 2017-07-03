package gc

import (
	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc/db"
)

type resourceCacheCollector struct {
	logger       lager.Logger
	cacheFactory db.ResourceCacheFactory
}

func NewResourceCacheCollector(
	logger lager.Logger,
	cacheFactory db.ResourceCacheFactory,
) Collector {
	return &resourceCacheCollector{
		logger:       logger.Session("resource-cache-collector"),
		cacheFactory: cacheFactory,
	}
}

func (rcuc *resourceCacheCollector) Run() error {
	return rcuc.cacheFactory.CleanUpInvalidCaches()
}
