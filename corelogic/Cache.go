package corelogic

type Cache struct {
  var GameMap [string] Game,
  var ProfileMap [string] PlayerProfile
}

func (cache Cache) initialise() {
  cache.gameMap = make(map[string]Game)
  cache.profileMap = make(map[string]PlayerProfile)
}

var instance *Cache

func GetCacheInstance() *Cache {
	if instance == nil {
    //not thread safe
		instance = &Cache{}
    instance.initialise()
	}
	return instance
}
