package version

const version = "2.10.0"

var (
	gitCommit = ""
	buildTime = ""
)

func Version() string {
	return version
}

func GitCommit() string {
	return gitCommit
}

func BuildTime() string {
	return buildTime
}
