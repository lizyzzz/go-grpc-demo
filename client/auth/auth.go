package auth

import "context"

type Authentication struct {
	User     string
	Password string
}

func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{"user": a.User, "password": a.Password}, nil
}

// 是否开启安全认证
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}
