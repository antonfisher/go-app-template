package params

// LDFLAGS build properties
var (
	// Name - application name
	// go build -ldflags "-X github.com/<USER_NAME>/<REPOSITORY_NAME>/pkg/params.Name=..."
	Name string

	// Version, to set use flag:
	// go build -ldflags "-X github.com/<USER_NAME>/<REPOSITORY_NAME>/pkg/params.Version=..."
	Version string

	// Commit, to set use flag:
	// go build -ldflags "-X github.com/<USER_NAME>/<REPOSITORY_NAME>/pkg/params.Commit=..."
	Commit string
)
