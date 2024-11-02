package repository

type Database interface {
	Connect(string, string, string, string, int, bool)
	Close()
	FindCityPopulation(city string) (int, error)
}
