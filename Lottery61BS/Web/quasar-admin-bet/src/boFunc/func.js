const arrDrawMoneyStateName = ['','申请状态','下分','下分完成','已汇款']
export function DrawMoneyStateName(state){
  return arrDrawMoneyStateName[state]
}

const arrSaveMoneyName = ['','申请状态','付款中','已付款','已上传凭证', '充值成功']
export function SaveMoneyStateName(state){
  return arrSaveMoneyName[state]
}


const arrBetStateName = ['','待开奖','已兑奖','不开奖退款']
export function BetStateName(state){
  return arrBetStateName[state]
}

const arrUserStateName = ['新加','启用','禁用']
export function UserStateName(state){
  return arrUserStateName[state]
}

const arrSysUserStateName = ['新加','启用','禁用']
export function SysUserStateName(state){
  return arrSysUserStateName[state]
}
