package repo

import "context"

func (r Repo) Ping(ctx context.Context) (string, error) {
	return "Hello World", nil
}
