package mock

import (
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"math/rand"
	"online-chat/global"
	"online-chat/online_server/models/requestx"
	"time"
)

func MockUserBasic() {
	var count int64
	err := global.DB.Find(&requestx.UserBasic{}).Count(&count).Debug().Error
	if err != nil {
		zap.S().Errorf("check database error: %v", err)
		return
	}
	if count > 0 {
		fmt.Println("count: ", count)
		return
	}

	for i := 0; i < 10; i++ {
		u := requestx.UserBasic{
			Name:       uuid.New().String(),
			Password:   uuid.New().String(),
			Phone:      generatePhone(),
			Email:      generateEmail(),
			Identity:   uuid.New().String(),
			ClientIp:   generateIp(),
			ClientPort: generatePort(),
		}
		err := global.DB.Model(&requestx.UserBasic{}).Create(&u).Error
		if err != nil {
			zap.S().Errorf("mock 数据失败")
		}
		zap.S().Info("mock data success.")
	}

}

func generatePhone() string {
	var headerNums = [...]string{"139", "138", "137", "136", "135", "134", "159", "158", "157", "150", "151", "152", "188", "187", "182", "183", "184", "178", "130", "131", "132", "156", "155", "186", "185", "176", "133", "153", "189", "180", "181", "177"}
	var headerNumsLen = len(headerNums)
	header := headerNums[rand.Intn(headerNumsLen)]
	body := fmt.Sprintf("%08d", rand.Intn(99999999))
	phone := header + body
	return phone
}

func generateEmail() string {
	emailSlice := []string{"vknawqq@265.com", "opwfqqqqnja@email.com.cn", "bihgnawksufa@hotmail.com", "avevnflf@yeah.net", "hehcpmchddifqp@qq.com", "qgss@eastday.com", "gcqsikklhske@163.net", "qusgsuugfrelsum@21cn.com", "cicuf@xinhuanet", "dotptnsm@265.com", "wtklwthatvnu@citiz.com", "vbhhssmkkvd@netease.com", "fjpw@sina.com", "buqrwdmstik@citiz.com", "agqrfk@yahoo.com.cn", "emilhtsbnewv@35.com", "pehi@netease.com", "apbap@263.net", "rlpjlvsuiifogr@qq.com", "ekhckomktgku@yahoo.com.cn", "htjqanvrk@etang.com", "pht@eyou.com", "umh@citiz.com", "mispdwem@sogou@com", "vivajf@xinhuanet", "dhjwgrrgqi@email.com.cn", "uncm@56.com", "nauoqefmjhjsp@xinhuanet", "teckmbbvt@eyou.com", "kunhfi@msn.com", "omtmne@sogou@com", "ikuuqoabjwcf@email.com.cn", "cbhuriwb@netease.com", "nrt@sina.com", "pmgsiplvicstl@126.com", "galckreqkjrgith@263.net", "fbrktkbtncwrvlk@citiz.com", "rnfr@hotmail.com", "kssjnic@163.net", "dbo@sina.com", "tjmkfupe@163.com", "ugdnbrqwvtgoaak@xinhuanet", "avoutawvh@xinhuanet", "mkm@hotmail.com", "hlhelw@265.com", "cqnuej@etang.com", "rioghnbge@enet.com.cn", "rvklgarpmdg@citiz.com", "oesdjbqvltsbvwn@sina.com", "aatfmloegukr@etang.com"}
	rand.Seed(time.Now().UnixMilli())
	index := rand.Intn(len(emailSlice) - 1)
	return emailSlice[index]
}

func generateIp() string {
	rand.Seed(time.Now().Unix())
	ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
	return ip
}

func generatePort() string {
	rand.Seed(time.Now().UnixMilli())
	return fmt.Sprintf("%d", rand.Intn(10000)+10000)
}
