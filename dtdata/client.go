package dtdata

import (
	"context"

	// "github.com/zhs007/jarviscore/base"
	// "go.uber.org/zap"

	"github.com/zhs007/dtdataserv/jarviscrawlercore"
	"google.golang.org/grpc"
)

type dtdataClient struct {
	servAddr string
	conn     *grpc.ClientConn
	client   jarviscrawlercore.JarvisCrawlerServiceClient
}

// newDTDataClient - new dtdataClient
func newDTDataClient(servAddr string) *dtdataClient {
	return &dtdataClient{
		servAddr: servAddr,
	}
}

// getDTData -
func (tc *dtdataClient) getDTData(ctx context.Context, mode string, startTime string, endTime string) (*jarviscrawlercore.ReplyDTData, error) {
	if tc.servAddr == "" {
		return nil, ErrNoServerAddress
	}

	if tc.conn == nil || tc.client == nil {
		conn, err := grpc.Dial(tc.servAddr, grpc.WithInsecure())
		if err != nil {
			// jarvisbase.Warn("dtdataClient.getDTData:grpc.Dial", zap.Error(err))

			return nil, err
		}

		tc.conn = conn

		tc.client = jarviscrawlercore.NewJarvisCrawlerServiceClient(conn)
	}

	reply, err := tc.client.GetDTData(ctx, &jarviscrawlercore.RequestDTData{
		Mode:      mode,
		StartTime: startTime,
		EndTime:   endTime,
	})
	if err != nil {
		// jarvisbase.Warn("dtdataClient.getDTData:GetDTData", zap.Error(err))

		// if error, close connect
		tc.conn.Close()

		tc.conn = nil
		tc.client = nil

		return nil, err
	}

	return reply, nil
}
