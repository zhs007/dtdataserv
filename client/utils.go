package dtdataclient

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/zhs007/dtdataserv/proto"
	"github.com/zhs007/jarviscore/proto"
)

// DTDataCtrlType - dtdata ctrl type
const DTDataCtrlType = "dtdata"

// BuildCtrlInfoForGameDayReport - build ctrlinfo for dtdata game day report
func BuildCtrlInfoForGameDayReport(env string, daytime string, currency string,
	scaleMoney int) (*jarviscorepb.CtrlInfo, error) {

	csd := &dtdatapb.DTDataServCtrlInfo{
		Type:       dtdatapb.DTDataType_DTDT_GAMEDAYREPORT,
		StartTime:  daytime,
		EndTime:    daytime,
		Env:        env,
		Currency:   currency,
		ScaleMoney: int32(scaleMoney),
	}

	dat, err := ptypes.MarshalAny(csd)
	if err != nil {
		return nil, err
	}

	ci := &jarviscorepb.CtrlInfo{
		CtrlType: DTDataCtrlType,
		Dat:      dat,
	}

	return ci, nil
}

// GetReplyFromCtrlResult - get DTDataServCtrlReply from CtrlResult
func GetReplyFromCtrlResult(cr *jarviscorepb.CtrlResult) (*dtdatapb.DTDataServCtrlReply, error) {

	reply := &dtdatapb.DTDataServCtrlReply{}

	err := ptypes.UnmarshalAny(cr.Dat, reply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}
