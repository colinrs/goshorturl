package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/colinrs/goshorturl/pkg/code"
	"github.com/colinrs/goshorturl/pkg/gosafe"
	"github.com/colinrs/goshorturl/pkg/httpc"
	"github.com/zeromicro/go-zero/core/logx"
)

type IdGenClient interface {
	NextId(ctx context.Context) (int64, error)
}

type idGenClient struct {
	host       string
	bizTagName string
	ids        chan int64
	httpClient httpc.Client
	stat       *stat
}

func NewIdGenClient(options ...Option) IdGenClient {
	idGen := &idGenClient{}
	o := &Options{
		host: "http://127.0.0.1:8888",
	}
	for _, option := range options {
		option(o)
	}
	idGen.host = o.host
	idGen.bizTagName = o.bizTagName
	idGen.ids = make(chan int64, 1000)
	idGen.httpClient = httpc.NewClient(idGen.host)
	idGen.stat = newStat()
	logx.Must(idGen.sendToChannel())
	logx.Must(idGen.start())
	return idGen
}

func (i *idGenClient) NextId(ctx context.Context) (int64, error) {
	for {
		select {
		case nextId := <-i.ids:
			return nextId, nil
		case <-ctx.Done():
			return 0, ctx.Err()
		}
	}
}

func (i *idGenClient) start() error {
	gosafe.GoSafe(context.Background(), func() {
		i.sendToChannelLoop()
	})
	gosafe.GoSafe(context.Background(), func() {
		i.logStat()
	})
	return nil
}

func (i *idGenClient) sendToChannelLoop() {
	for {
		err := i.sendToChannel()
		if err != nil {
			logx.Errorf("goleaf idGenClient sendToChannel error: %v", err)
		}
	}
}

func (i *idGenClient) sendToChannel() error {
	nextIds, err := i.getNextIds()
	if err != nil {
		logx.Errorf("goleaf idGenClient getNextIds error: %v", err)
		return err
	}
	i.stat.lastIdUpdate.Store(time.Now().Unix())
	for _, nextId := range nextIds {
		i.ids <- nextId
	}
	i.stat.lastIdUpdate.Store(time.Now().Unix())
	return nil
}

func (i *idGenClient) logStat() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		<-ticker.C
		logx.Debugf("goleaf idGenClient lens: %d, last id update: %d",
			len(i.ids), i.stat.lastIdUpdate.Load())
	}
}

func (i *idGenClient) getNextIds() ([]int64, error) {
	if i.bizTagName == "" {
		return i.getFromSnowflake()
	}
	return i.getFromSegment()
}

func (i *idGenClient) getFromSegment() ([]int64, error) {
	resp, err := i.httpClient.Get(context.Background(), fmt.Sprintf("/api/v1/segment/get?biz_tag=%s", i.bizTagName))
	if err != nil {
		return nil, err
	}
	segmentResponse := &SegmentResponse{}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, segmentResponse); err != nil {
		return nil, err
	}
	if segmentResponse.Code != 0 {
		return nil, code.NewErr(code.WithCode(segmentResponse.Code), code.WithMsg(segmentResponse.Msg))
	}
	logx.Debugf("getFromSegment lens: %d, minId: %d, maxId: %d",
		segmentResponse.Data.Step, segmentResponse.Data.MinID, segmentResponse.Data.MaxID)
	ids := make([]int64, 0, segmentResponse.Data.Step)
	for i := segmentResponse.Data.MinID; i < segmentResponse.Data.MaxID; i++ {
		ids = append(ids, i)
	}
	return ids, nil
}

func (i *idGenClient) getFromSnowflake() ([]int64, error) {
	resp, err := i.httpClient.Get(context.Background(), "/api/v1/snowflake/get?step=100")
	if err != nil {
		return nil, err
	}
	snowflakeResponse := &SnowflakeResponse{}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, snowflakeResponse); err != nil {
		return nil, err
	}
	if snowflakeResponse.Code != 0 {
		return nil, code.NewErr(code.WithCode(snowflakeResponse.Code), code.WithMsg(snowflakeResponse.Msg))
	}
	logx.Debugf("getFromSnowflake lens: %d,start:%d,end:%d", snowflakeResponse.Data.Total,
		snowflakeResponse.Data.List[0], snowflakeResponse.Data.List[len(snowflakeResponse.Data.List)-1])
	return snowflakeResponse.Data.List, nil
}
