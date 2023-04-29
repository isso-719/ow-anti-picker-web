package src

import (
	owantipicker "github.com/isso-719/ow-anti-picker"
	"github.com/isso-719/ow-anti-picker/owhero"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type reqBody struct {
	EnemyHeroList []string `json:"enemy_hero_list"`
}

func GetAntiPick(c echo.Context) error {
	var reqBody reqBody
	if err := c.Bind(&reqBody); err != nil {
		return err
	}

	enemyHeroListJSON := reqBody.EnemyHeroList

	var enemyHeroList []owhero.HeroID
	for _, heroIDStr := range enemyHeroListJSON {
		heroID, err := strconv.Atoi(heroIDStr)
		if err != nil {
			return err
		}
		enemyHeroList = append(enemyHeroList, owhero.HeroID(heroID))
	}

	antiPickMap := owantipicker.GetAntiPickMap(enemyHeroList)
	antiPickList := antiPickMap.GetAntiPickMapToSortByImportantListWithHeroName()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"anti_pick_list": antiPickList,
	})
}
