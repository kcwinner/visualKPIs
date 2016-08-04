package viewModels

import (
	"github.com/kcwinner/visualKPIs/models"
)

//SlideshowViewModel Slideshow view model for data template
type SlideshowViewModel struct {
	Soldiers []models.Soldier
	APFTData models.APFTData
	SSDData  models.SSDData
}
