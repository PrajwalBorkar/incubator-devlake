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

package archived

import (
	"time"
)

type PullRequest struct {
	DomainEntity
	BaseRepoId     string `gorm:"index"`
	HeadRepoId     string `gorm:"index"`
	Status         string `gorm:"type:varchar(100);comment:open/closed or other"`
	Number         int
	Title          string
	Description    string
	Url            string `gorm:"type:varchar(255)"`
	AuthorName     string `gorm:"type:varchar(100)"`
	AuthorId       string `gorm:"type:varchar(100)"`
	ParentPrId     string `gorm:"index;type:varchar(100)"`
	Key            int
	CreatedDate    time.Time
	MergedDate     *time.Time
	ClosedDate     *time.Time
	Type           string `gorm:"type:varchar(100)"`
	Component      string `gorm:"type:varchar(100)"`
	MergeCommitSha string `gorm:"type:varchar(40)"`
	HeadRef        string `gorm:"type:varchar(255)"`
	BaseRef        string `gorm:"type:varchar(255)"`
	BaseCommitSha  string `gorm:"type:varchar(40)"`
	HeadCommitSha  string `gorm:"type:varchar(40)"`
}

type PullRequestCommit struct {
	CommitSha     string `gorm:"primaryKey;type:varchar(40)"`
	PullRequestId string `json:"id" gorm:"primaryKey;type:varchar(255);comment:This key is generated based on details from the original plugin"` // format: <Plugin>:<Entity>:<PK0>:<PK1>
	NoPKModel
}

type PullRequestIssue struct {
	PullRequestId     string `json:"id" gorm:"primaryKey;type:varchar(255);comment:This key is generated based on details from the original plugin"` // format: <Plugin>:<Entity>:<PK0>:<PK1>
	IssueId           string `gorm:"primaryKey;type:varchar(255)"`
	PullRequestNumber int
	IssueNumber       int
	NoPKModel
}

// Please note that Issue Labels can also apply to Pull Requests.
// Pull Requests are considered Issues in GitHub.

type PullRequestLabel struct {
	PullRequestId string `json:"id" gorm:"primaryKey;type:varchar(255);comment:This key is generated based on details from the original plugin"` // format: <Plugin>:<Entity>:<PK0>:<PK1>
	LabelName     string `gorm:"primaryKey;type:varchar(255)"`
	NoPKModel
}

type PullRequestComment struct {
	DomainEntity
	PullRequestId string `gorm:"index"`
	Body          string
	UserId        string `gorm:"type:varchar(255)"`
	CreatedDate   time.Time
	CommitSha     string `gorm:"type:varchar(255)"`
	Position      int
}
