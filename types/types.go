package types

type Vote struct {
	VoteName     string   `json:"VoteName" binding:"required"`
	VoteVariants []string `json:"VoteVariants" binding:"required"`
}

type UserVote struct {
	UserID     int `json:"userid" binding:"required"`
	VoteID     int `json:"voteid" binding:"required"`
	VariantNum int `json:"variantnum" binding:"required"`
}

type VoteInfo struct {
	VoteID int `json:"voteid" uri:"voteid" binding:"required"`
}
