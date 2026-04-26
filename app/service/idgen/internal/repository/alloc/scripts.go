package alloc

const setSeqScript = `
local key = KEYS[1]
local lockValue = ARGV[1]
local dataSecond = ARGV[2]
local curr_seq = tonumber(ARGV[3])
local last_seq = tonumber(ARGV[4])
local mallocTime = ARGV[5]
if redis.call("EXISTS", key) == 0 then
	redis.call("HSET", key, "CURR", curr_seq, "LAST", last_seq, "TIME", mallocTime)
	redis.call("EXPIRE", key, dataSecond)
	return 1
end
if redis.call("HGET", key, "LOCK") ~= lockValue then
	return 2
end
redis.call("HDEL", key, "LOCK")
redis.call("HSET", key, "CURR", curr_seq, "LAST", last_seq, "TIME", mallocTime)
redis.call("EXPIRE", key, dataSecond)
return 0
`

const mallocScript = `
local key = KEYS[1]
local size = tonumber(ARGV[1])
local lockSecond = ARGV[2]
local dataSecond = ARGV[3]
local mallocTime = ARGV[4]
local result = {}
if redis.call("EXISTS", key) == 0 then
	local lockValue = math.random(0, 999999999)
	redis.call("HSET", key, "LOCK", lockValue)
	redis.call("EXPIRE", key, lockSecond)
	table.insert(result, 1)
	table.insert(result, lockValue)
	table.insert(result, mallocTime)
	return result
end
if redis.call("HEXISTS", key, "LOCK") == 1 then
	table.insert(result, 2)
	return result
end
local curr_seq = tonumber(redis.call("HGET", key, "CURR"))
local last_seq = tonumber(redis.call("HGET", key, "LAST"))
if size == 0 then
	redis.call("EXPIRE", key, dataSecond)
	table.insert(result, 0)
	table.insert(result, curr_seq)
	table.insert(result, last_seq)
	local setTime = redis.call("HGET", key, "TIME")
	if setTime then
		table.insert(result, setTime)
	else
		table.insert(result, 0)
	end
	return result
end
local max_seq = curr_seq + size
if max_seq > last_seq then
	local lockValue = math.random(0, 999999999)
	redis.call("HSET", key, "LOCK", lockValue)
	redis.call("HSET", key, "CURR", last_seq)
	redis.call("HSET", key, "TIME", mallocTime)
	redis.call("EXPIRE", key, lockSecond)
	table.insert(result, 3)
	table.insert(result, curr_seq)
	table.insert(result, last_seq)
	table.insert(result, lockValue)
	table.insert(result, mallocTime)
	return result
end
redis.call("HSET", key, "CURR", max_seq)
redis.call("HSET", key, "TIME", mallocTime)
redis.call("EXPIRE", key, dataSecond)
table.insert(result, 0)
table.insert(result, curr_seq)
table.insert(result, last_seq)
table.insert(result, mallocTime)
return result
`
