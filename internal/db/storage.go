package db

type Storage struct {
	Filename string
}

func newStorage(filename string) *Storage {
	return &Storage{
		Filename: filename,
	}
}

func (s *Storage) Write() {
}

func (s *Storage) Read() {
}
