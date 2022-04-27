package function

import (
	"github.com/google/uuid"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	inGameMode []InGameMode
	rankLevel  []RankLevelType
)

func init() {
	inGameMode = append(inGameMode, InGameMode_TEAM_COLLECT_STAR_4_VS_4)
	inGameMode = append(inGameMode, InGameMode_TOWER)
	inGameMode = append(inGameMode, InGameMode_DEATH_MATCH)
	inGameMode = append(inGameMode, InGameMode_SOLO_SURVIVAL)
	inGameMode = append(inGameMode, InGameMode_DUAL_SURVIVAL)

	rankLevel = append(rankLevel, RankRecruit)
	rankLevel = append(rankLevel, RankPrivate)
	rankLevel = append(rankLevel, RankBronze)
	rankLevel = append(rankLevel, RankSilver)
	rankLevel = append(rankLevel, RankGolden)
	rankLevel = append(rankLevel, RankPlatinum)
	rankLevel = append(rankLevel, RankDiamond)
	rankLevel = append(rankLevel, RankMaster)
	rankLevel = append(rankLevel, RankChampion)
}

func randomSkill() int32 {
	skillPool := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 35, 36, 37}

	return skillPool[rand.Intn(len(skillPool))]
}

func randomRegion() int32 {
	skillPool := []int32{0, 1, 2, 3, 7, 9, 11, 10, 13}

	return skillPool[rand.Intn(len(skillPool))]
}

func randomInGameMode() InGameMode {
	rand.Seed(int64(time.Now().Nanosecond()))
	return inGameMode[rand.Intn(len(inGameMode))]
}

func randomRankLevel() RankLevelType {
	rand.Seed(int64(time.Now().Nanosecond()))
	return rankLevel[rand.Intn(len(rankLevel))]
}

func getMinMaxTrophies(rank RankLevelType) (min int32, max int32) {
	switch rank {
	case RankRecruit:
		return 48, 527
	case RankPrivate:
		return 528, 1727
	case RankBronze:
		return 1728, 4727
	case RankSilver:
		return 4728, 9727
	case RankGolden:
		return 9728, 16727
	case RankPlatinum:
		return 16728, 26727
	case RankDiamond:
		return 26728, 37977
	case RankMaster:
		return 37978, 52977
	case RankChampion:
		return 52978, 60000
	}
	return 0, 0
}

func randate(min int64, max int64) time.Time {
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func GenerateUserData(numUser int) []UserData {
	var listUser []UserData
	for i := 0; i < numUser; i++ {
		rand.Seed(time.Now().UnixNano())
		ingameMode := randomInGameMode()
		skill1 := randomSkill()
		skill2 := randomSkill()
		rank := randomRankLevel()
		minRank, maxRank := getMinMaxTrophies(rank)
		trophies := int32(rand.Intn(int(maxRank-minRank))) + minRank
		herotypeId := rand.Int31n(24)
		heroLevel := rand.Int31n(11)
		skillLevel := rand.Int31n(11)
		playerId := strings.Replace(uuid.New().String(), "-", "", -1)
		heroId := strings.Replace(uuid.New().String(), "-", "", -1)
		battleCount := rand.Int31n(100)
		dateMin := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
		dateMax1 := time.Date(2021, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
		dateMax2 := time.Date(2022, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
		region := randomRegion()
		{

			userdatadetail := UserDetail{
				Id:                 playerId,
				Email:              playerId[:6] + "@gmail.com",
				CreatedAt:          randate(dateMin, dateMax1),
				OpenedAt:           randate(dateMax1, dateMax2),
				AddressConnectTime: randate(dateMax1, dateMax2),
			}

			userdata := UserData{}
			userdata.InGameMode = ingameMode
			userdata.InGameModeString = ingameMode.String()
			userdata.HeroTypeId = herotypeId
			userdata.HeroTypeIdString = strconv.Itoa(int(herotypeId))
			userdata.Skill1 = skill1
			userdata.Skill1String = strconv.Itoa(int(skill1))
			userdata.Skill2 = skill2
			userdata.Skill2String = strconv.Itoa(int(skill2))
			userdata.Trophies = trophies
			userdata.TrophiesString = strconv.Itoa(int(trophies))
			userdata.HeroLevel = heroLevel
			userdata.HeroLevelString = strconv.Itoa(int(heroLevel))
			userdata.Skill1Level = skillLevel
			userdata.Skill2Level = skillLevel
			userdata.Skill1LevelString = strconv.Itoa(int(skillLevel))
			userdata.Skill2LevelString = strconv.Itoa(int(skillLevel))
			userdata.PlayerName = playerId
			userdata.AvatarId = 1
			userdata.HeroId = heroId
			userdata.BattleCount = battleCount
			userdata.BattleCountString = strconv.Itoa(int(battleCount))
			userdata.UserDataDetail = userdatadetail
			userdata.Region = []int32{region}

			listUser = append(listUser, userdata)
		}
	}

	return listUser
}
