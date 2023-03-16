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
	msg string
}

func (b *BizAction) Name() string {
	return "biz-action"
}

func (b *BizAction) ParseConf(in []byte) (conf interface{}, err error) {
	actionConf := BizActionConf{}
	err = json.Unmarshal(in, &actionConf)
	return actionConf, err
}

func (b *BizAction) RequestFilter(conf interface{}, w http.ResponseWriter, r pkgHTTP.Request) {
	actionConf := conf.(BizActionConf)
	if len(actionConf.msg) == 0 {
		return
	}
	_, err := w.Write([]byte(actionConf.msg))
	if err != nil {
		log.Errorf("err occurred on RequestFilter", err)
	}
}
