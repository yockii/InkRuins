package util

import (
	snowflake "github.com/yockii/snowflake_ext"
)

var worker *snowflake.Worker

func InitSnowflake(workerId uint64) error {
	var err error
	worker, err = snowflake.NewSnowflake(workerId)
	if err != nil {
		return err
	}
	return nil
}

func NextID() uint64 {
	if worker == nil {
		return 0
	}
	return worker.NextId()
}
