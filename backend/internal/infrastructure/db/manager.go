package db

type Manager struct {
	client DbClient
}

func New(c DbClient) *Manager {
	return &Manager{
		client: c,
	}
}
