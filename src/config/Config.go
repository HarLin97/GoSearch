package config

type Config struct {
	RootDirs []string
	Regulars []Regular
	Search   string
	Suffix   []string
}
type Regular struct {
	Description string
	Expression  string
}
