package WeChatCustomerServiceSDK

import (
	"github.com/NICEXAI/WeChatCustomerServiceSDK/util"
	"strconv"
	"strings"
	"time"
)

// MonitorSchema SDK监控数据
type MonitorSchema map[string]map[string]int

// Monitor 获取当前服务SDK调用统计数据
func (r *Client) Monitor() MonitorSchema {
	keyList := make([]string, 0)
	cursor := uint64(0)
	hasMore := true

	for hasMore {
		keys, newCursor, err := r.cache.Scan(cursor, "wechat:log:" + r.corpID + "*", 1000)
		if err != nil {
			break
		}
		if newCursor == 0 || len(keys) == 0 {
			hasMore = false
		}
		cursor = newCursor
		keyList = append(keyList, keys...)
	}

	monitorInfo := MonitorSchema{}
	for _, key := range keyList {
		options := strings.Split(key, ":")
		if len(options) != 5 {
			continue
		}
		con, err := r.cache.Get(key)
		if err != nil {
			continue
		}
		timeDic := monitorInfo[options[3]]
		if timeDic == nil {
			monitorInfo[options[3]] = map[string]int{}
		}
		count, err := strconv.Atoi(con)
		if err != nil {
			continue
		}
		monitorInfo[options[3]][options[4]] = count
	}
	return monitorInfo
}

//记录SDK调用信息
func (r *Client) recordUpdate(path string) {
	if !r.monitorOpen {
		return
	}
	path = util.ParseRoute(path)
	if path == "" {
		return
	}
	recordKey := "wechat:log:" + r.corpID + ":" + time.Now().Format("2006010215") + ":" + path[1:]
	con, err := r.cache.Get(recordKey)
	if err != nil {
		return
	}
	if con == "" {
		if err = r.cache.Set(recordKey, strconv.Itoa(1), r.monitorLogExpireTime); err != nil {
			return
		}
		return
	}
	count, err := strconv.Atoi(con)
	if err != nil {
		return
	}
	_ = r.cache.Set(recordKey, strconv.Itoa(count + 1), r.monitorLogExpireTime)
}