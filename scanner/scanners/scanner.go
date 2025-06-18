package scanners

type Scanner interface {
	Name() string
	Run(target string) (string, error)
}