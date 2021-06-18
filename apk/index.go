package apk

import "context"

//CreateIndex 创建索引
func (c *Client) IndexCreate(index, mapping string) error {

	var (
		err   error
		exist bool
		ctx   = context.Background()
	)
	//判断索引是否存在
	if exist, err = c.IndexExists(index).Do(ctx); nil != err {
		return err
	}
	if exist {
		//索引存在
		return nil
	}
	if _, err = c.CreateIndex(index).BodyString(mapping).Do(ctx); nil != err {
		return err
	}
	return nil
}

//IndexDelete 删除索引
func (c Client) IndexDelete(index string) error {

	var (
		err error
		ctx = context.Background()
	)
	if _, err = c.DeleteIndex(index).Do(ctx); nil != err {
		return err
	}
	return nil
}

//Save 向索引加入数据
func (c Client) Save(id, index string, data interface{}) error {
	var (
		err error
		ctx = context.Background()
	)
	if _, err = c.Index().Index(index).Id(id).BodyJson(data).Do(ctx); nil != err {
		return err
	}

	return nil
}
