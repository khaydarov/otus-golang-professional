package auth

import ssov1 "github.com/khaydarov/otus-golang-professional/sample_projects/protos/gen/go/sso"

type ServerAPI struct {
	ssov1.UnimplementedAuthServer
}
