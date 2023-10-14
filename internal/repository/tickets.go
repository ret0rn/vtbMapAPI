package repository

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

func (r *Repository) GetCountPeopleByOffice(ctx context.Context, officeIds []int64, handlingType int64) (workloadMap map[int64]int64, err error) {
	const q = `SELECT office_id, count(*) FROM public."queue_tickets" qt 
               WHERE office_id = ANY($1) AND handling_id = $2 GROUP BY office_id`

	ctxDb, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	rows, err := r.MasterPool.Query(ctxDb, q, pq.Array(officeIds), handlingType)
	if err != nil && err != pgx.ErrNoRows {
		return nil, errors.Wrapf(err, "query error")
	}
	defer rows.Close()
	workloadMap = make(map[int64]int64, 0)
	for rows.Next() {
		var officeId, countPeople int64
		err := rows.Scan(&officeId, &countPeople)
		if err != nil {
			return nil, errors.Wrapf(err, "scan error")
		}
		workloadMap[officeId] = countPeople
	}
	return workloadMap, rows.Err()
}
