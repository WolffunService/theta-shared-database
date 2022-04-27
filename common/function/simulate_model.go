package function

import (
	"time"
)

type InGameMode int32

const (
	InGameMode_SOLO_SURVIVAL            InGameMode = 1
	InGameMode_DUAL_SURVIVAL            InGameMode = 2
	InGameMode_TEAM_COLLECT_STAR_4_VS_4 InGameMode = 3
	InGameMode_DEATH_MATCH              InGameMode = 6
	InGameMode_TOWER                    InGameMode = 9
)

var inGameModeString = map[InGameMode]string{
	InGameMode_SOLO_SURVIVAL:            "InGameMode_SOLO_SURVIVAL",
	InGameMode_DUAL_SURVIVAL:            "InGameMode_DUAL_SURVIVAL",
	InGameMode_TEAM_COLLECT_STAR_4_VS_4: "InGameMode_TEAM_COLLECT_STAR_4_VS_4",
	InGameMode_DEATH_MATCH:              "InGameMode_DEATH_MATCH",
	InGameMode_TOWER:                    "InGameMode_TOWER",
}

func (i InGameMode) String() string {
	return inGameModeString[i]
}

type RankLevelType int32

// Type List Rank
const (
	RankRecruit RankLevelType = 1 + iota
	RankPrivate
	RankBronze
	RankSilver
	RankGolden
	RankPlatinum
	RankDiamond
	RankMaster
	RankChampion
)

type UserData struct {
	// heroId
	UserDataDetail    UserDetail
	InGameMode        InGameMode
	InGameModeString  string
	HeroTypeId        int32
	HeroTypeIdString  string
	Skill1            int32
	Skill1String      string
	Skill2            int32
	Skill2String      string
	Trophies          int32
	TrophiesString    string
	HeroLevel         int32
	HeroLevelString   string
	Skill1Level       int32
	Skill1LevelString string
	Skill2Level       int32
	Skill2LevelString string
	PlayerName        string
	AvatarId          int32
	HeroId            string
	TrophyMatching    int32
	BattleCount       int32
	BattleCountString string
	Region            []int32
}

type UserDetail struct {
	Id                   string
	CreatedAt            time.Time
	OpenedAt             time.Time
	Role                 int
	Version              int
	Email                string
	PublicEmail          string
	UserName             string
	NumChangeName        int
	Address              string
	AddressConnectTime   time.Time
	Nonce                int
	CanClaimFreeHero     bool
	CanClaimBetaReward   bool
	Suspicious           int
	SuspiciousWrongData  int
	SuspiciousAbnormal   int
	IpAddress            string
	Country              string
	AvatarId             int
	FrameId              int
	IsCreator            bool
	IsCreatorProgram     bool
	BanReason            string
	LastTimeBattle       time.Time
	IsCheckBehaviorPoint bool
	IsBot                bool

	// TODO: tmp field
	HasNewAvatar bool `json:"hasNewAvatar" bson:"hasNewAvatar"`
}
