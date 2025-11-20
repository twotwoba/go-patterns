package bridge

import "fmt"

/* ============== 理论 ============== */
// 桥接模式，目的是避免类的爆炸性增长
// 经典比喻：开关与电器**
// *   我们有不同类型的**开关**：墙壁开关、拉绳开关、遥控开关。这是一个变化的维度（**抽象部分**）。
// *   我们有不同类型的**电器**：电灯、风扇、电视机。这是另一个独立变化的维度（**实现部分**）。

// **如果不使用桥接模式**，我们可能会用继承来创建所有的组合：
// *   `WallSwitchForLight`, `WallSwitchForFan`, `WallSwitchForTV`
// *   `PullCordSwitchForLight`, `PullCordSwitchForFan`, `PullCordSwitchForTV`
// *   `RemoteControlForLight`, `RemoteControlForFan`, `RemoteControlForTV`
// ...
// 你看，3 种开关 x 3 种电器 = 9 个类。如果再增加一种“声控开关”，就得再加 3 个类。这就是**类的数量爆炸**。

/*
	分离抽象部分和实现部分， 桥接设计的核心在于抽象接口和组合抽象接口的结构体
	设计思想：
		1. 抽象接口，(实现该接口的具体struct，可扩展多个struct)
		2. 属性为抽象接口的struct Phone（桥接层）
		3. 与Phone组合的具体struct（可以是多个struct）
*/

// 抽象接口
type SoftWare interface {
	Run()
}

// 具体类型CPU和Storage
type Cpu struct{}

func (c *Cpu) Run() {
	fmt.Println("this is cpu run")
}

type Storage struct{}

func (s *Storage) Run() {
	fmt.Println("this is storage run")
}

// 桥接层结构体  -- 这里是桥接模式的核心
// 在下面具体品牌和 Phone 连接，Phone 和 software 连接
type Phone struct {
	software SoftWare
}

// 赋值具体software
func (s *Phone) SetSoftWare(soft SoftWare) {
	s.software = soft
}

// Apple 结构体
type Apple struct {
	phone Phone
}

func (p *Apple) SetShape(soft SoftWare) {
	p.phone.SetSoftWare(soft)
}

func (p *Apple) Print() {
	p.phone.software.Run()
}

// HuaWei结构体
type HuaWei struct {
	phone Phone
}

func (p *HuaWei) SetShape(soft SoftWare) {
	p.phone.SetSoftWare(soft)
}

func (p *HuaWei) Print() {
	p.phone.software.Run()
}

/**
  // 1. 创建具体的“实现”
  cpuSoftware := &bridge.Cpu{}
  storageSoftware := &bridge.Storage{}

  // 2. 创建具体的“抽象”
  applePhone := &bridge.Apple{}

  // 3. 将“实现”和“抽象”桥接起来
  applePhone.SetShape(cpuSoftware) // 让苹果手机运行 CPU 软件
  applePhone.Print() // 输出: this is cpu run

  // 4. 在运行时可以轻松改变桥接的另一端
  applePhone.SetShape(storageSoftware) // 现在让同一部苹果手机运行 Storage 软件
  applePhone.Print() // 输出: this is storage run


  将**手机品牌**和**软件**这两个独立变化的维度分离开来。
  *   通过 `Phone` 这个中间层结构体作为**桥梁**，连接了这两个维度。
  *   使得我们可以在不修改 `Apple` 或 `HuaWei` 代码的情况下，任意扩展新的 `SoftWare`；也可以在不修改 `Cpu` 或 `Storage` 的情况下，任意扩展新的手机品牌。

  这正是桥接模式的精髓所在，它是一种应对多维度变化的强大设计武器
*/
