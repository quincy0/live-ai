package douyinService

type Douyin struct {
	cookies string
}

type Redpacket struct {
	RedpacketData *RedpacketData `json:"redpacket_data"`
}

type RedpacketData struct {
	RedpacketActivity *RedpacketActivity `json:"redpacket_activity"`
	RedpacketMetaList []*RedpacketMeta   `json:"redpacket_meta_list"`
}

type RedpacketActivity struct {
	ActivityBizType            int    `json:"activity_biz_type"`
	RedpacketActivityName      string `json:"redpacket_activity_name"`
	MaxApplyTimes              int    `json:"max_apply_times"`
	ValidityType               int    `json:"validity_type"`
	LiveRedpackActivitySubType int    `json:"live_redpack_activity_sub_type"`
	KolUserTag                 int    `json:"kol_user_tag"`
	RedpacketSubType           int    `json:"redpacket_sub_type"`
	RedpackType                int    `json:"redpack_type"`
	ValidStartTime             int64  `json:"valid_start_time"`
	ValidEndTime               int64  `json:"valid_end_time"`
	TotalCredit                int    `json:"total_credit"`
}

type RedpacketMeta struct {
	TotalAmount int        `json:"total_amount"`
	TotalCredit int        `json:"total_credit"`
	CreditType  int        `json:"credit_type"`
	ExtraInfo   *ExtraInfo `json:"extra_info"`
	AvgCredit   int        `json:"avg_credit"`
}

type ExtraInfo struct {
	StrategyGoal int `json:"strategy_goal"`
}

func InitRedpacket() *Redpacket {
	return &Redpacket{
		RedpacketData: &RedpacketData{
			RedpacketActivity: &RedpacketActivity{
				ActivityBizType:            1,
				RedpacketActivityName:      "0619",
				MaxApplyTimes:              1,
				ValidityType:               1,
				LiveRedpackActivitySubType: 3,
				KolUserTag:                 0,
				RedpacketSubType:           1,
				RedpackType:                11,
				ValidStartTime:             1718801453,
				ValidEndTime:               1718805053,
				TotalCredit:                10000,
			},
			RedpacketMetaList: []*RedpacketMeta{
				&RedpacketMeta{
					TotalAmount: 20,
					TotalCredit: 10000,
					CreditType:  1,
					ExtraInfo: &ExtraInfo{
						StrategyGoal: 0,
					},
					AvgCredit: 500,
				},
			},
		},
	}
}

type CreateRedPacketResp struct {
	St    int    `json:"st"`
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Extra struct {
		Now   int64  `json:"now"`
		LogId string `json:"log_id"`
	} `json:"extra"`
	Data struct {
		RedpacketActivityId string   `json:"redpacket_activity_id"`
		RedpacketMetaIds    []string `json:"redpacket_meta_ids"`
	} `json:"data"`
}

type DisplayResp struct {
	St    int    `json:"st"`
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Extra struct {
		LogId string `json:"log_id"`
		Now   int64  `json:"now"`
	} `json:"extra"`
	Data struct {
	} `json:"data"`
}

type CheckDisplayTimeResp struct {
	St    int    `json:"st"`
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Extra struct {
		LogId string `json:"log_id"`
		Now   int64  `json:"now"`
	} `json:"extra"`
	Data struct {
		Result int    `json:"result"`
		Desc   string `json:"desc"`
	} `json:"data"`
}

/*

{
  "redpacket_data": {
    "redpacket_activity": {
      "activity_biz_type": 1,
      "redpacket_activity_name": "0619",
      "max_apply_times": 1,
      "validity_type": 1,
      "live_redpack_activity_sub_type": 3,
      "kol_user_tag": 0,
      "redpacket_sub_type": 1,
      "redpack_type": 11,
      "valid_start_time": 1718801453,
      "valid_end_time": 1718805053,
      "total_credit": 10000
    },
    "redpacket_meta_list": [
      {
        "total_amount": 20,
        "total_credit": 10000,
        "credit_type": 1,
        "extra_info": {
          "strategy_goal": 0
        },
        "avg_credit": 500
      }
    ]
  }
}

*/
