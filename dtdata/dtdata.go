package dtdata

import (
	"context"

	// "github.com/zhs007/jarviscore"

	"github.com/zhs007/ankadb"
	"github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"
)

// DTData - DTData database
type DTData struct {
	ankaDB ankadb.AnkaDB
}

// NewDTData - new dtdata db
func NewDTData(dbpath string, httpAddr string, engine string) (*DTData, error) {
	cfg := ankadb.NewConfig()

	cfg.AddrHTTP = httpAddr
	cfg.PathDBRoot = dbpath
	cfg.ListDB = append(cfg.ListDB, ankadb.DBConfig{
		Name:   DTDataDBName,
		Engine: engine,
		PathDB: DTDataDBName,
	})

	ankaDB, err := ankadb.NewAnkaDB(cfg, nil)
	if ankaDB == nil {
		jarvisbase.Error("NewDTData", zap.Error(err))

		return nil, err
	}

	jarvisbase.Info("NewDTData", zap.String("dbpath", dbpath),
		zap.String("httpAddr", httpAddr), zap.String("engine", engine))

	db := &DTData{
		ankaDB: ankaDB,
	}

	return db, err
}

func (db *DTData) AddBusinessDayData(ctx context.Context) (bool, error) {
	// has, err := db.getArticle(ctx, uid, website, url)
	// if err != nil {
	// 	return false, err
	// }

	// if has {
	// 	return false, nil
	// }

	// err = db.ankaDB.Set(ctx, CrawlerDBName, jarviscore.AppendString(uid, ":", website, ":", url), []byte(url))
	// if err != nil {
	// 	return false, err
	// }

	return true, nil
}

func (db *DTData) getArticle(ctx context.Context, uid string, website string, url string) (bool, error) {
	// _, err := db.ankaDB.Get(ctx, CrawlerDBName, jarviscore.AppendString(uid, ":", website, ":", url))
	// if err != nil {
	// 	if err == ankadb.ErrNotFoundKey {
	// 		return false, nil
	// 	}

	// 	return false, err
	// }

	return true, nil
}
