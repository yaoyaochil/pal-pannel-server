package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"reflect"
	"strings"
	model "web-server/model/pal_server"
)

// ReadIni 读取PalWorldSettings.ini文件
func ReadIni() (data model.PalServerOption, err error) {
	sourc, err := os.ReadFile("source/game/PalWorldSettings.ini")
	if err != nil {
		return data, err
	}
	input := string(sourc)
	formattedOutput := parseAndFormat(input)

	cfg, err := ini.Load([]byte(formattedOutput))
	if err != nil {
		return data, err
	}

	bEnablePlayerToPlayerDamage, err := cfg.Section("").Key("bEnablePlayerToPlayerDamage").Bool()
	bEnableFriendlyFire, err := cfg.Section("").Key("bEnableFriendlyFire").Bool()
	bEnableInvaderEnemy, err := cfg.Section("").Key("bEnableInvaderEnemy").Bool()
	bActiveUNKO, err := cfg.Section("").Key("bActiveUNKO").Bool()
	bEnableAimAssistPad, err := cfg.Section("").Key("bEnableAimAssistPad").Bool()
	bEnableAimAssistKeyboard, err := cfg.Section("").Key("bEnableAimAssistKeyboard").Bool()
	BAutoResetGuildNoOnlinePlayers, err := cfg.Section("").Key("bAutoResetGuildNoOnlinePlayers").Bool()
	bIsMultiplay, err := cfg.Section("").Key("bIsMultiplay").Bool()
	bIsPvP, err := cfg.Section("").Key("bIsPvP").Bool()
	bCanPickupOtherGuildDeathPenaltyDrop, err := cfg.Section("").Key("bCanPickupOtherGuildDeathPenaltyDrop").Bool()
	bEnableNonLoginPenalty, err := cfg.Section("").Key("bEnableNonLoginPenalty").Bool()
	bEnableFastTravel, err := cfg.Section("").Key("bEnableFastTravel").Bool()
	bIsStartLocationSelectByMap, err := cfg.Section("").Key("bIsStartLocationSelectByMap").Bool()
	bExistPlayerAfterLogout, err := cfg.Section("").Key("bExistPlayerAfterLogout").Bool()
	bEnableDefenseOtherGuildPlayer, err := cfg.Section("").Key("bEnableDefenseOtherGuildPlayer").Bool()
	bShowPlayerList, err := cfg.Section("").Key("bShowPlayerList").Bool()
	RCONEnabled, err := cfg.Section("").Key("RCONEnabled").Bool()
	BUseAuth, err := cfg.Section("").Key("bUseAuth").Bool()
	err = cfg.Section("").MapTo(&data)
	if err != nil {
		return data, err
	}

	var boolData model.PalServerOption
	boolData = model.PalServerOption{
		BEnablePlayerToPlayerDamage:          bEnablePlayerToPlayerDamage,
		BEnableFriendlyFire:                  bEnableFriendlyFire,
		BEnableInvaderEnemy:                  bEnableInvaderEnemy,
		BActiveUNKO:                          bActiveUNKO,
		BEnableAimAssistPad:                  bEnableAimAssistPad,
		BEnableAimAssistKeyboard:             bEnableAimAssistKeyboard,
		BAutoResetGuildNoOnlinePlayers:       BAutoResetGuildNoOnlinePlayers,
		BIsMultiplay:                         bIsMultiplay,
		BIsPvP:                               bIsPvP,
		BCanPickupOtherGuildDeathPenaltyDrop: bCanPickupOtherGuildDeathPenaltyDrop,
		BEnableNonLoginPenalty:               bEnableNonLoginPenalty,
		BEnableFastTravel:                    bEnableFastTravel,
		BIsStartLocationSelectByMap:          bIsStartLocationSelectByMap,
		BExistPlayerAfterLogout:              bExistPlayerAfterLogout,
		BEnableDefenseOtherGuildPlayer:       bEnableDefenseOtherGuildPlayer,
		BShowPlayerList:                      bShowPlayerList,
		RCONEnabled:                          RCONEnabled,
		BUseAuth:                             BUseAuth,
	}

	updateStructFields(&data, boolData)

	return data, nil
}

// parseAndFormat 解析并格式化PalWorldSettings.ini文件
func parseAndFormat(input string) string {
	// 去除头尾的方括号和末尾的括号
	input = strings.TrimPrefix(input, "[/Script/Pal.PalGameWorldSettings]\nOptionSettings=(")
	input = strings.TrimSuffix(input, ")")

	// 替换所有的,为\n
	input = strings.ReplaceAll(input, ",", "\n")

	return input
}

