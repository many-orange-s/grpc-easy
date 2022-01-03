package freedefinit

import (
	"context"
	"encoding/base64"
)

func (b *BasicAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	auth := b.Secret
	enc := base64.StdEncoding.EncodeToString([]byte(auth))
	u := make(map[string]string, 1)
	u["Product"] = "Basic " + enc
	return u, nil
}

func (b *BasicAuth) RequireTransportSecurity() bool {
	return true
}
