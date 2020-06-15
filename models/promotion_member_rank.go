package models

type PromotionMemberRank struct {
	Promotions  *Promotion  `orm:"column(promotions);rel(fk)"`
	MemberRanks *MemberRank `orm:"column(member_ranks);rel(fk)"`
}
