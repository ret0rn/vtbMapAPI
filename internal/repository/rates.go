package repository

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

func (r *Repository) GetRatesByOffice(ctx context.Context, officeIds []int64) (ratesMap map[int64]float64, err error) {
	const q = `SELECT office_id, rating FROM public."office_rating" 
            	WHERE office_id = ANY($1)`

	ctxDb, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()
	rows, err := r.MasterPool.Query(ctxDb, q, pq.Array(officeIds))
	if err != nil && err != pgx.ErrNoRows {
		return nil, errors.Wrapf(err, "query error")
	}
	defer rows.Close()

	ratesMap = make(map[int64]float64, 0)
	for rows.Next() {
		var officeId int64
		var rating float64
		err = rows.Scan(&officeId, &rating)
		if err != nil {
			return nil, errors.Wrapf(err, "scan error")
		}
		ratesMap[officeId] = rating
	}
	return ratesMap, rows.Err()
}
