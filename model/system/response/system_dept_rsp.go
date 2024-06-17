package response

import "github.com/eryajf/xirang/model/system"

type DeptListRsp struct {
	Total int64         `json:"total"`
	Depts []system.Dept `json:"depts"`
}

// type Guser struct {
// 	UserId       int64  `json:"userId"`
// 	UserName     string `json:"userName"`
// 	NickName     string `json:"nickName"`
// 	Mail         string `json:"mail"`
// 	JobNumber    string `json:"jobNumber"`
// 	Mobile       string `json:"mobile"`
// 	Introduction string `json:"introduction"`
// }

// type GroupUsers struct {
// 	GroupId     int64   `json:"groupId"`
// 	GroupName   string  `json:"groupName"`
// 	GroupRemark string  `json:"groupRemark"`
// 	UserList    []Guser `json:"userList"`
// }
