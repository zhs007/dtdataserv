package dtdata

import (
	"context"

	"github.com/zhs007/dtdataserv/jarviscrawlercore"
)

// DTData - DTData
type DTData struct {
	cfg    *Config
	client *dtdataClient
	db     *dtDataDB
}

// NewDTData - new dtdata
func NewDTData(filename string) (*DTData, error) {
	cfg, err := LoadConfig(filename)
	if err != nil {
		return nil, err
	}

	db, err := newDTDataDB(cfg.AnkaDB.DBPath, cfg.AnkaDB.HTTPAddr, cfg.AnkaDB.Engine)
	if err != nil {
		return nil, err
	}

	dtdata := &DTData{
		cfg:    cfg,
		db:     db,
		client: newDTDataClient(cfg.DTDataServAddr),
	}

	return dtdata, nil
}

// GetBusinessDayData -
func (dtdata *DTData) GetBusinessDayData(ctx context.Context, env string, daytime string) (*jarviscrawlercore.ReplyDTData, error) {
	data, err := dtdata.db.getBusinessDayData(ctx, env, daytime)
	if err != nil {
		return nil, err
	}

	if data == nil {
		data, err := dtdata.client.getDTData(ctx, "gamedatareport", daytime, daytime)
		if err != nil {
			return nil, err
		}

		err = dtdata.db.addBusinessDayData(ctx, env, daytime, data)
		if err != nil {
			return nil, err
		}

		return data, nil
	}

	return data, nil
}
