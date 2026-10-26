local player = {}

local health = 3
local tix = 0

local moveToggle = false -- true = no move

local gunCd = .5
local gunTracker = true --true = shootable; false = not shootable

function player.toggleMovement(t)
    moveToggle = t
end

function player.getToggle()
    return moveToggle
end

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