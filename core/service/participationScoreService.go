package service

import (
	"github.com/zoobc/zoobc-core/common/blocker"
	"github.com/zoobc/zoobc-core/common/model"
	"github.com/zoobc/zoobc-core/common/query"
)

type (
	ParticipationScoreServiceInterface interface {
		GetParticipationScore(nodePublicKey []byte) (int64, error)
	}

	ParticipationScoreService struct {
		ParticipationScoreQuery query.ParticipationScoreQueryInterface
		QueryExecutor           query.ExecutorInterface
	}
)

func NewParticipationScoreService(
	participationScoreQuery query.ParticipationScoreQueryInterface,
	queryExecutor query.ExecutorInterface,
) *ParticipationScoreService {
	return &ParticipationScoreService{
		ParticipationScoreQuery: participationScoreQuery,
		QueryExecutor:           queryExecutor,
	}
}

// GetParticipationScore handle received block from another node
func (pss *ParticipationScoreService) GetParticipationScore(nodePublicKey []byte) (int64, error) {
	var (
		participationScores []*model.ParticipationScore
	)
	participationScoreQ, args := pss.ParticipationScoreQuery.GetParticipationScoreByNodePublicKey(nodePublicKey)
	rows, err := pss.QueryExecutor.ExecuteSelect(participationScoreQ, false, args...)
	if err != nil {
		return 0, blocker.NewBlocker(blocker.DBErr, err.Error())
	}
	defer rows.Close()
	participationScores, err = pss.ParticipationScoreQuery.BuildModel(participationScores, rows)
	// if there aren't participation scores for this address/node, return 0
	if (err != nil) || len(participationScores) == 0 {
		return 0, nil
	}
	return participationScores[0].Score, nil
}
