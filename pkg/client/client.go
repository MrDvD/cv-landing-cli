package client

import (
	"fmt"
	"net/http"
)

type BaseClient struct {
	HTTPClient *http.Client
	Endpoints  map[string]string
}

func (b *BaseClient) Resolve(key string, path string) (string, error) {
	base, ok := b.Endpoints[key]
	if !ok {
		return "", fmt.Errorf("could not resolve key '%s' from config", key)
	}
	return fmt.Sprintf("%s/%s", base, path), nil
}
