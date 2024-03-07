package pal_server

import (
	"errors"
	"fmt"
	"web-server/global"
	model "web-server/model/pal_server"
	"web-server/model/pal_server/request"
	"web-server/utils"
)

type PalSaveArchiveService struct{}

// SaveArchive 存档保存
func (s *PalSaveArchiveService) SaveArchive(info request.ArchiveSave) (err error) {
	sourcePath := fmt.Sprintf("%sPal/Saved", global.PalConfig.Palu.ServerPath)

	// 创建备份目录
	// 示例：/Users/tim/OrbStack/docker/containers/palserver/home/steam/Steam/steamapps/common/PalServer/backups
	backupPath := fmt.Sprintf("%sbackups/", global.PalConfig.Palu.ServerPath)

	// 执行备份
	fileName, err := utils.TarGzDir(sourcePath, backupPath)
	if err != nil {
		return errors.New(fmt.Sprintf("备份失败: %s", err.Error()))
	}

	// 获取备份文件大小
	archiveSize, err := utils.GetFileSize(backupPath + fileName)
	if err != nil {
		return errors.New("获取备份文件大小失败")
	}

	// 存储备份信息
	if err = global.PalDB.Create(&model.PalArchive{
		ArchiveName: info.ArchiveName,
		ArchiveDesc: info.ArchiveDesc,
		ArchivePath: backupPath + fileName,
		ArchiveSize: archiveSize,
	}).Error; err != nil {
		return errors.New("备份信息存储失败")
	}

	return err
}

// RestoreArchive 存档恢复
func (s *PalSaveArchiveService) RestoreArchive(backFileID uint) (err error) {
	// 示例：/Users/tim/OrbStack/docker/containers/palserver/home/steam/Steam/steamapps/common/PalServer/Pal/Saved
	savedPath := fmt.Sprintf("%sPal/Saved", global.PalConfig.Palu.ServerPath)
	var archive model.PalArchive
	if err = global.PalDB.Where("id = ?", backFileID).First(&archive).Error; err != nil {
		return errors.New("备份文件不存在")
	}

	// 执行恢复
	err = utils.UnTarGzDir(archive.ArchivePath, savedPath)
	if err != nil {
		return errors.New("恢复失败")
	}
	return
}

// GetArchiveList 分页获取存档列表
func (s *PalSaveArchiveService) GetArchiveList(info request.ArchiveList) (data []model.PalArchive, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.PalDB.Model(&model.PalArchive{})
	var archiveList []model.PalArchive
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&archiveList).Error
	return archiveList, total, err
}

// DeleteArchive 存档删除
func (s *PalSaveArchiveService) DeleteArchive(backFileID uint) (err error) {
	var archive model.PalArchive
	if err = global.PalDB.Where("id = ?", backFileID).First(&archive).Error; err != nil {
		return errors.New("备份文件不存在")
	}

	// 执行删除
	err = utils.Removefile(archive.ArchivePath)
	if err != nil {
		return errors.New("删除失败")
	}

	// 删除数据库记录
	err = global.PalDB.Delete(&archive).Error
	return
}
