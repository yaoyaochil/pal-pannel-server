package model

// PalServerOption 结构体用于定义 Palworld 服务器选项
type PalServerOption struct {
	Difficulty                           string  `json:"Difficulty"`                           // 难度
	DayTimeSpeedRate                     float64 `json:"DayTimeSpeedRate"`                     // 白天流逝速度 (0.1 - 5)
	NightTimeSpeedRate                   float64 `json:"NightTimeSpeedRate"`                   // 夜间流逝速度 (0.1 - 5)
	ExpRate                              float64 `json:"ExpRate"`                              // 经验值倍率
	PalCaptureRate                       float64 `json:"PalCaptureRate"`                       // 捕捉概率倍率 (0.5 - 2)
	PalSpawnNumRate                      float64 `json:"PalSpawnNumRate"`                      // 帕鲁出现数量倍率
	PalDamageRateAttack                  float64 `json:"PalDamageRateAttack"`                  // 帕鲁攻击伤害倍率
	PalDamageRateDefense                 float64 `json:"PalDamageRateDefense"`                 // 帕鲁承受伤害倍率
	PlayerDamageRateAttack               float64 `json:"PlayerDamageRateAttack"`               // 玩家攻击伤害倍率 (0.1 - 5)
	PlayerDamageRateDefense              float64 `json:"PlayerDamageRateDefense"`              // 玩家承受伤害倍率 (0.1 - 5)
	PlayerStomachDecreaceRate            float64 `json:"PlayerStomachDecreaceRate"`            // 玩家饱食度降低倍率 (0.1 - 5)
	PlayerStaminaDecreaceRate            float64 `json:"PlayerStaminaDecreaceRate"`            // 玩家耐力降低倍率 (0.1 - 5)
	PlayerAutoHPRegeneRate               float64 `json:"PlayerAutoHPRegeneRate"`               // 玩家生命值自然回复倍率
	PlayerAutoHpRegeneRateInSleep        float64 `json:"PlayerAutoHpRegeneRateInSleep"`        // 玩家睡眠时生命值回复倍率
	PalStomachDecreaceRate               float64 `json:"PalStomachDecreaceRate"`               // 帕鲁饱食度降低倍率 (0.1 - 5)
	PalStaminaDecreaceRate               float64 `json:"PalStaminaDecreaceRate"`               // 帕鲁耐力降低倍率 (0.1 - 5)
	PalAutoHPRegeneRate                  float64 `json:"PalAutoHPRegeneRate"`                  // 帕鲁生命值自然回复倍率
	PalAutoHpRegeneRateInSleep           float64 `json:"PalAutoHpRegeneRateInSleep"`           // 帕鲁睡眠时生命值回复倍率
	BuildObjectDamageRate                float64 `json:"BuildObjectDamageRate"`                // 对建筑物伤害倍率 (0.5 - 3)
	BuildObjectDeteriorationDamageRate   float64 `json:"BuildObjectDeteriorationDamageRate"`   // 非基地圈内建筑物的劣化速度倍率 (0 - 10)
	CollectionDropRate                   float64 `json:"CollectionDropRate"`                   // 道具采集量倍率
	CollectionObjectHpRate               float64 `json:"CollectionObjectHpRate"`               // 可采集物品生命值倍率 (0.5 - 3)
	CollectionObjectRespawnSpeedRate     float64 `json:"CollectionObjectRespawnSpeedRate"`     // 可采集物品重生间隔倍率 (0；0.5 - 3)
	EnemyDropItemRate                    float64 `json:"EnemyDropItemRate"`                    // 道具掉落量倍率
	DeathPenalty                         string  `json:"DeathPenalty"`                         // 死亡惩罚
	BEnablePlayerToPlayerDamage          bool    `json:"bEnablePlayerToPlayerDamage"`          // 启用玩家对玩家伤害
	BEnableFriendlyFire                  bool    `json:"bEnableFriendlyFire"`                  // 启用友伤
	BEnableInvaderEnemy                  bool    `json:"bEnableInvaderEnemy"`                  // 启用袭击事件
	BActiveUNKO                          bool    `json:"bActiveUNKO"`                          // 激活帕鲁便便
	BEnableAimAssistPad                  bool    `json:"bEnableAimAssistPad"`                  // 启用手柄瞄准辅助
	BEnableAimAssistKeyboard             bool    `json:"bEnableAimAssistKeyboard"`             // 启用键盘瞄准辅助
	DropItemMaxNum                       int     `json:"DropItemMaxNum"`                       // 掉落物品最大存在数量
	DropItemMaxNum_UNKO                  int     `json:"DropItemMaxNum_UNKO"`                  // 帕鲁便便掉落最大数量
	BaseCampMaxNum                       int     `json:"BaseCampMaxNum"`                       // 全地图据点最大数量
	BaseCampWorkerMaxNum                 int     `json:"BaseCampWorkerMaxNum"`                 // 可分配至据点工作的帕鲁数量上限
	DropItemAliveMaxHours                float64 `json:"DropItemAliveMaxHours"`                // 掉落物品存活最大小时数
	BAutoResetGuildNoOnlinePlayers       bool    `json:"bAutoResetGuildNoOnlinePlayers"`       // 自动重置无在线玩家的公会
	AutoResetGuildTimeNoOnlinePlayers    float64 `json:"AutoResetGuildTimeNoOnlinePlayers"`    // 自动重置无在线玩家的公会时间(小时)
	GuildPlayerMaxNum                    int     `json:"GuildPlayerMaxNum"`                    // 公会玩家最大数量
	PalEggDefaultHatchingTime            float64 `json:"PalEggDefaultHatchingTime"`            // 巨大蛋孵化所需时间(小时)
	WorkSpeedRate                        float64 `json:"WorkSpeedRate"`                        // 工作速率
	BIsMultiplay                         bool    `json:"bIsMultiplay"`                         // 是否多人游戏
	BIsPvP                               bool    `json:"bIsPvP"`                               // 是否 PvP
	BCanPickupOtherGuildDeathPenaltyDrop bool    `json:"bCanPickupOtherGuildDeathPenaltyDrop"` // 是否捡取其他公会的死亡惩罚掉落
	BEnableNonLoginPenalty               bool    `json:"bEnableNonLoginPenalty"`               // 启用超时未登录惩罚
	BEnableFastTravel                    bool    `json:"bEnableFastTravel"`                    // 启用快速传送
	BIsStartLocationSelectByMap          bool    `json:"bIsStartLocationSelectByMap"`          // 是否通过地图选择复活位置
	BExistPlayerAfterLogout              bool    `json:"bExistPlayerAfterLogout"`              // 登出后玩家人物是否存在
	BEnableDefenseOtherGuildPlayer       bool    `json:"bEnableDefenseOtherGuildPlayer"`       // 启用据点内防御其他公会玩家
	CoopPlayerMaxNum                     int     `json:"CoopPlayerMaxNum"`                     // 合作玩家最大数量
	BShowPlayerList                      bool    `json:"bShowPlayerList"`                      // 启用玩家列表显示
	ServerPlayerMaxNum                   int     `json:"ServerPlayerMaxNum"`                   // 服务器最大玩家数量
	ServerName                           string  `json:"ServerName"`                           // 服务器名称
	ServerDescription                    string  `json:"ServerDescription"`                    // 服务器描述
	AdminPassword                        string  `json:"AdminPassword"`                        // 管理员密码
	ServerPassword                       string  `json:"ServerPassword"`                       // 服务器密码
	PublicPort                           int     `json:"PublicPort"`                           // 公开端口号
	PublicIP                             string  `json:"PublicIP"`                             // 公开 IP 地址
	RCONEnabled                          bool    `json:"RCONEnabled"`                          // 是否启用 RCON
	RCONPort                             int     `json:"RCONPort"`                             // RCON 端口号
	Region                               string  `json:"Region"`                               // 区域
	BUseAuth                             bool    `json:"bUseAuth"`                             // 是否使用授权服务
	BanListURL                           string  `json:"BanListURL"`                           // 封禁列表 URL
}
