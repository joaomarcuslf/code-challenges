package challenge02

type Person struct {
	ID      string
	Friends []string
}

func unique(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

func getFriend(devs []Person, id string, depth int) []string {
	var friends []string
	for _, dev := range devs {
		if dev.ID == id {
			friends = dev.Friends
			break
		}
	}

	return friends
}

func GetSocialBubbles(devs []Person, id string) ([]string, []string, []string) {
	var friends []string
	var friendsOfFriends []string
	var friendsOfFriendsOfFriends []string

	for _, dev := range devs {
		if dev.ID == id {
			friends = dev.Friends
			break
		}
	}

	for _, friend := range friends {
		for _, dev := range devs {
			if dev.ID == friend {
				friendsOfFriends = append(friendsOfFriends, dev.Friends...)
				break
			}
		}
	}

	for _, friendOfFriend := range friendsOfFriends {
		for _, dev := range devs {
			if dev.ID == friendOfFriend {
				friendsOfFriendsOfFriends = append(friendsOfFriendsOfFriends, dev.Friends...)
				break
			}
		}
	}

	return friends, unique(friendsOfFriends), unique(friendsOfFriendsOfFriends)
}

// func mapIdsToPersons(ids []string) []*Person {
// 	var output []*Person

// 	for _, dev := range devs {
// 		for _, id := range ids {
// 			if dev.ID == id {
// 				output = append(output, &dev)
// 				break
// 			}
// 		}
// 	}

// 	return output
// }

// func recursiveExtractFriends(devs []*Person, depth int) []string {
// 	friends := []string{}

// 	if (depth == 0) {
// 		return friends
// 	}

// 	for _, dev := range devs {
// 		friends = append(friends, dev.Friends...)
// 	}
// }

// func GetSocialBubbles02(devs []Person, id string) ([]string, []string, []string) {
// 	return recursiveExtractFriends([]*Person{dev}, 3, aux)
// }
