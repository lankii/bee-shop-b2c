package common

// 等待付款 0
// pendingPayment

// 等待接单中 1
// pendingReview

// 等待发货 2
// pendingShipment

// 已发货 3
// shipped

// 已收货 4
// received

// 已完成 5
// completed

// 已失败 6
// failed

// 已取消 7
// canceled

// 已拒绝 8
// denied

// 未完成 9
// unfinished

// 退款中 10
// refunding

// 退款完成 11
// refunded

const (
	pendingPayment = iota
	pendingReview
	pendingShipment
	shipped
	received
	completed
	failed
	canceled
	denied
	unfinished
	refunding
	refunded
)

func GetOrderStatusName(status int) string {
	switch status {
	case pendingPayment:
		return "等待付款"
	case pendingReview:
		return "等待接单中"
	case pendingShipment:
		return "等待发货"
	case shipped:
		return "已发货"
	case received:
		return "已收货"
	case completed:
		return "已完成"
	case failed:
		return "已失败"
	case canceled:
		return "已取消"
	case denied:
		return "已拒绝"
	case unfinished:
		return "未完成"
	case refunding:
		return "退款中"
	case refunded:
		return "退款完成"
	default:
		return ""
	}
}
