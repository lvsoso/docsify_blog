import time
import json

from flask import g, request


def get_identifiers():

    ret = ["ip:" + request.remote_addr]

    if g.user.is_authenticated():
        ret.append("user:%s"%g.user.get_id())
    
    return ret



def over_limit_multi(conn, limits=[(1, 10), (60, 120), (3600, 240)]):

    if not hasattr(conn, 'over_limit_multi_lua'):
        conn.over_limit_multi_lua = conn.register_script(over_limit_multi_lua_)

    
    return conn.over_limit_multi_lua(
        keys = get_identifiers(), args=[json.dumps(limits), time.time()])


over_limit_multi_lua_ = '''
local limits = cjson.decode(ARGV[1])
local now = tonumber(ARGV[2])
for i, limit in ipairs(limits) do
    local duration = limit[1]
 
    local bucket = ':' .. duration .. ':' .. math.floor(now / duration)
    for j, id in ipairs(KEYS) do
        local key = id .. bucket
 
        local count = redis.call('INCR', key)
        redis.call('EXPIRE', key, duration)
        if tonumber(count) > limit[2] then
            return 1
        end
    end
end
return 0
'''


