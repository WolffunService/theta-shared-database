package usermodel

import "time"

type PlayerStatAccount struct {
	FirstOpenDate     time.Time `json:"firstOpenDate"`
	ConnectWalletDate time.Time `json:"connectWalletDate"`
	AccountDate       time.Time `json:"accountDate"`
	AccountAge        int32     `json:"accountAge"`
}

type PlayerStatBalance struct {
	THCBalance int64
	THGBalance int64
	PPBalance  int64
	PTBalance  int64
}

type PlayerStatMarketplace struct {
}

type PlayerStatSocial struct {
}

type PlayerStat struct {
	PlayerId              string  `json:"userId"`
	PlayerDetail          NewUser `json:"userDetail"`
	PlayerStatAccount     PlayerStatAccount
	PlayerStatMarketplace PlayerStatMarketplace
	PlayerStatBalance     PlayerStatBalance
}
