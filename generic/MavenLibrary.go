package generic

type MavenLibrary struct {
	Coordinates string `json:"coordinates,omitempty"`
	Repo string `json:"repo,omitempty"`
	Exclusions []string `json:"exclusions,omitempty"`
}