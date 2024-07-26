local bloomKey =KEYS[1]
local bitsCnt =ARGS[1]

for i=1,bitsCnt,1 do
	local offset = ARGS[1+i]
	redis.call('setbit',bloomKey,offset,1)
end
return 1