package models

type RedisKey struct {
}

const RedisPrefix = "mhjy:"

func NewRedisKey() *RedisKey{
	return &RedisKey{}
}

//
func (key *RedisKey)AdminSessionKey(sessionId string) string{
	return RedisPrefix + "admin:" + sessionId
}