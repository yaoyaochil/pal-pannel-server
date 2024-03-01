package pal_server

import (
	model "web-server/model/pal_server"
	"web-server/utils"
)

type PalServerOptionService struct{}

// Save 保存幻兽帕鲁服务器配置
func (s *PalServerOptionService) Save(request model.PalServerOption) (response model.PalServerOption, err error) {
	response = request
	return
}

// Get 获取幻兽帕鲁服务器配置
func (s *PalServerOptionService) Get() (response model.PalServerOption, err error) {
	response, err = utils.ReadIni()
	return
}

// WriteIni 写入PalWorldSettings.ini文件
func (s *PalServerOptionService) WriteIni(request model.PalServerOption) (err error) {
	err = utils.WriteIni(request)
	return
}
