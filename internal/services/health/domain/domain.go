package domain

type HealthResponse struct {
	Status       string       `json:"status"`
	Dependencies []Dependency `json:"dependencies"`
}

type Dependency struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}
