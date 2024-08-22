// Сущности слоя БД

package database

import "log/slog"

type database struct {
	Database string
}

type cardsStorage struct {
	log *slog.Logger
	// db          *database.Database
	loggerLevel slog.Attr
}

// type CardsStorage interface {
// 	CreateCardsStorage(eventID string, graphUUID uuid.UUID) error
// 	GetExistsCardsStorage(eventID string) (uuid.UUID, error)
// }

// func New(log *slog.Logger, db *database.Database) CardsStorage {
// 	return &cardsStorage{
// 	  log:         log,
// 	  db:          db,
// 	  loggerLevel: slog.String(logger.LoggerLevelArg, chartsFilesLogLevel),
// 	}
// }
