package corelogic

type PlayerProfile struct {
	Name  string
	Id    string
	Email string
}

func getPlayerProfile() *PlayerProfile {
	playerProfile, keyPresent = m[key]
	if keyPresent {
		return &playerProfile;
	}
	else {
		return createDummyPlayerProfile()
	}

}

func createDummyPlayerProfile() *PlayerProfile {

}

func createPlayerProfile() *PlayerProfile {

}
