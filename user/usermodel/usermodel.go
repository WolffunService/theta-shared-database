package usermodel

import (
	"github.com/WolffunGame/theta-shared-database/database/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (User) CollectionName() string {
	return "Users"
}

type User struct {
	mongodb.DefaultModel `bson:",inline"`
	mongodb.DateFields   `bson:",inline"`
	Version              int                `json:"version" bson:"version"`
	Email                string             `json:"email" bson:"email"`
	UserName             string             `json:"username" bson:"username"`
	NumChangeName        int                `json:"numChangeName" bson:"numChangeName"`
	Status               UserStatus         `json:"status" bson:"status"`
	Address              string             `json:"address" bson:"address"`
	Nonce                int                `json:"nonce" bson:"nonce"`
	IsClaimedFreeHero    bool               `json:"canClaimFreeHero" bson:"canClaimFreeHero"`
	UserProfile          UserProfile        `json:"userProfile" bson:"userProfile"`
	PlayerStatistic      PlayerStatistic    `json:"playerStatistic" bson:"playerStatistic"`
	Suspicious           int                `json:"-" bson:"suspicious"`
	SuspiciousWrongData  int                `json:"-" bson:"suspiciousWrongData"`
	Country              string             `json:"country" bson:"country"`
	AvatarId             int                `json:"avatarId" bson:"avatarId"`
	FrameId              int                `json:"frameId" bson:"frameId"`
	TicketBanFindMatch   TicketBanFindMatch `json:"-" bson:"ticketBanFindMatch"`
}

type PlayerStatistic struct {
	Battle    int32 `json:"battle" bson:"battle"`
	Victory   int32 `json:"victory" bson:"victory"`
	Streak    int32 `json:"streak" bson:"streak"`
	CurStreak int32 `json:"-" bson:"curStreak"`
	Triple    int32 `json:"triple" bson:"triple"`
	Mega      int32 `json:"mega" bson:"mega"`
	Mvp       int32 `json:"mvp" bson:"mvp"`
	Hero      int32 `json:"hero" bson:"hero"`
}

type TicketBanFindMatch struct {
	NumBans  int   `json:"numBans" bson:"numBans"`
	Duration int64 `json:"duration" bson:"duration"`
}

func (t TicketBanFindMatch) IsBanned() bool {
	return t.Duration >= time.Now().Unix()
}

func (u User) IsBanned() bool {
	return u.Status == BANNED
}

func (user *User) GetUserId() string {
	return user.ID.(primitive.ObjectID).Hex()
}

type UserStatus int

const (
	ACTIVE UserStatus = 1
	BANNED UserStatus = -1
)

type UserProfile struct {
	Level      int `bson:"level" json:"level"`
	XP         int `bson:"xp" json:"xp"`
	LevelUpGPP int `bson:"levelUpGPP" json:"levelUpGPP"`
}
