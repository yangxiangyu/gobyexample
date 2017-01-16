package cmd

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/smallnest/gofsm"
	"github.com/spf13/cobra"
)

// fsmCmd represents the fsm command
var fsmCmd = &cobra.Command{
	Use:   "fsm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		TFSM()
		fmt.Println("fsm called")
	},
}

type Turnstile struct {
	ID         uint64
	EventCount uint64   //事件统计
	CoinCount  uint64   //投币事件统计
	PassCount  uint64   //顾客通过事件统计
	State      string   //当前状态
	States     []string //历史经过的状态
}

type TurnstileEventProcessor struct{}

func (p *TurnstileEventProcessor) OnExit(fromState string, args []interface{}) {
	t := args[0].(*Turnstile)
	if t.State != fromState {
		panic(fmt.Errorf("转门 %v 的状态与期望的状态 %s 不一致，可能在状态机外被改变了", t, fromState))
	}

	log.Infof("转门 %d 从状态 %s 改变", t.ID, fromState)
}

func (p *TurnstileEventProcessor) Action(action string, fromState string, toState string, args []interface{}) {
	t := args[0].(*Turnstile)
	t.EventCount++

	switch action {
	case "pass": //用户通过的action
		t.PassCount++
	case "check", "repeat-check": //刷卡或者投币的action
		t.CoinCount++
	default: //其它action
	}
}

func (p *TurnstileEventProcessor) OnEnter(toState string, args []interface{}) {
	t := args[0].(*Turnstile)
	t.State = toState
	t.States = append(t.States, toState)

	log.Infof("转门 %d 的状态改变为 %s ", t.ID, toState)
}

func initFSM() *fsm.StateMachine {
	delegate := &fsm.DefaultDelegate{P: &TurnstileEventProcessor{}}

	transitions := []fsm.Transition{
		fsm.Transition{From: "Locked", Event: "Coin", To: "Unlocked", Action: "check"},
		fsm.Transition{From: "Locked", Event: "Push", To: "Locked", Action: "invalid-push"},
		fsm.Transition{From: "Unlocked", Event: "Push", To: "Locked", Action: "pass"},
		fsm.Transition{From: "Unlocked", Event: "Coin", To: "Unlocked", Action: "repeat-check"},
	}

	return fsm.NewStateMachine(delegate, transitions...)
}

func TFSM() {
	ts := &Turnstile{
		ID:     1,
		State:  "Locked",
		States: []string{"Locked"},
	}
	fsm := initFSM()

	//推门
	//没刷卡/投币不可进入
	fsm.Trigger(ts.State, "Push", ts)

	//推门
	//没刷卡/投币不可进入
	fsm.Trigger(ts.State, "Push", ts)

	//刷卡或者投币
	//不容易啊，终于解锁了
	fsm.Trigger(ts.State, "Coin", ts)

	//刷卡或者投币
	//这时才解锁
	fsm.Trigger(ts.State, "Coin", ts)

	//推门
	//这时才能进入，进入后闸门被锁
	fsm.Trigger(ts.State, "Push", ts)

	//推门
	//无法进入，闸门已锁
	fsm.Trigger(ts.State, "Push", ts)

}

func init() {
	RootCmd.AddCommand(fsmCmd)

}