package models

import (
	"github.com/apache/incubator-devlake/models/common"
	"time"
)

type GitlabIssue struct {
	GitlabId        int    `gorm:"primaryKey"`
	ProjectId       int    `gorm:"index"`
	Number          int    `gorm:"index;comment:Used in API requests ex. api/repo/1/issue/<THIS_NUMBER>"`
	State           string `gorm:"type:varchar(255)"`
	Title           string
	Body            string
	Priority        string `gorm:"type:varchar(255)"`
	Type            string `gorm:"type:varchar(100)"`
	Status          string `gorm:"type:varchar(255)"`
	AssigneeId      string
	AssigneeName    string `gorm:"type:varchar(255)"`
	LeadTimeMinutes uint
	Url             string `gorm:"type:varchar(255)"`
	ClosedAt        *time.Time
	GitlabCreatedAt time.Time
	GitlabUpdatedAt time.Time `gorm:"index"`
	Severity        string    `gorm:"type:varchar(255)"`
	Component       string    `gorm:"type:varchar(255)"`
	TimeEstimate 	int64
	TotalTimeSpent 	int64
	common.NoPKModel
}
func (GitlabIssue) TableName() string {
	return "_tool_gitlab_issues"
}

