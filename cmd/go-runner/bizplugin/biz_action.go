package bizplugin

import (
	"encoding/json"
	pkgHTTP "github.com/apache/apisix-go-plugin-runner/pkg/http"
	"github.com/apache/apisix-go-plugin-runner/pkg/log"
	"github.com/apache/apisix-go-plugin-runner/pkg/plugin"
	"net/http"
)

func init() {
	err := plugin.RegisterPlugin(&BizAction{})
	if err != nil {
		log.Errorf("err occurred on init", err)
	}
}

type BizAction struct {
	plugin.DefaultPlugin
}

type BizActionConf struct {
	Msg string `json:"msg"`
}

func (b *BizAction) Name() string {
	return "biz-action"
}

func (b *BizAction) ParseConf(in []byte) (conf interface{}, err error) {
	actionConf := BizActionConf{}
	log.Infof("conf: %s", string(in))
	err = json.Unmarshal(in, &actionConf)
	return actionConf, err
}

func (b *BizAction) RequestFilter(conf interface{}, w http.ResponseWriter, r pkgHTTP.Request) {
	actionConf := conf.(BizActionConf)
	marshal, err := json.Marshal(actionConf)
	if err != nil {
		log.Errorf("json Marshal failed", err)
	}
	log.Infof("info: %s", marshal)
	msg := actionConf.Msg
	if len(msg) == 0 {
		return
	}
	w.Header().Add("x-biz-action", "filtered")
	_, err = w.Write([]byte(msg))
	if err != nil {
		log.Errorf("err occurred on RequestFilter", err)
	}
}
