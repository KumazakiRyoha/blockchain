package core

type Storage interface {
	Put(block *Block) error
}

type MemoryStore struct {
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

func (m *MemoryStore) Put(b *Block) error {
	return nil
}