// WriteIni 写入PalWorldSettings.ini文件
func WriteIni(data model.PalServerOption) (err error) {
	source := fmt.Sprintf(`[/Script/Pal.PalGameWorldSettings]
OptionSettings=(Difficulty=%s,DayTimeSpeedRate=%.6f,NightTimeSpeedRate=%.6f,ExpRate=%.6f,PalCaptureRate=%.6f,PalSpawnNumRate=%.6f,PalDamageRateAttack=%.6f,PalDamageRateDefense=%.6f,PlayerDamageRateAttack=%.6f,PlayerDamageRateDefense=%.6f,PlayerStomachDecreaceRate=%.6f,PlayerStaminaDecreaceRate=%.6f,PlayerAutoHPRegeneRate=%.6f,PlayerAutoHpRegeneRateInSleep=%.6f,PalStomachDecreaceRate=%.6f,PalStaminaDecreaceRate=%.6f,PalAutoHPRegeneRate=%.6f,PalAutoHpRegeneRateInSleep=%.6f,BuildObjectDamageRate=%.6f,BuildObjectDeteriorationDamageRate=%.6f,CollectionDropRate=%.6f,CollectionObjectHpRate=%.6f,CollectionObjectRespawnSpeedRate=%.6f,EnemyDropItemRate=%.6f,DeathPenalty=%s,bEnablePlayerToPlayerDamage=%t,bEnableFriendlyFire=%t,bEnableInvaderEnemy=%t,bActiveUNKO=%t,bEnableAimAssistPad=%t,bEnableAimAssistKeyboard=%t,DropItemMaxNum=%d,DropItemMaxNum_UNKO=%d,BaseCampMaxNum=%d,BaseCampWorkerMaxNum=%d,DropItemAliveMaxHours=%.6f,bAutoResetGuildNoOnlinePlayers=%t,AutoResetGuildTimeNoOnlinePlayers=%.6f,GuildPlayerMaxNum=%d,PalEggDefaultHatchingTime=%.6f,WorkSpeedRate=%.6f,bIsMultiplay=%t,bIsPvP=%t,bCanPickupOtherGuildDeathPenaltyDrop=%t,bEnableNonLoginPenalty=%t,bEnableFastTravel=%t,bIsStartLocationSelectByMap=%t,bExistPlayerAfterLogout=%t,bEnableDefenseOtherGuildPlayer=%t,bShowPlayerList=%t,CoopPlayerMaxNum=%d,ServerPlayerMaxNum=%d,ServerName="%s",ServerDescription="%s",AdminPassword="%s",ServerPassword="%s",PublicPort=%d,PublicIP="%s",RCONEnabled=%t,RCONPort=%d,Region="%s",bUseAuth=%t,BanListURL="%s")`,
		data.Difficulty, data.DayTimeSpeedRate, data.NightTimeSpeedRate, data.ExpRate, data.PalCaptureRate, data.PalSpawnNumRate, data.PalDamageRateAttack, data.PalDamageRateDefense, data.PlayerDamageRateAttack, data.PlayerDamageRateDefense, data.PlayerStomachDecreaceRate, data.PlayerStaminaDecreaceRate, data.PlayerAutoHPRegeneRate, data.PlayerAutoHpRegeneRateInSleep, data.PalStomachDecreaceRate, data.PalStaminaDecreaceRate, data.PalAutoHPRegeneRate, data.PalAutoHpRegeneRateInSleep, data.BuildObjectDamageRate, data.BuildObjectDeteriorationDamageRate, data.CollectionDropRate, data.CollectionObjectHpRate, data.CollectionObjectRespawnSpeedRate, data.EnemyDropItemRate, data.DeathPenalty, data.BEnablePlayerToPlayerDamage, data.BEnableFriendlyFire, data.BEnableInvaderEnemy, data.BActiveUNKO, data.BEnableAimAssistPad, data.BEnableAimAssistKeyboard, data.DropItemMaxNum, data.DropItemMaxNum_UNKO, data.BaseCampMaxNum, data.BaseCampWorkerMaxNum, data.DropItemAliveMaxHours, data.BAutoResetGuildNoOnlinePlayers, data.AutoResetGuildTimeNoOnlinePlayers, data.GuildPlayerMaxNum, data.PalEggDefaultHatchingTime, data.WorkSpeedRate, data.BIsMultiplay, data.BIsPvP, data.BCanPickupOtherGuildDeathPenaltyDrop, data.BEnableNonLoginPenalty, data.BEnableFastTravel, data.BIsStartLocationSelectByMap, data.BExistPlayerAfterLogout, data.BEnableDefenseOtherGuildPlayer, data.BShowPlayerList, data.CoopPlayerMaxNum, data.ServerPlayerMaxNum, data.ServerName, data.ServerDescription, data.AdminPassword, data.ServerPassword, data.PublicPort, data.PublicIP, data.RCONEnabled, data.RCONPort, data.Region, data.BUseAuth, data.BanListURL)
	source = strings.ReplaceAll(source, "true", "True")
	source = strings.ReplaceAll(source, "false", "False")
	err = os.WriteFile("source/game/PalWorldSettings.ini", []byte(source), 0644)
	if err != nil {
		return err
	}

	return nil
}

// 更新结构体字段的函数
func updateStructFields(source interface{}, target interface{}) {
	sourceValue := reflect.ValueOf(source).Elem()
	targetValue := reflect.ValueOf(target)

	for i := 0; i < sourceValue.NumField(); i++ {
		fieldName := sourceValue.Type().Field(i).Name
		targetFieldValue := targetValue.FieldByName(fieldName)

		if targetFieldValue.IsValid() && !isEmpty(targetFieldValue) {
			sourceValue.Field(i).Set(targetFieldValue)
		}
	}
}

// 判断字段是否为空的辅助函数
func isEmpty(field reflect.Value) bool {
	zero := reflect.Zero(field.Type())
	return reflect.DeepEqual(field.Interface(), zero.Interface())
}
