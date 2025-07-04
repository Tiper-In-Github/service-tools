package snowflakeid

import (
	"fmt"
	"math/rand"
	"net"
	"service-tools/utils"
	"strconv"
	"time"
)

// SnowflakeId 实例
type SnowflakeId struct {
	mode          GenMode // 模式
	dataCenterId  int64   // 数据中心ID
	machineId     int64   // 机器ID
	sequence      int64   // 当前毫秒内序列号
	lastTimestamp int64   // 上次生成ID的时间戳

	value int64
}

// NewInt64 创建int64雪花ID实例并应用默认设置
func NewInt64() int64 {
	id, _ := New().NextId()
	return id
}

// NewString 创建string雪花ID实例并应用默认设置
func NewString() string {
	id, _ := New().NextId()
	return strconv.Itoa(int(id))
}

// New 创建雪花ID实例并应用默认设置
func New() *SnowflakeId {
	// Go 1.20+ 随机数发生器已弃用
	// rand.Seed(time.Now().UnixNano()) // 初始化随机种子
	return &SnowflakeId{mode: Mode_Standard, dataCenterId: 0, machineId: defaultMachineId, sequence: 0}
}

func ipToMachineId(ip net.IP) int64 {
	// 确保处理的是IPv4地址
	ip = ip.To4()
	if ip == nil {
		return -1 // 返回-1表示错误或不适用
	}

	// 使用最后一个字节并缩小范围
	return int64(ip[3] % 32) // 32对应于maxMachineNum+1
}

// SetMode 设置模式
//
// 注意:这会覆盖s.SetMachineId的配置
func (s *SnowflakeId) SetMode(mode GenMode) {
	s.mode = mode

	switch s.mode {
	case Mode_Standard:
		s.SetMachineId(defaultMachineId)
	case Mode_IP:
		var a int64 = 7
		ip, err := utils.GetLocalIP()
		if err == nil {
			a = ipToMachineId(ip)
		}
		s.SetMachineId(a)
	case Mode_PseudoRandom:
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		s.SetMachineId(r.Int63n(maxMachineNum + 1))
	}
}

// SetDataCenterId 设置数据中心ID
//
// 如果设置失败，将会使用默认值
func (s *SnowflakeId) SetDataCenterId(dataCenterId int64) error {
	if dataCenterId > maxDatacenterNum || dataCenterId < 0 {
		return fmt.Errorf("dataCenterId can't be greater than %d or less than 0", maxDatacenterNum)
	}
	s.dataCenterId = dataCenterId
	return nil
}

// SetMachineId 设置机器ID
//
// 如果设置失败，将会使用默认值
func (s *SnowflakeId) SetMachineId(machineId int64) error {
	if machineId > maxMachineNum || machineId < 0 {
		return fmt.Errorf("machineId can't be greater than %d or less than 0", maxMachineNum)
	}
	s.machineId = machineId
	return nil
}

// NextId 生成ID
func (s *SnowflakeId) NextId() (int64, error) {

	switch s.mode {
	case Mode_Standard:
		if s.machineId > maxMachineNum || s.machineId < 0 {
			return 0, fmt.Errorf("machineId can't be greater than %d or less than 0", maxMachineNum)
		}
	}

	// 雪花算法生成
	timestamp := time.Now().UnixNano() / 1e6

	if timestamp < s.lastTimestamp {
		return 0, fmt.Errorf("clock moved backwards. Refusing to generate id for %d milliseconds", s.lastTimestamp-timestamp)
	}

	if timestamp == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & maxSequenceNum
		if s.sequence == 0 {
			for timestamp <= s.lastTimestamp {
				timestamp = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.sequence = 0
	}

	s.lastTimestamp = timestamp

	id := ((timestamp - epoch) << timestampShift) |
		(s.dataCenterId << datacenterShift) |
		(s.machineId << machineShift) |
		s.sequence

	s.value = id
	return id, nil
}

// Before 返回t是否在s之前
//
// 注意:此方法仅适用于同一机器下同一模式的ID
func (s *SnowflakeId) Before(t *SnowflakeId) bool {
	a := s.value & maxSequenceNum
	b := t.value & maxSequenceNum
	return a < b
}

// After 返回t是否在s之后
//
// 注意:此方法仅适用于同一机器下同一模式的ID
func (s *SnowflakeId) After(t *SnowflakeId) bool {
	a := s.value & maxSequenceNum
	b := t.value & maxSequenceNum
	return a > b
}
