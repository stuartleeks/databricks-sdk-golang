package models

type Library struct {
	Jar   string            `json:"jar,omitempty"`
	Egg   string            `json:"egg,omitempty"`
	Whl   string            `json:"whl,omitempty"`
	Pypi  PythonPyPiLibrary `json:"pypi,omitempty"`
	Maven MavenLibrary      `json:"maven,omitempty"`
	Cran  RCranLibrary      `json:"cran,omitempty"`
}
