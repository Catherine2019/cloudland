/*
Copyright <holder> All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package grpcs

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/IBM/cloudland/web/clui/model"
	"github.com/IBM/cloudland/web/sca/dbs"
)

func init() {
	Add("launch_vm", LaunchVM)
}

func LaunchVM(ctx context.Context, job *model.Job, args []string) (status string, err error) {
	//|:-COMMAND-:| launch_vm.sh '127' 'running' '3'
	db := dbs.DB()
	argn := len(args)
	if argn < 4 {
		err = fmt.Errorf("Wrong params")
		log.Println("Invalid args", err)
		return
	}
	instID, err := strconv.Atoi(args[1])
	if err != nil {
		log.Println("Invalid instance ID", err)
		return
	}
	reason := ""
	instance := &model.Instance{Model: model.Model{ID: int64(instID)}}
	err = db.Where(instance).Take(instance).Error
	if err != nil {
		log.Println("Invalid instance ID", err)
		reason = err.Error()
		return
	}
	serverStatus := args[2]
	hyperID := -1
	if serverStatus == "running" {
		hyperID, err = strconv.Atoi(args[3])
		if err != nil {
			log.Println("Invalid hyper ID", err)
			reason = err.Error()
			return
		}
	} else {
		reason = "Instance not launched for unknown reason"
	}
	err = db.Model(&instance).Updates(map[string]interface{}{
		"status": serverStatus,
		"hyper":  int32(hyperID),
		"reason": reason}).Error
	if err != nil {
		return
	}
	return
}
