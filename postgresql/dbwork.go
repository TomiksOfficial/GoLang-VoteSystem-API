package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"strconv"
	"voteAPI/types"
)

func (conn *PGDatabase) CreateVote(ctx context.Context, vote *types.Vote) error {

	if conn == nil {
		//fmt.Println("DB IS NULL")
		return fmt.Errorf("DB IS NULL")
	}

	query := `INSERT INTO voteapi_vote("voteName") VALUES(@voteName) RETURNING id;`
	args := pgx.NamedArgs{
		"voteName": vote.VoteName,
	}

	tx, err := conn.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	row := tx.QueryRow(ctx, query, args)

	var id int
	err = row.Scan(&id)
	if err != nil {
		return err
	}

	query = `	INSERT INTO voteapi_vote_variants("voteVariant", "voteId", "variantNum") 
				VALUES(@voteVariant, @voteId, @variantNum);`

	for i, variant := range vote.VoteVariants {
		args = pgx.NamedArgs{
			"voteVariant": variant,
			"voteId":      id,
			"variantNum":  i + 1,
		}

		_, err = tx.Exec(ctx, query, args)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (conn *PGDatabase) DoVote(ctx context.Context, vote *types.UserVote) error {
	if conn == nil {
		return fmt.Errorf("DB IS NULL")
	}

	query := `	INSERT INTO voteapi_users_vote(userid, voteid, "variantNum") 
				VALUES(@userId, @voteId, @variantNum) ON CONFLICT DO NOTHING;`

	args := pgx.NamedArgs{
		"userId":     vote.UserID,
		"voteId":     vote.VoteID,
		"variantNum": vote.VariantNum,
	}

	_, err := conn.db.Exec(ctx, query, args)
	if err != nil {
		return err
	}

	return nil
}

func (conn *PGDatabase) GetCountVotes(ctx context.Context, voteinfo *types.VoteInfo) (map[string]any, error) {
	if conn == nil {
		return nil, fmt.Errorf("DB IS NULL")
	}

	query := ` SELECT COUNT(id) AS count_votes, "variantNum" FROM voteapi_users_vote GROUP BY "variantNum", voteid HAVING voteid = @voteId;`

	args := pgx.NamedArgs{
		"voteId": voteinfo.VoteID,
	}

	rows, err := conn.db.Query(ctx, query, args)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var countVotes int
	var variantNum int
	result := make(map[string]any)

	for rows.Next() {
		err = rows.Scan(&countVotes, &variantNum)
		if err != nil {
			return nil, err
		}

		result[strconv.Itoa(variantNum)] = countVotes
	}

	return result, nil
}
