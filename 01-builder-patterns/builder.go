package builder

/* ============== 理论部分：建造者模式四要素 (看最下面更常用的函数式选项模式) ============== */
// 标准的建造者可以把构建过程和最终表示分离，构建过程往往可以链式调用来构建Product的各个部分
// 可读性和拓展性都很高
//
// 1. Product，就是需要构建的复杂对象（属性很多~），在go中即对应的struct
// 2. Builder，是一个接口，定义了构建Product的各个部分的方法
// 3. Concrete Builder，实现了Builder接口，提供了具体的构建Product的方法
// 4. Director，负责调用Builder的一系列方法来构建Product，可以看成流水线
//
// product := NewProduct(a,b,c,d,...)
// 或
// product := &Product{}
// product.a = 1
// product.b = 2
// ...
// 当创建一个复杂对象的时候，很可能使用👆2种方式之一来创建，弊端也很明显：
// 第一种，参数太多，容易错乱，并且当某些参数是可选时，就需要多次调用构造函数来区分不同实例
// 第二种，如果漏了某个属性就无法正常工作
//

// 这就是需要创建的产品
type Vehicle struct {
	Wheels    int
	Seats     int
	Structure string
}

// 抽象建造者，提供构建产品的各个部分的方法
type Builder interface {
	SetWheels() Builder
	SetSeats() Builder
	SetStructure() Builder
	Build() Vehicle // 严格的情况下，只有调了 Builder 后才是一个成品诞生
}

// 具体建造者，在 go 中起内部组合了需要实现的产品
type Car struct {
	vehicle Vehicle
}

// 实现继承Builder
func (car *Car) SetWheels() Builder {
	car.vehicle.Wheels = 4
	return car
}
func (car *Car) SetSeats() Builder {
	car.vehicle.Seats = 4
	return car
}
func (car *Car) SetStructure() Builder {
	car.vehicle.Structure = "Car"
	return car
}
func (car *Car) Build() Vehicle {
	return car.vehicle
}

// 导演（指挥者）
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

// 直接对创建流程进行了完整编排，如果要创建 Bus、Bike等其他产品就可以快速开辟出另一个产线
func (director *Director) ConstructCar() {
	director.builder.SetWheels().SetSeats().SetStructure() //链式调用
}

/* ============== 实践：在 gin 开发中 ============== */
// Web 开发中，对象的构建过程往往不是固定的，而是**高度依赖于前端传来的动态参数**
// 一些可选参数，比如用户头像、年龄等，可以使用链式调用来优雅实现
// 1.往往会省略调 Builder 接口
// 2.director 的角色会由分层替代

/* ============== 补充：在go中函数式选项模式更优雅，也更常用 ============== */
type Server struct {
	Host    string
	Port    int
	Timeout int
}

type Option func(*Server) // 定义一个函数类型

func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}

func NewServer(opts ...Option) *Server {
	// 默认值
	s := &Server{Host: "localhost", Port: 8080}
	// 遍历应用选项
	for _, opt := range opts {
		opt(s)
	}
	return s
}

//调用：
// s := NewServer(WithPort(9000))
