package containers

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

type MongoContainer struct {
	testcontainers.Container
	Port string
}

func SetupMongoContainer(ctx context.Context) (*MongoContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        "mongo",
		ExposedPorts: []string{"27017/tcp"},
	}
	container, err := testcontainers.GenericContainer(
		ctx,
		testcontainers.GenericContainerRequest{ContainerRequest: req, Started: true},
	)
	if err != nil {
		return nil, err
	}

	mappedPort, err := container.MappedPort(ctx, "27017")
	if err != nil {
		return nil, err
	}

	return &MongoContainer{Container: container, Port: mappedPort.Port()}, nil
}
