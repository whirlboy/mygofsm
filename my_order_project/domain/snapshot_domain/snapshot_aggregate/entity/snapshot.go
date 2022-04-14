package entity

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

/**
 * @Author waizixi
 * @Description //TODO $
 **/

//go:generate easytags $GOFILE json
type TaskDtoSubmitSnapshot struct {
	Id            int64
	UserId        int64
	SubmitType    int64
	SubmitCommand string
	ReqToken      string
	ExtraInfo     map[interface{}]interface{}
}

func (snapshotEntity *TaskDtoSubmitSnapshot) GenerateSnapshotToken() error {
	a, err := json.Marshal(snapshotEntity.SubmitCommand)
	if err != nil {
		return err
	}
	md5Hash := md5.New()
	md5Hash.Write(a)
	snapshotEntity.ReqToken = hex.EncodeToString(md5Hash.Sum(nil))
	return nil
}
