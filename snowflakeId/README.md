# SnowflaskeId 雪花ID

### 生成雪花id
```go
// 以下演示默认配置
// 使用方式一：
id := snowflakeid.NewInt64()
// 使用方式二：
id := snowflakeid.Newstring()
// 使用方式三：
sf := snowflakeid.New()
id := sf.NextId()  // 返回一个int64雪花id
```

## 概述

Snowflake ID Generator 是一个用于生成分布式唯一ID的库。它基于Twitter的Snowflake算法，并提供了一些额外的功能来适应不同的使用场景。

## 函数列表

### 创建实例

#### `NewInt64() int64`

创建并返回一个新的雪花ID（int64类型），使用默认设置。

#### `NewString() string`

创建并返回一个新的雪花ID（string类型），使用默认设置。

#### `New() *SnowflakeId`

创建一个新的`SnowflakeId`实例，并应用默认设置。

### 设置属性

#### `SetMode(mode GenMode)`

设置模式，可能的值包括`Mode_Standard`, `Mode_IP`, 和 `Mode_PseudoRandom`。

#### `SetDataCenterId(dataCenterId int64) error`

设置数据中心ID，确保其在有效范围内。

#### `SetMachineId(machineId int64) error`

设置机器ID，确保其在有效范围内。

### 生成ID

#### `NextId() (int64, error)`

生成下一个ID，基于当前配置。

### 比较方法

#### `Before(t *SnowflakeId) bool`

判断t是否在s之前。

#### `After(t *SnowflakeId) bool`

判断t是否在s之后。

## 使用示例

```go
// 创建并初始化SnowflakeId实例
snowflake := snowflakeid.New()

// 设置模式为IP模式
snowflake.SetMode(snowflakeid.Mode_IP)

// 设置数据中心ID为1
err := snowflake.SetDataCenterId(1)
if err != nil {
	fmt.Println(err)
}

// 获取下一个ID
id, err := snowflake.NextId()
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Generated ID:", id)
}
```

注意：请确保所有依赖项都已正确导入，并且相关常量如`maxDatacenterNum`, `maxMachineNum`, `maxSequenceNum`, `timestampShift`, `datacenterShift`, `machineShift`, 和 `epoch`都已经定义。
