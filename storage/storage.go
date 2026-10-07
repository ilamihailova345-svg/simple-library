package storage

// интерфейс для типов, которые можно сохранять и загружать
type Storable interface {
	Save() error
	Load() error
}
