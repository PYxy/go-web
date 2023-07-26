package store

var (
	mysqlClient Factory
	redisClient Factory
)

type Factory interface {
	CustomerInfoOption() CustomerInfoStore
	CustomerGoodOption() CustomerGoodStore
	Close() error
}

// Client return the store client instance.
func Client(storeType string) Factory {
	switch storeType {
	case "mysql":
		return mysqlClient
	case "redis":
		return redisClient
	default:
		return mysqlClient
	}
}

// SetClient set the iam store client.
func SetClient(storeType string, factory Factory) {
	switch storeType {
	case "mysql":
		mysqlClient = factory
	case "redis":
		redisClient = factory
	default:
		mysqlClient = factory
	}
}
