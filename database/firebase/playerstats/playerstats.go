package playerstats

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WolffunService/theta-shared-common/common/enum/playerstatenum"
	"github.com/WolffunService/theta-shared-database/common/util/request"
	"github.com/WolffunService/theta-shared-database/database/firebase/playerstatsmodel"
)

var thetanFireBaseAddress string = ""

func Initialize(address string) {
	thetanFireBaseAddress = address
}

func GetPlayerStatWithStatName(userId string, statName playerstatenum.StatName) (int64, error) {
	return getPlayerStatFull(userId, statName.String())
}

func GetPlayerStatWithCode(userId string, statName playerstatenum.StatName, statCode int) (int64, error) {
	return getPlayerStatFull(userId, statName.GetWithCode(statCode))
}

func getPlayerStatFull(userId string, statName string) (int64, error) {
	operation := "GetPlayerStatFull"

	if thetanFireBaseAddress == "" {
		return 0, fmt.Errorf("ThetanFireBaseAddress is empty")
	}
	method := "/internal/playerStats/getPlayerStatFull"
	var mapQueryParams = make(map[string]string)
	mapQueryParams["UserId"] = userId
	mapQueryParams["StatName"] = statName
	data, statusCode, err := request.NewGetRequestParams(thetanFireBaseAddress, method, mapQueryParams)
	if err != nil {
		return 0, err
	}
	var res = playerstatsmodel.PlayerGetInt64Response{}
	if statusCode == http.StatusOK {
		errParseJson := json.Unmarshal(data, &res)
		if errParseJson != nil {
			return 0, errParseJson
		}
		return res.Data, nil
	}
	if data != nil {
		errParseJson := json.Unmarshal(data, &res)
		if errParseJson != nil {
			return 0, errParseJson
		}
	}

	return 0, fmt.Errorf("%v err, statusCode: %v", operation, statusCode)
}

func TrackPlayerStatWithStatName(userId string, statName playerstatenum.StatName, statValue int64, isFullValue bool) (isFirst bool, _ error) {
	return trackPlayerStat(userId, statName.String(), statValue, isFullValue)
}

func TrackPlayerStatWithCode(userId string, statName playerstatenum.StatName, statCode int, statValue int64, isFullValue bool) (isFirst bool, _ error) {
	return trackPlayerStat(userId, statName.GetWithCode(statCode), statValue, isFullValue)
}

func trackPlayerStat(userId string, statName string, statValue int64, isFullValue bool) (isFirst bool, _ error) {
	operation := "trackPlayerStat"

	isFirstRS := false

	if thetanFireBaseAddress == "" {
		return isFirstRS, fmt.Errorf("ThetanFireBaseAddress is empty")
	}

	if !isFullValue {
		stat, err := getPlayerStatFull(userId, statName)
		if err != nil {
			return isFirstRS, err
		}
		// Check if is first track
		if stat == 0 {
			isFirstRS = true
		}
	}
	method := "/internal/playerStats/trackPlayerStats"
	req := playerstatsmodel.TrackStatsRequest{
		UserId:    userId,
		StatName:  statName,
		StatValue: statValue,
	}
	statusCode, data, err := request.NewPostRequest(thetanFireBaseAddress, method, req, "")
	if err != nil {
		return isFirstRS, err
	}
	var res = playerstatsmodel.PlayerTrackStatResponse{}
	if statusCode == http.StatusOK {
		errParseJson := json.Unmarshal(data, &res)
		if errParseJson != nil {
			return isFirstRS, errParseJson
		}
		return isFirstRS, nil
	}
	if data != nil {
		errParseJson := json.Unmarshal(data, &res)
		if errParseJson != nil {
			return isFirstRS, errParseJson
		}
	}

	return isFirstRS, fmt.Errorf("%v err, statusCode: %v", operation, statusCode)
}
