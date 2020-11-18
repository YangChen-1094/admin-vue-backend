package models

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
	"my_gin/pkg/global"
	"my_gin/pkg/logger"
	"strconv"
	"sync"
	"time"
)

type MemSessionData struct {
	Id     string
	Data   map[string]interface{}
	Expire int64
	Lock   sync.RWMutex
}

func NewMemSessionData(sessionId string) MemSessionData {
	return MemSessionData{
		Id:   sessionId,
		Data: make(map[string]interface{}, 8),
	}
}

func (memSession *MemSessionData) Get(key string) (val interface{}, err error) {
	memSession.Lock.RLock()
	defer memSession.Lock.RUnlock()
	val, ok := memSession.Data[key]
	if !ok {
		err = fmt.Errorf("[memSession] invalid key")
	}
	return
}

func (memSession *MemSessionData) GetId() string {
	return memSession.Id
}

func (memSession *MemSessionData) Set(key string, val interface{}) {
	memSession.Lock.Lock()
	defer memSession.Lock.Unlock()
	memSession.Data[key] = val
}

//
func (memSession *MemSessionData) UpdateExpire(timestamp int64) {
	memSession.Lock.Lock()
	defer memSession.Lock.Unlock()
	sessionId := memSession.GetId()
	//取值
	thisSessionData := RedisMgr.SessionData[sessionId]
	thisSessionData.Expire = timestamp                //赋值
	RedisMgr.SessionData[sessionId] = thisSessionData //更新
}

func (memSession *MemSessionData) Save() {
	memSession.Lock.Lock()
	defer memSession.Lock.Unlock()
	modelKey := NewRedisKey()
	key := modelKey.AdminSessionKey(memSession.GetId())
	value, err := json.Marshal(memSession.Data)
	if err != nil {
		logger.Info("redisSession", "[Save] redis 序列化sessiondata失败")
		return
	}
	RedisMgr.rdsClient.Set(key, value, time.Duration(global.WEB_ADMINS_LOGIN_EXPIRE)*time.Second)
}

func (memSession *MemSessionData) Del(key string) {
	memSession.Lock.Lock()
	defer memSession.Lock.Unlock()
	delete(memSession.Data, key)
}

type RedisManager struct {
	SessionData map[string]MemSessionData
	Lock        sync.RWMutex
	rdsClient   *redis.Client
}

var (
	RedisMgr *RedisManager
)

func NewRedisManager() *RedisManager {
	return &RedisManager{
		SessionData: make(map[string]MemSessionData, 2048),
	}
}

func (rds *RedisManager) Init(addr string, options ...string) {
	var (
		passwd string
		Db     string
	)
	if len(options) == 1 {
		passwd = options[0]
	} else if len(options) == 1 {
		passwd = options[0]
		Db = options[1]
	}
	dbVal, err := strconv.Atoi(Db)
	if err != nil {
		dbVal = 0
	}

	param := &redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       dbVal,
	}
	rds.rdsClient = redis.NewClient(param)
	_, err = rds.rdsClient.Ping().Result()
	if err != nil {
		logger.Info("redisSession", "[Init] 初始化redis失败")
		panic(err)
	}
}

//从redis加载session到内存中
func (rds *RedisManager) LoadRedisData(sessionId string) (err error) {
	modelKey := NewRedisKey()
	key := modelKey.AdminSessionKey(sessionId)
	val, err := rds.rdsClient.Get(key).Result()
	if err != nil {
		return
	}
	sessionData := &MemSessionData{}
	err = json.Unmarshal([]byte(val), sessionData)
	if err != nil {
		logger.Info("redisSession", "[LoadRedisData] json反序列化失败，sessionId：", sessionId, "val=", val)
		return
	}

	rds.Lock.Lock()
	defer rds.Lock.Unlock()
	rds.SessionData[sessionId] = *sessionData
	return
}

//获取session
func (rds *RedisManager) GetSessionData(sessionId string) (sd MemSessionData, err error) {
	rds.Lock.RLock()
	defer rds.Lock.RUnlock()
	sd, ok := rds.SessionData[sessionId]
	if !ok { //首次进程初始化
		err = rds.LoadRedisData(sessionId)
		if err != nil {
			return sd, err
		}
	}
	rds.Lock.RLock()
	defer rds.Lock.RUnlock()
	sd, ok = rds.SessionData[sessionId]
	if !ok {
		err = fmt.Errorf("redisSession, [GetSessionData] sessionId没有对应的数据")
		return
	}
	return
}

//创建session
func (rds *RedisManager) CreateSessionData() (sd MemSessionData) {
	sessionIdObj := uuid.NewV4()
	sd = NewMemSessionData(sessionIdObj.String())
	rds.SessionData[sd.GetId()] = sd
	return
}
