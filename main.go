package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"golang-elasticsearch/apk"
	"reflect"
)

type Good struct {
	Id        int    `json:"id"`
	Brand     string `json:"brand"`
	BrandId   int    `json:"brand_id"`
	GoodsName string `json:"goods_name"`
}

func main() {
	var (
		err     error
		ctx     = context.Background()
		client  *apk.Client
		result  []Good
		mapping = `{
						"mappings": {
							"properties": {
								"id": {
									"type": "long"
								},
								"goods_name": {
									"type": "text"
								},
								"brand_id": {
									"type": "long"
								},
								"brand": {
									"type": "keyword"
									//"type": "text"
								}
							}
						}
					}`
		index = "cytest"
	)
	if client, err = apk.InitClient("http://192.168.13.199:9200"); nil != err {
		return
	}

	/*	if err = client.IndexDelete(index); nil != err {
		return
	}*/

	if err = client.IndexCreate(index, mapping); nil != err {
		return
	}

	//插入2条数据
	if err = client.Save("1", index, map[string]interface{}{

		"id":         123,
		"goods_name": "测试商品1",
		"brand_id":   123,
		"brand":      "测试品牌1",
	}); nil != err {
		return
	}
	if err = client.Save("2", index, map[string]interface{}{

		"id":         456,
		"goods_name": "原味瓜子",
		"brand_id":   456,
		"brand":      "洽洽瓜子",
	}); nil != err {
		return
	}
	//查数据

	//布尔查询
	boolQuery := client.BoolQuery()
	//过滤
	/**
	范围查询
	elastic.NewRangeQuery("id").Gt(0)  id>0
	*/

	//boolQuery.Filter(elastic.NewRangeQuery("id").Gt(123))

	//必须满足
	//boolQuery.Must(elastic.NewMatchQuery("id", 123))

	//MatchPhraseQuery分词 keyword类型不会被分词 match的需要跟keyword的完全匹配可以。
	//match分词，text也分词，只要match的分词结果和text的分词结果有相同的就匹配。
	brandName := elastic.NewMatchPhraseQuery("brand", "测试")
	var queryShould = []elastic.Query{brandName}
	boolQuery.Should(queryShould...)
	boolQuery.MinimumNumberShouldMatch(0) //修改数字自行测试，

	//按照id 排序 ，分页
	r, err := client.Search().Index(index).Query(boolQuery).Sort("id", true).From(0).Size(10).Do(ctx)
	if nil != err {
		return
	}
	var good Good
	//这段代码 翻译自elastic.SearchResult.Each(typ reflect.Type) 方法   因为原方法遇见并不会报异常
	typ := reflect.TypeOf(good)
	if r.Hits != nil || r.Hits.Hits != nil || len(r.Hits.Hits) != 0 {
		for _, hit := range r.Hits.Hits {
			v := reflect.New(typ).Elem()
			if hit.Source == nil {
				result = append(result, v.Interface().(Good))
				continue
			}
			err := json.Unmarshal(hit.Source, v.Addr().Interface())
			if err == nil {
				result = append(result, v.Interface().(Good))
			}
		}
	}
	fmt.Printf("%+v \n", result)
	fmt.Println("done")
}
