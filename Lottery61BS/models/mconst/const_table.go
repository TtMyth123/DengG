package mconst

const (
	TableName_LoAgentUser       = "lo_agent_user"
	TableName_LoAgentType       = "lo_agent_type"
	TableName_LoAgentBetNumData = "lo_agent_bet_num_data"
)

const (
	TableName_LoBetUser    = "lo_bet_user"
	TableName_LoBetType    = "lo_bet_type"
	TableName_LoBetBetData = "lo_bet_bet_data"
)

const (
	TableName_LoNumInfo = "lo_num_info"
	TableName_LoConfig  = "lo_config"

	TableName_LoInputBetM   = "lo_input_bet_m"
	TableName_LoLotteryType = "lo_lottery_type"
)

type a struct {
	LoBetBetData bool `json:"lo_bet_bet_data"`
	LoInputBetM  bool `json:"lo_input_bet_m"`
	LoConfig     bool `json:"lo_config"`
	LoNumInfo    bool `json:"lo_num_info"`

	LoAgentUser       bool `json:"lo_agent_user"`
	LoAgentType       bool `json:"lo_agent_type"`
	LoAgentBetNumData bool `json:"lo_agent_bet_num_data"`
}
