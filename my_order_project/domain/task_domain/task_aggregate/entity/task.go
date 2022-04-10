package entity

/**
 * @Author waizixi
 * @Description //TODO $
 **/

type Task struct {
	TaskId         int64
	TaskType       int64
	TaskName       int64
	TaskStatus     int64
	TaskSkuId      int64
	OrderId        int64
	TaskResourceId int64
	ExtraInfo      map[interface{}]interface{}
}

type Resource struct {
	ResourceId     int64
	ResourceUrl    string
	TaskId         int64
	ResourceStatus int64
	ExtraInfo      map[interface{}]interface{}
}
