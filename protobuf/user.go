package protobuf

import "github.com/dundee/gogenerate/pkg/log"

func (u *User) MarshalLog() log.Fields {
	return log.Fields{
		"user.id":   u.Id,
		"user.name": u.Name,
	}
}
