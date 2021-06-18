package apk

import "github.com/olivere/elastic/v7"

func (c Client) BoolQuery() *elastic.BoolQuery {
	return elastic.NewBoolQuery()
}
