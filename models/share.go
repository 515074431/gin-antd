package models

import (
	"database/sql"
	"errors"
	"github.com/jinzhu/gorm"
	"path/filepath"
	"time"
)

type Share struct {
	gorm.Model
	ShareType sql.NullInt64
	ShareRoot string	`gorm:"type:varchar(255)"`//分享到的根
	ShareWith string	`gorm:"type:varchar(50)"`
	Password string	`gorm:"type:varchar(32)"`
	UidOwner string `gorm:"type:varchar(50)"`
	UidInitiator string
	Parent int
	ItemType string
	ItemSource string
	ItemTarget string
	FileSource string
	FileTarget string
	Permissions int
	Accepted int
	Expiration *time.Time
	Token string
	MailSend int
	ShareName string	`gorm:"type:varchar(255)"`
	PasswordByTalk int
	Note string
	HideDownload int
	label string


}
//获取根目录下的分享文件列表
func ShareRootList(path string) (shares []Share,err error){
	//var shares [] Share
	result :=Db.Where("share_root = ?",path).Find(&shares)
	if result.Error !=nil{
		err = result.Error
	}
	return
}
//获取文件的分享信息
func ShareInfo(path string) (sh Share,err error)  {

	err = errors.New("没有分享信息")
	var shares [] Share
	result :=Db.Where("item_target IN (?)",spitList(path)).Find(&shares)
	//log.Println("shares array:",shares)
	if result.Error !=nil{
		return
	}
	if len(shares) > 0{//没有分享信息
		 sh = shares[len(shares)-1]
		 return sh,nil
	}
	return
}
//处理路径为数组
func spitList(path string) (list []string) {
	if path == "" {
		return
	}
	for i := 0; i < len(path); i++ {
		if path[i]  == '/' || path[i]  ==  '\\'  {
			list = append(list,path[:i])
		}
	}
	if path[len(path)-1] != filepath.Separator {
		list = append(list,path)
	}
	return
}