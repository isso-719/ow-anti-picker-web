package src

import (
	"github.com/isso-719/ow-anti-picker/owhero"
	"github.com/labstack/echo"
	"net/http"
)

type heroList struct {
	ID    owhero.HeroID
	Name  string
	Image string
}

func Index(c echo.Context) error {
	var tankHeroList []heroList
	var damageHeroList []heroList
	var supportHeroList []heroList
	for _, hero := range owhero.HeroList {
		heroImgUrl := HeroImageMap[hero.ID]

		switch hero.Role {
		case owhero.Tank:
			tankHeroList = append(tankHeroList, heroList{ID: hero.ID, Name: hero.Name, Image: heroImgUrl})
		case owhero.Damage:
			damageHeroList = append(damageHeroList, heroList{ID: hero.ID, Name: hero.Name, Image: heroImgUrl})
		case owhero.Support:
			supportHeroList = append(supportHeroList, heroList{ID: hero.ID, Name: hero.Name, Image: heroImgUrl})
		}
	}

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"tank_hero_list":    tankHeroList,
		"damage_hero_list":  damageHeroList,
		"support_hero_list": supportHeroList,
	})
}
