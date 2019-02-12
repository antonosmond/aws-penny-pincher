package config

type Config struct {
	Rules []Rule `json:"rules"`
}

type Rule struct {
	Name      string     `json:"name"`
	Regions   []string   `json:"regions"`
	Resources []Resource `json:"resources"`
}

type Resource struct {
	Type    string   `json:"type"`
	Filters []Filter `json:"filters"`
	Actions []string `json:"actions"`
}

type Filter struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}
