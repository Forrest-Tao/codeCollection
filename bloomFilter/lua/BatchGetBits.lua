local bloomKey = KEYS[1]
local bitsCnt = ARGS[1]

for i=1,bitsCnt,1 do
	local offset = ARGS[i+1]
	local reply = redis.call('getbit',bloomKey,offset)
	if (not reply) then
		error('Fatal')
		return 0
	end
	if (reply ==0) then
			return 0
	end
end
return 1