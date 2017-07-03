package image

import (
	"errors"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc/creds"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/worker"
)

var ErrUnsupportedResourceType = errors.New("unsupported resource type")

type imageFactory struct {
	imageResourceFetcherFactory ImageResourceFetcherFactory
}

func NewImageFactory(
	imageResourceFetcherFactory ImageResourceFetcherFactory,
) worker.ImageFactory {
	return &imageFactory{
		imageResourceFetcherFactory: imageResourceFetcherFactory,
	}
}

func (f *imageFactory) GetImage(
	logger lager.Logger,
	workerClient worker.Worker,
	volumeClient worker.VolumeClient,
	imageSpec worker.ImageSpec,
	teamID int,
	delegate worker.ImageFetchingDelegate,
	resourceUser db.ResourceUser,
	resourceTypes creds.VersionedResourceTypes,
) (worker.Image, error) {
	if imageSpec.ImageArtifactSource != nil {
		artifactVolume, existsOnWorker, err := imageSpec.ImageArtifactSource.VolumeOn(workerClient)
		if err != nil {
			logger.Error("failed-to-check-if-volume-exists-on-worker", err)
			return nil, err
		}

		if existsOnWorker {
			return &imageProvidedByPreviousStepOnSameWorker{
				artifactVolume: artifactVolume,
				imageSpec:      imageSpec,
				teamID:         teamID,
				volumeClient:   volumeClient,
			}, nil
		}

		return &imageProvidedByPreviousStepOnDifferentWorker{
			imageSpec:    imageSpec,
			teamID:       teamID,
			volumeClient: volumeClient,
		}, nil
	}

	// check if custom resource
	resourceType, found := resourceTypes.Lookup(imageSpec.ResourceType)
	if found {
		// source, err := resourceType.Source.Evaluate()
		// if err != nil {
		// 	return nil, err
		// }
		//
		imageResourceFetcher := f.imageResourceFetcherFactory.NewImageResourceFetcher(
			workerClient,
			resourceUser,
			worker.ImageResource{
				Type:   resourceType.Type,
				Source: resourceType.Source,
			},
			resourceType.Version,
			teamID,
			resourceTypes.Without(imageSpec.ResourceType),
			delegate,
		)

		return &imageFromResource{
			imageResourceFetcher: imageResourceFetcher,

			privileged:   resourceType.Privileged,
			teamID:       teamID,
			volumeClient: volumeClient,
		}, nil
	}

	if imageSpec.ImageResource != nil {
		imageResourceFetcher := f.imageResourceFetcherFactory.NewImageResourceFetcher(
			workerClient,
			resourceUser,
			*imageSpec.ImageResource,
			nil,
			teamID,
			resourceTypes,
			delegate,
		)

		return &imageFromResource{
			imageResourceFetcher: imageResourceFetcher,

			privileged:   imageSpec.Privileged,
			teamID:       teamID,
			volumeClient: volumeClient,
		}, nil
	}

	if imageSpec.ResourceType != "" {
		return &imageFromBaseResourceType{
			worker:           workerClient,
			resourceTypeName: imageSpec.ResourceType,
			teamID:           teamID,
			volumeClient:     volumeClient,
		}, nil
	}

	return &imageFromRootfsURI{
		url: imageSpec.ImageURL,
	}, nil
}
