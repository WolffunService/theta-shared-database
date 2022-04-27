package usermodel

import (
	"github.com/WolffunGame/theta-shared-database/database/mongodb"
	"time"
)

func (UserRanking) CollectionName() string {
	return "UserRanking"
}

type UserRanking struct {
	mongodb.DefaultModel `json:",inline" bson:",inline"`
	Trophy               int       `json:"trophy" bson:"trophy"`
	TrophyHighest        int       `json:"trophyHighest" bson:"trophyHighest"`
	SeasonTrophyHighest  int       `json:"seasonTrophyHighest" bson:"seasonTrophyHighest"`
	RankingLevel         int       `json:"rankingLevel" bson:"rankingLevel"`
	RankingLevelHighest  int       `json:"rankingLevelHighest" bson:"rankingLevelHighest"`
	TrophyCurRank        int       `json:"trophyCurRank" bson:"trophyCurRank"`
	LockedSkills         []int     `json:"lockedSkills"`
	CreateAt             time.Time `json:"createAt"`
}
