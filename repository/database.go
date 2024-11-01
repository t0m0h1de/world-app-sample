package repository

type Database interface {
	Connect(string, string, string, int, bool) error
	Close()
	FindPopulation(city string) (int, error)
}
