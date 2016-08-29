package cvm

// InstanceStatus represents instance status, refer https://www.qcloud.com/doc/api/229/831
type InstanceStatus int

const (
	// Failure 故障
	Failure = InstanceStatus(1)
	// Running 运行中
	Running = InstanceStatus(2)
	// Creating 创建中
	Creating = InstanceStatus(3)
	// Stopped 已关机
	Stopped = InstanceStatus(4)
	// Refund 已退还
	Refund = InstanceStatus(5)
	// Refunding 退还中
	Refunding = InstanceStatus(6)
	// Rebooting 重启中
	Rebooting = InstanceStatus(7)
	// Starting 开机中
	Starting = InstanceStatus(8)
	// Stopping 关机中
	Stopping = InstanceStatus(9)
	// PasswordReseting 密码重置中
	PasswordReseting = InstanceStatus(10)
	// Formatting 格式化中
	Formatting = InstanceStatus(11)
	// ImageBuilding 镜像制作中
	ImageBuilding = InstanceStatus(12)
	// BandwidthSetting 宽带设置中
	BandwidthSetting = InstanceStatus(13)
	// OsInstalling 重装系统中
	OsInstalling = InstanceStatus(14)
	// DomainBinding 域名绑定中
	DomainBinding = InstanceStatus(15)
	// DomainUnbinding 域名解绑中
	DomainUnbinding = InstanceStatus(16)
	// LoadBalanceBinding 负载均衡绑定中
	LoadBalanceBinding = InstanceStatus(17)
	// LoadBalanceUnbinding 负载均衡解绑中
	LoadBalanceUnbinding = InstanceStatus(18)
	// Upgrading 升级中
	Upgrading = InstanceStatus(19)
	// KeyDelivering 密钥下发中
	KeyDelivering = InstanceStatus(20)
	// Maintaining 维护中
	Maintaining = InstanceStatus(21)
)
