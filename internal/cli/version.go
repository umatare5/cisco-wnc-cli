package cli

// version is stamped at link time by the Makefile and by goreleaser. The default
// marks a build made straight from a working tree.
var version = "dev"

// Version reports the build's version, which is also the suffix UserAgent carries so a
// controller's own logs identify the client.
func Version() string {
	return version
}

// UserAgent is the name every RESTCONF request sends.
func UserAgent() string {
	return "wnc/" + version
}
