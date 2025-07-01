package snowflakeid

type GenMode int

// tools 工具

const (
	Mode_Standard     = iota // 标准模式 Standard (64位整数:| 1 Bit | 41 Bits 时间戳 | 5 Bits 数据中心ID | 5 Bits 机器ID | 12 Bits 序列号 |)
	Mode_IP                  // IP模式 IP (64位整数:| 1 Bit | 41 Bits 时间戳 | 5 Bits 数据中心ID | 5 Bits IP地址截取 | 12 Bits 序列号 |)
	Mode_PseudoRandom        // 伪随机模式 PseudoRandom (64位整数:| 1 Bit | 41 Bits 时间戳 | 5 Bits 数据中心ID | 5 Bits 伪随机整数 | 12 Bits 序列号 |)
)

// 业务用

const (
	timestampBits  = 41
	datacenterBits = 5
	machineBits    = 5
	sequenceBits   = 12

	maxDatacenterNum = -1 ^ (-1 << datacenterBits) // 最大数据中心id
	maxMachineNum    = -1 ^ (-1 << machineBits)    // 最大机器id
	maxSequenceNum   = -1 ^ (-1 << sequenceBits)   // 最大序列号

	datacenterShift = sequenceBits + machineBits + datacenterBits
	machineShift    = sequenceBits + machineBits
	timestampShift  = sequenceBits + machineBits + datacenterBits

	defaultMachineId = 1 // 默认机器ID
)

var (
	epoch int64 = 1735660800000 // 自定义起始时间戳（毫秒）
)
