package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"ushield_bot/internal/application/services"
	"ushield_bot/internal/domain"
	"ushield_bot/pkg/switcher"
)

type IBot interface {
	SendMessage(message domain.MessageToSend, channel int) error
	GetUserStates() map[int64]string
	GetServices() *services.Service
	GetBot() *tgbotapi.BotAPI
	GetTaskManager() *switcher.TaskFlowManager
	GetCookie() string
}

const (
	CommandStart = "start"

	CommandHelp = "help"

	CommandScoreEnergy                 = "check"
	CommandExchangeEnergy              = "exchange_energy"
	CommandCheckBlacklist              = "bind"
	GET_TODAY_FROZEN_TOTAL             = "get_today_frozen_total"
	GET_TODAY_FROZEN_ADDRESSES         = "get_today_frozen_addresses"
	GET_PENDING_FROZEN_ADDRESSES       = "get_pending_frozen_addresses"
	GET_HISTORICAL_STATS               = "get_historical_addresses_stats"
	ASSOCIATED_RECOMMENDATION_RELATION = "associated_relation "
	ADDRESS_BEHAVIOR_REPORT            = "get_address_behavior_report "
	GET_VIP                            = "upgrade_vip "
	MONITOR_ADDRESS                    = "monitor_address"
	CommandGetAccount                  = "get_account"
)

const (
	WaitingId          = "WaitingId"
	WaitingGroup       = "WaitingGroup"
	WaitingName        = "WaitingName"
	WaitingDescription = "WaitingDescription"
	WaitingImages      = "WaitingImages"
	WaitingTags        = "WaitingTags"
	WaitingDeadline    = "WaitingDeadline"
)

const (
	DefaultChannel     = 0
	ChannelInformation = 2
	ChannelBot         = 5
)
