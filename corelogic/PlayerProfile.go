package corelogic

type PlayerProfile struct {
	Name  string
	Id    string
	Email string
}

func (profile PlayerProfile) setPlayerId(id string) {
	profile.Id = id
}

func getPlayerProfile(id string) *PlayerProfile {
	Cache cache = GetCacheInstance()
	playerProfile, keyPresent := cache.profileMap[id]
	if keyPresent {
		return &playerProfile
	} else {
		playerProfile = createDummyPlayerProfile()
		cache.profileMap[playerProfile.Id] = playerProfile
		return playerProfile;
	}

}

func createDummyPlayerProfile() *PlayerProfile {
	return createPlayerProfile("dummy", "dummy@abc.edu")
}

func createPlayerProfile(name string, email string) *PlayerProfile {
	playerProfile := &PlayerProfile{Name: name, Email: email}
	playerProfile.setPlayerId(getUniquePlayerId())
	return playerProfile
}
