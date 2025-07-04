# SnowflaskeId Snowflake ID

### Generate snowflake id
```go
// The following demonstrates the default configuration
// Usage Method One：
id := snowflakeid.NewInt64()
// Usage Method Two：
id := snowflakeid.Newstring()
// Usage Method Three：
sf := snowflakeid.New()
id := sf.NextId()  // Return an int64 snowflake id
```

## Overview

Snowflake ID Generator is a library used to generate distributed unique ids. It is based on Twitter's Snowflake algorithm and offers some additional features to adapt to different usage scenarios.

## Function list

### Create an instance

#### `NewInt64() int64`

Create and return a new snowflake ID (int64 type) using the default Settings.

#### `NewString() string`

Create and return a new snowflake ID (of string type) using the default Settings.

#### `New() *SnowflakeId`

Create a new instance of `SnowflakeId` and apply the default Settings.

### Set attributes

#### `SetMode(mode GenMode)`

Set the mode. Possible values include`Mode_Standard`, `Mode_IP`,And `Mode_PseudoRandom`。

#### `SetDataCenterId(dataCenterId int64) error`

Set the data center ID and ensure it is within the valid range.

#### `SetMachineId(machineId int64) error`

Set the machine ID and ensure it is within the valid range.

### Generate ID

#### `NextId() (int64, error)`

Generate the next ID based on the current configuration.

### Comparison method

#### `Before(t *SnowflakeId) bool`

Determine whether t is before s.

#### `After(t *SnowflakeId) bool`

Determine whether t is after s.

## Usage example

```go
// Create and initialize a SnowflakeId instance
snowflake := snowflakeid.New()

// Set the mode to IP mode
snowflake.SetMode(snowflakeid.Mode_IP)

// Set the data center ID to 1
err := snowflake.SetDataCenterId(1)
if err != nil {
	fmt.Println(err)
}

// Get the next ID
id, err := snowflake.NextId()
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Generated ID:", id)
}
```

Note: Please ensure that all dependencies have been imported correctly And related constants such as`maxDatacenterNum`, `maxMachineNum`, `maxSequenceNum`, `timestampShift`, `datacenterShift`, and `machineShift` Both `epoch` and `epoch` have been defined.