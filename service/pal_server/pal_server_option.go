package pal_server

import (
	"errors"
	"fmt"
	"os"
	"web-server/global"
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

// unlockTo20 解锁据点上限到20
func (s *PalServerOptionService) UnlockTo20(is bool) (err error) {
	savePath, err := os.Open(fmt.Sprintf("%sPal/Saved/SaveGames/0", global.PalConfig.Palu.ServerPath))
	if err != nil {
		return err
	}

	defer savePath.Close()

	firstFloder, err := savePath.Readdir(1)
	if err != nil {
		return err
	}

	switch is {
	case true:
		// 拷贝source->game->WorldOption.sav到global.PalConfig.Palu.ServerPath/Pal/Saved/SaveGames/0/第一个文件夹下面
		err = utils.CopyFile("source/game/WorldOption.sav", fmt.Sprintf("%sPal/Saved/SaveGames/0/%s/WorldOption.sav", global.PalConfig.Palu.ServerPath, firstFloder[0].Name()))
		if err != nil {
			return err
		}
	case false:
		// 删除global.PalConfig.Palu.ServerPath/Pal/Saved/SaveGames/0/第一个文件夹下面的WorldOption.sav
		err = os.Remove(fmt.Sprintf("%sPal/Saved/SaveGames/0/%s/WorldOption.sav", global.PalConfig.Palu.ServerPath, firstFloder[0].Name()))
		if err != nil {
			return err
		}
	}
	return nil
}

// checkUnlockTo20 检查是否解锁据点上限到20
func (s *PalServerOptionService) CheckUnlockTo20() (isUnlock bool, err error) {
	savePath, err := os.Open(fmt.Sprintf("%sPal/Saved/SaveGames/0", global.PalConfig.Palu.ServerPath))
	if err != nil {
		return false, errors.New("没有存档文件夹")
	}

	defer savePath.Close()

	firstFloder, err := savePath.Readdir(1)
	if err != nil {
		return false, errors.New("没有存档文件夹")
	}

	// 检查是否存在global.PalConfig.Palu.ServerPath/Pal/Saved/SaveGames/0/第一个文件夹下面的WorldOption.sav
	_, err = os.Stat(fmt.Sprintf("%sPal/Saved/SaveGames/0/%s/WorldOption.sav", global.PalConfig.Palu.ServerPath, firstFloder[0].Name()))
	if err != nil {
		return false, nil
	}
	isUnlock = true
	return isUnlock, nil
}
