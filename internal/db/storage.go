package db

type Storage struct {
	Filename string
}

func newStorage(filename string) (*Storage, error) {
	return &Storage{
		Filename: filename,
	}, nil
}

func (s *Storage) Write() {
}

func (s *Storage) Read() {
}
