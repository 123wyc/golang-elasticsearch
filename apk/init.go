package apk

import (
	"github.com/olivere/elastic/v7"
	"time"
)

//InitClient 初始化elas
func InitClient(addr string) (*Client, error) {

	c, err := elastic.NewClient(elastic.SetURL(addr), elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second))
	if nil != err {
		return nil, err
	}
	return &Client{c}, nil
}
