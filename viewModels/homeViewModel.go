package viewModels

import (
	"github.com/kcwinner/visualKPIs/models"
)

//HomeViewModel Home view model for data template
type HomeViewModel struct {
	APFTData models.APFTData
	SSDData  models.SSDData
	Soldiers []models.Soldier
}
