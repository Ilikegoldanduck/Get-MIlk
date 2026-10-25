local player = {}

local health = 3
local tix = 0


local gunCd = 1
local gunTracker = true --true = shootable; false = not shootable

function player.reset()
    health = 3
    tix = 0
    gunTracker = true
end

function player.getHealth()
    return health
end

function player.setHealth(h)
    health = h
end

function player.gunCheck()
    if gunTracker then
        gunTracker = false
        timer.delay(gunCd, false, function()
            gunTracker = true
        end)
        return true
    else 
        return false
    end
end

function player.gunCheckGui()
   return gunTracker 
end

function player.getTix()
    return tix
end

function player.addTix(v)
    tix = tix + v
end

return player