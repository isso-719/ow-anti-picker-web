package src

import (
	"github.com/isso-719/ow-anti-picker/owhero"
	"github.com/labstack/echo"
	"net/http"
)

type heroList struct {
	ID   owhero.HeroID
	Name string
}

func Index(c echo.Context) error {
	// allHeroList
	var allHeroList []heroList
	for _, hero := range owhero.HeroList {
		allHeroList = append(allHeroList, heroList{ID: hero.ID, Name: hero.Name})
	}

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"allHeroList": allHeroList,
	})
}
