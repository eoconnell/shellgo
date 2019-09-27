package schema

type Config struct {
  Build BuildConfig `json:"build"`
}

type BuildConfig struct {
  Language  string `json:"language"`
  Env     []string `json:"env"`
  Install []string `json:"install"`
  Script  []string `json:"script"`
}
