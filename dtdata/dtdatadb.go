package dtdata

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/ankadb"
	"github.com/zhs007/dtdataserv/jarviscrawlercore"
	"github.com/zhs007/dtdataserv/proto"
	// "github.com/zhs007/jarviscore"
	"github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"
)

// dtDataDB - DTData database
type dtDataDB struct {
	ankaDB ankadb.AnkaDB
}

// newDTDataDB - new dtdata db
func newDTDataDB(dbpath string, httpAddr string, engine string) (*dtDataDB, error) {
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
		jarvisbase.Error("newDTDataDB", zap.Error(err))

		return nil, err
	}

	jarvisbase.Info("newDTDataDB", zap.String("dbpath", dbpath),
		zap.String("httpAddr", httpAddr), zap.String("engine", engine))

	db := &dtDataDB{
		ankaDB: ankaDB,
	}

	return db, nil
}

// addBusinessDayData -
func (db *dtDataDB) addBusinessDayData(ctx context.Context, env string, daytime string, dtdata *jarviscrawlercore.ReplyDTData) error {
	curkey := makeBusinessDayDataKey(env, daytime)

	buf, err := proto.Marshal(dtdata)
	if err != nil {
		jarvisbase.Warn("dtDataDB.addBusinessDayData:Marshal", zap.Error(err))

		return err
	}

	err = db.ankaDB.Set(ctx, DTDataDBName, curkey, buf)
	if err != nil {
		jarvisbase.Warn("dtDataDB.addBusinessDayData:Set", zap.Error(err))

		return err
	}

	return nil
}

// getBusinessDayData -
func (db *dtDataDB) getBusinessDayData(ctx context.Context, env string, daytime string) (*jarviscrawlercore.ReplyDTData, error) {
	curkey := makeBusinessDayDataKey(env, daytime)

	buf, err := db.ankaDB.Get(ctx, DTDataDBName, curkey)
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			return nil, nil
		}

		jarvisbase.Warn("dtDataDB.getBusinessDayData:Get", zap.Error(err))

		return nil, err
	}

	data := &jarviscrawlercore.ReplyDTData{}

	err = proto.Unmarshal(buf, data)
	if err != nil {
		jarvisbase.Warn("dtDataDB.getBusinessDayData:Unmarshal", zap.Error(err))

		return nil, err
	}

	return data, nil
}

// getCache -
func (db *dtDataDB) getCache(ctx context.Context, hashstr string) (*dtdatapb.DTReport, error) {
	curkey := makeCacheKey(hashstr)

	buf, err := db.ankaDB.Get(ctx, DTDataDBName, curkey)
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			return nil, nil
		}

		jarvisbase.Warn("dtDataDB.getCache:Get", zap.Error(err))

		return nil, err
	}

	data := &dtdatapb.DTReport{}

	err = proto.Unmarshal(buf, data)
	if err != nil {
		jarvisbase.Warn("dtDataDB.getCache:Unmarshal", zap.Error(err))

		return nil, err
	}

	return data, nil
}

// setCache -
func (db *dtDataDB) setCache(ctx context.Context, hashstr string, report *dtdatapb.DTReport) error {
	curkey := makeCacheKey(hashstr)

	buf, err := proto.Marshal(report)
	if err != nil {
		jarvisbase.Warn("dtDataDB.setCache:Marshal", zap.Error(err))

		return err
	}

	err = db.ankaDB.Set(ctx, DTDataDBName, curkey, buf)
	if err != nil {
		jarvisbase.Warn("dtDataDB.setCache:Set", zap.Error(err))

		return err
	}

	return nil
}
