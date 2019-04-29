package dtdata

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/zhs007/dtdataserv/jarviscrawlercore"
	"github.com/zhs007/dtdataserv/proto"
	"github.com/zhs007/jarviscore"
	"github.com/zhs007/jarviscore/base"
	"github.com/zhs007/jarviscore/proto"
	"go.uber.org/zap"
)

// DTData - DTData
type DTData struct {
	cfg      *Config
	client   *dtdataClient
	db       *dtDataDB
	httpserv *HTTPServer
}

// NewDTData - new dtdata
func NewDTData(filename string) (*DTData, error) {
	cfg, err := LoadConfig(filename)
	if err != nil {
		jarvisbase.Error("NewDTData", zap.Error(err))

		return nil, err
	}

	db, err := newDTDataDB(cfg.AnkaDB.DBPath, cfg.AnkaDB.HTTPAddr, cfg.AnkaDB.Engine)
	if err != nil {
		jarvisbase.Error("NewDTData:newDTDataDB", zap.Error(err))

		return nil, err
	}

	httpserv, err := newHTTPServer(cfg.HTTPAddr, db)
	if err != nil {
		jarvisbase.Error("NewDTData:newHTTPServer", zap.Error(err))

		return nil, err
	}

	dtdata := &DTData{
		cfg:      cfg,
		db:       db,
		client:   newDTDataClient(cfg.DTDataServAddr),
		httpserv: httpserv,
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
			jarvisbase.Warn("DTData.GetBusinessDayData:getDTData", zap.Error(err))

			return nil, err
		}

		err = dtdata.db.addBusinessDayData(ctx, env, daytime, data)
		if err != nil {
			jarvisbase.Warn("DTData.GetBusinessDayData:addBusinessDayData", zap.Error(err))

			return nil, err
		}

		return data, nil
	}

	return data, nil
}

// Run - jarviscore.Ctrl
func (dtdata *DTData) Run(ctx context.Context, jarvisnode jarviscore.JarvisNode, srcAddr string, msgid int64, ci *jarviscorepb.CtrlInfo) []*jarviscorepb.JarvisMsg {

	var dtdataci dtdatapb.DTDataServCtrlInfo
	err := ptypes.UnmarshalAny(ci.Dat, &dtdataci)
	if err != nil {
		jarvisbase.Warn("DTData.Run:UnmarshalAny", zap.Error(err))

		return []*jarviscorepb.JarvisMsg{
			jarviscore.NewErrorMsg(jarvisnode, srcAddr, err.Error(), msgid),
		}
	}

	hashstr := buildHashBuffer(ci.Dat.Value)

	report, err := dtdata.db.getCache(ctx, hashstr)
	if err != nil {
		jarvisbase.Warn("DTData.Run:getCache", zap.Error(err))

		return []*jarviscorepb.JarvisMsg{
			jarviscore.NewErrorMsg(jarvisnode, srcAddr, err.Error(), msgid),
		}
	}

	if report != nil {
		rc := &dtdatapb.DTDataServCtrlReply{
			Token: hashstr,
		}

		msgrc, err := jarviscore.NewCtrlResult(jarvisnode, srcAddr, msgid, rc)
		if err != nil {
			jarvisbase.Warn("DTData.Run:NewCtrlResult", zap.Error(err))

			return []*jarviscorepb.JarvisMsg{
				jarviscore.NewErrorMsg(jarvisnode, srcAddr, err.Error(), msgid),
			}
		}

		return []*jarviscorepb.JarvisMsg{
			msgrc,
		}
	}

	if dtdataci.Type == dtdatapb.DTDataType_DTDT_GAMEDAYREPORT {
		bdddat, err := dtdata.GetBusinessDayData(ctx, dtdataci.Env, dtdataci.StartTime)
		if err != nil {
			jarvisbase.Warn("DTData.Run:GetBusinessDayData", zap.Error(err))

			return []*jarviscorepb.JarvisMsg{
				jarviscore.NewErrorMsg(jarvisnode, srcAddr, err.Error(), msgid),
			}
		}

		if bdddat == nil {
			jarvisbase.Warn("DTData.Run:GetBusinessDayData", zap.Error(ErrNoBusinessDayData))

			return []*jarviscorepb.JarvisMsg{
				jarviscore.NewErrorMsg(jarvisnode, srcAddr, ErrNoBusinessDayData.Error(), msgid),
			}
		}

		curreport := countDTReportWithBusinessGameReport(bdddat, dtdataci.Currency, int(dtdataci.ScaleMoney))
		if curreport != nil {
			err = dtdata.db.setCache(ctx, hashstr, curreport)
			if err != nil {
				return []*jarviscorepb.JarvisMsg{
					jarviscore.NewErrorMsg(jarvisnode, srcAddr, err.Error(), msgid),
				}
			}

			rc := &dtdatapb.DTDataServCtrlReply{
				Token: hashstr,
			}

			msgrc, err := jarviscore.NewCtrlResult(jarvisnode, srcAddr, msgid, rc)
			if err != nil {
				jarvisbase.Warn("DTData.Run:NewCtrlResult", zap.Error(err))

				return []*jarviscorepb.JarvisMsg{
					jarviscore.NewErrorMsg(jarvisnode, srcAddr, err.Error(), msgid),
				}
			}

			return []*jarviscorepb.JarvisMsg{
				msgrc,
			}
		}

		return []*jarviscorepb.JarvisMsg{
			jarviscore.NewErrorMsg(jarvisnode, srcAddr, ErrGameDayReport.Error(), msgid),
		}
	}

	return []*jarviscorepb.JarvisMsg{
		jarviscore.NewErrorMsg(jarvisnode, srcAddr, ErrInvliadDTDataType.Error(), msgid),
	}
}

// Start - start a service
func (dtdata *DTData) Start(ctx context.Context, jarvisnode jarviscore.JarvisNode) {
	jarvisnode.RegCtrl(DTDataCtrlType, dtdata)

	dtdata.httpserv.start(ctx)
}
