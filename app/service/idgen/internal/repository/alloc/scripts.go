// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package alloc

// Lua scripts for the Redis segment cache.
//
// Both scripts operate on a single Redis hash key whose fields are:
//
//	CURR    - next sequence to be allocated (int64)
//	LAST    - exclusive end of the cached segment (int64)
//	TIME    - last successful malloc/setSeq timestamp in millis (int64)
//	LOCK    - owner string of the in-flight replenish round, "" or absent if free
//	LOCK_AT - millis when LOCK was taken (used for logical lock expiration)
//
// Lock TTL is implemented logically (now - LOCK_AT < lockMillis) so the data
// key's own EXPIRE never has to be shortened to enforce lock expiration. The
// data key's TTL is always refreshed to dataSecond on every successful op.

// mallocScript reserves up to `size` ids from the cached segment.
//
// KEYS:
//
//	[1] data key
//
// ARGV:
//
//	[1] size        - number of ids to reserve; 0 means "peek current curr/last"
//	[2] lockMillis  - logical lock TTL in millis
//	[3] dataSecond  - data key TTL in seconds
//	[4] nowMillis   - caller-supplied wall clock millis
//	[5] owner       - caller-generated unique id (string)
//
// Return values are arrays whose first element is the state code:
//
//	{0, currSeq, lastSeq, time}        - Success: caller may use [currSeq, currSeq+size)
//	{1, owner, time}                   - Miss:    cache is empty; if size>0 the
//	                                              lock is acquired and the caller
//	                                              must replenish from store and
//	                                              call setSeq with this owner.
//	                                              For a size==0 peek the lock is
//	                                              NOT taken and owner is "" so the
//	                                              caller can simply read-through.
//	{2}                                - Locked:  another caller holds the lock
//	{3, currSeq, lastSeq, owner, time} - Exceed:  segment exhausted, lock acquired;
//	                                              caller may use [currSeq, lastSeq)
//	                                              and must replenish for the rest
//
// `time` is monotonically non-decreasing within a key: if nowMillis is older
// than the previously stored TIME, the previously stored TIME is returned and
// kept as the canonical clock.
const mallocScript = `
local size = tonumber(ARGV[1])
local lockMillis = tonumber(ARGV[2])
local dataSecond = tonumber(ARGV[3])
local nowMillis = tonumber(ARGV[4])
local owner = ARGV[5]

local key = KEYS[1]

local lockOwner = redis.call("HGET", key, "LOCK")
if lockOwner and lockOwner ~= "" and lockOwner ~= false then
    local lockAt = tonumber(redis.call("HGET", key, "LOCK_AT")) or 0
    if nowMillis - lockAt < lockMillis then
        return {2}
    end
    redis.call("HDEL", key, "LOCK", "LOCK_AT")
end

local prev = tonumber(redis.call("HGET", key, "TIME")) or 0
if nowMillis < prev then
    nowMillis = prev
end

if redis.call("HEXISTS", key, "CURR") == 0 then
    -- Read-only peek on a cold key must not mutate the cache: returning a
    -- lock-less Miss lets the Go caller fall through to the store without
    -- leaving an orphaned LOCK that would block a concurrent real Malloc.
    if size == 0 then
        return {1, "", nowMillis}
    end
    redis.call("HSET", key, "LOCK", owner, "LOCK_AT", nowMillis)
    redis.call("EXPIRE", key, dataSecond)
    return {1, owner, nowMillis}
end

local curr = tonumber(redis.call("HGET", key, "CURR"))
local last = tonumber(redis.call("HGET", key, "LAST"))

if size == 0 then
    redis.call("EXPIRE", key, dataSecond)
    return {0, curr, last, nowMillis}
end

local nextCurr = curr + size
if nextCurr > last then
    redis.call("HSET", key, "CURR", last, "TIME", nowMillis, "LOCK", owner, "LOCK_AT", nowMillis)
    redis.call("EXPIRE", key, dataSecond)
    return {3, curr, last, owner, nowMillis}
end

redis.call("HSET", key, "CURR", nextCurr, "TIME", nowMillis)
redis.call("EXPIRE", key, dataSecond)
return {0, curr, last, nowMillis}
`

// setSeqScript commits a freshly fetched segment and releases the lock.
//
// It must be called with the same owner that mallocScript returned in the
// preceding Miss/Exceed result. If LOCK does not match, the script aborts
// without touching CURR/LAST so a stale caller cannot overwrite a segment
// that another owner has already committed.
//
// KEYS:
//
//	[1] data key
//
// ARGV:
//
//	[1] owner       - must equal the current LOCK field
//	[2] dataSecond  - data key TTL in seconds
//	[3] currSeq     - new CURR value
//	[4] lastSeq     - new LAST value (must be >= currSeq)
//	[5] nowMillis   - caller-supplied wall clock millis
//
// Return values:
//
//	0 - Success
//	1 - LockLost: lock was missing or held by another owner; caller must
//	    treat the freshly fetched segment as wasted (gap on the producer side)
const setSeqScript = `
local owner = ARGV[1]
local dataSecond = tonumber(ARGV[2])
local currSeq = tonumber(ARGV[3])
local lastSeq = tonumber(ARGV[4])
local nowMillis = tonumber(ARGV[5])

local key = KEYS[1]

local lockOwner = redis.call("HGET", key, "LOCK")
if (not lockOwner) or lockOwner == false or lockOwner ~= owner then
    return 1
end

local prev = tonumber(redis.call("HGET", key, "TIME")) or 0
if nowMillis < prev then
    nowMillis = prev
end

redis.call("HDEL", key, "LOCK", "LOCK_AT")
redis.call("HSET", key, "CURR", currSeq, "LAST", lastSeq, "TIME", nowMillis)
redis.call("EXPIRE", key, dataSecond)
return 0
`
