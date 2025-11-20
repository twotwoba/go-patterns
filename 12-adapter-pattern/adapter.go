package adapter

import "fmt"

/*
	设计思想:
		1.目标接口（示例中的Player）
		2.被适配者
		3.核心是通过适配器Adapter转换为目标接口（组合的方式包含被适配者）

		If the Target and Adaptee has similarities then the adapter has just to delegate
		the requests from the Target to the Adaptee.
		If Target and Adaptee are not similar, then the adapter might have to convert the
		data structures between those and to implement the operations required by the Target
		but not implemented by the Adaptee
*/
//目标接口- 音乐播放器
type Player interface {
	PlayMusic()
}

type MusicPlayer struct {
	Src string
}

func (music *MusicPlayer) PlayMusic() {
	fmt.Println("play music: " + music.Src)
}

// 播放器的使用方法，只接受 Player 类型的播放器
func Play(player Player) {
	player.PlayMusic()
}

// 被适配者-在音乐播放基础上实现游戏播放
type GamePlayer struct {
	Src string
}

func (game *GamePlayer) PlaySound() {
	fmt.Println("play sound: " + game.Src)
}

// 这里要实现调用play方法的时候，实现GamePlayer的播放
// 通过组合的方式，声明一个GamePlayer的适配器
type GamePlayerAdapter struct {
	Adapt GamePlayer
}

// 继承Player接口, 内部调用GamePlayer的方法
func (game *GamePlayerAdapter) PlayMusic() {
	game.Adapt.PlaySound()
}
