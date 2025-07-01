# SnowflaskeId 雪花ID

### 生成雪花id
```go
sf := snowflakeid.New()
id := sf.NextId()  // 返回一个int64雪花id
```