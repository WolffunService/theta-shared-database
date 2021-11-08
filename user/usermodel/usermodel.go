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
	Role                 int                `json:"role" bson:"role"`
	Version              int                `json:"version" bson:"version"`
	Email                string             `json:"email" bson:"email"`
	UserName             string             `json:"username" bson:"username"`
	NumChangeName        int                `json:"numChangeName" bson:"numChangeName"`
	Status               UserStatus         `json:"status" bson:"status"`
	Address              string             `json:"address" bson:"address"`
	Nonce                int                `json:"nonce" bson:"nonce"`
	CanClaimFreeHero     bool               `json:"canClaimFreeHero" bson:"canClaimFreeHero"`
	CanClaimBetaReward   bool               `json:"canClaimBetaReward" bson:"canClaimBetaReward"`
	UserProfile          UserProfile        `json:"userProfile" bson:"userProfile"`
	PlayerStatistic      PlayerStatistic    `json:"playerStatistic" bson:"playerStatistic"`
	Suspicious           int                `json:"-" bson:"suspicious"`
	SuspiciousWrongData  int                `json:"-" bson:"suspiciousWrongData"`
	Country              string             `json:"country" bson:"country"`
	AvatarId             int                `json:"avatarId" bson:"avatarId"`
	FrameId              int                `json:"frameId" bson:"frameId"`
	TicketBanFindMatch   TicketBanFindMatch `json:"-" bson:"ticketBanFindMatch"`
	Referral             Referral           `json:"referral" bson:"referral"`
	IsCreator            bool               `json:"isCreator" bson:"isCreator"`
}

type Referral struct {
	NumInviteFriends      int32  `json:"numInviteFriends" bson:"numInviteFriends"`
	TotalNumInviteFriends int32  `json:"totalNumInviteFriends" bson:"totalNumInviteFriends"`
	ReferralID            string `json:"referralID" bson:"referralID"`
	AcceptAt              int64  `json:"acceptAt" bson:"acceptAt"`
	Status                bool   `json:"status" bson:"status"`
}

type PlayerStatistic struct {
	Battle        int32 `json:"battle" bson:"battle"`
	Victory       int32 `json:"victory" bson:"victory"`
	Streak        int32 `json:"streak" bson:"streak"`
	CurStreak     int32 `json:"-" bson:"curStreak"`
	Triple        int32 `json:"triple" bson:"triple"`
	Mega          int32 `json:"mega" bson:"mega"`
	Mvp           int32 `json:"mvp" bson:"mvp"`
	Hero          int32 `json:"hero" bson:"hero"`
	BehaviorPoint int32 `json:"behaviorPoint" bson:"behaviorPoint"`
}

type TicketBanFindMatch struct {
	NumBans     int   `json:"numBans" bson:"numBans"`
	ExpiredTime int64 `json:"expiredTime" bson:"expiredTime"`
}

func (t TicketBanFindMatch) IsBanned() bool {
	return t.ExpiredTime >= time.Now().Unix()
}

func (u User) IsBanned() bool {
	return u.Status == BANNED
}

func (u *User) GetUserId() string {
	return u.ID.(primitive.ObjectID).Hex()
}

func (u *User) GetUserName() string {
	return u.UserName
}

func (u *User) GetAddress() string {
	return u.Address
}

func (u *User) GetRole() int {
	return u.Role
}

func (u *User) GetBehaviorPoint() int32 {
	return 100 - u.PlayerStatistic.BehaviorPoint
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

// Trạng thái điểm hành vi: Tốt, Khá, Tệ
type BehaviorStatus int

const (
	EXCELLENT BehaviorStatus = iota
	GOOD
	BAD
)

func GetBehaviorStatus(u *User) BehaviorStatus {
	// Behavior Point
	// 70 - 100: EXCELLENT
	// 50 - 69: GOOD
	// < 50: BAD

	bPoint := GetBehaviorStatus(u)
	if bPoint >= 70 {
		return EXCELLENT
	}
	if bPoint >= 50 {
		return GOOD
	}
	return BAD
}
