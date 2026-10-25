local player = {}

local health = 3

local gunCd = 1
local gunTracker = true --true = shootable; false = not shootable

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

return player