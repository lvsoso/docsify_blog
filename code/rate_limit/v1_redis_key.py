import time

from flask import g, request


def get_identifiers():

    ret = ["ip:" + request.remote_addr]

    if g.user.is_authenticated():
        ret.append("user:%s"%g.user.get_id())
    
    return ret


def over_limit(conn, duration=3600, limit=240):

    pipe = conn.pipeline(transaction=True)

    bucket = ":%i:%i"%(duration, time.time())

    for id in get_identifiers():
        key = id + bucket

        pipe.incr(key)

        pipe.expire(key, duration)

        if pipe.execute()[0] > limit :
            return True
        
    return False


def over_limit_multi(conn, limits=[(1, 10), (60, 120), (3600, 240)]):

    for duration, limit in limits:

        if over_limit(conn, duration, limit):

            return True
    
    return False
