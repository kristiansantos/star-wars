package entities

type PlanetCreateDto struct {
	Name    string `json:"name"`
	Climate string `json:"climate"`
	Ground  string `json:"ground"`
}

type FilmsResults struct {
	Films []string `json:"films"`
}

type FilmsResponseBody struct {
	Results []FilmsResults `json:"results"`
}
