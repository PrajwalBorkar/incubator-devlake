/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tasks

import (
	"reflect"

	"github.com/apache/incubator-devlake/models/domainlayer"
	"github.com/apache/incubator-devlake/models/domainlayer/didgen"
	"github.com/apache/incubator-devlake/models/domainlayer/ticket"
	"github.com/apache/incubator-devlake/plugins/core"
	"github.com/apache/incubator-devlake/plugins/helper"
	"github.com/apache/incubator-devlake/plugins/jira/models"
)

func ConvertBoard(taskCtx core.SubTaskContext) error {
	data := taskCtx.GetData().(*JiraTaskData)
	logger := taskCtx.GetLogger()
	db := taskCtx.GetDb()
	logger.Info("collect board:%d", data.Options.BoardId)
	idGen := didgen.NewDomainIdGenerator(&models.JiraBoard{})
	cursor, err := db.Model(&models.JiraBoard{}).Where("connection_id = ? AND board_id = ?", data.Connection.ID, data.Options.BoardId).Rows()
	if err != nil {
		return err
	}
	defer cursor.Close()
	converter, err := helper.NewDataConverter(helper.DataConverterArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: JiraApiParams{
				ConnectionId: data.Connection.ID,
				BoardId:      data.Options.BoardId,
			},
			Table: RAW_BOARD_TABLE,
		},
		InputRowType: reflect.TypeOf(models.JiraBoard{}),
		Input:        cursor,
		Convert: func(inputRow interface{}) ([]interface{}, error) {
			board := inputRow.(*models.JiraBoard)
			domainBoard := &ticket.Board{
				DomainEntity: domainlayer.DomainEntity{Id: idGen.Generate(data.Connection.ID, data.Options.BoardId)},
				Name:         board.Name,
				Url:          board.Self,
			}
			return []interface{}{
				domainBoard,
			}, nil
		},
	})
	if err != nil {
		return err
	}

	return converter.Execute()
}
