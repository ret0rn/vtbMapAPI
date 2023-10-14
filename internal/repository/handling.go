package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/ret0rn/vtbMapAPI/internal/model"
)

func (r *Repository) GetHandlingList(ctx context.Context, filter *model.HandlingFilter) (list model.HandlingList, err error) {
	const q = `SELECT handling_id, title, client_type, handling_duration FROM public.handling WHERE handling_id > 0 %s;`

	var addQ = ""
	var params = make([]interface{}, 0)
	if filter != nil {
		if filter.ClientType != 0 {
			params = append(params, filter.ClientType)
			addQ += fmt.Sprintf(" and client_type = $%d", len(params))
		}
		if len(filter.HandlingIds) > 0 {
			params = append(params, pq.Array(filter.HandlingIds))
			addQ += fmt.Sprintf(" and handling_id = $%d", len(params))
		}
	}
	ctxDB, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	rows, err := r.MasterPool.Query(ctxDB, fmt.Sprintf(q, addQ), params...) // TODO: заменить на SlavePool
	if err != nil && err != pgx.ErrNoRows {
		return nil, errors.Wrapf(err, "query error")
	}
	defer rows.Close()
	list = make([]*model.Handling, 0)
	for rows.Next() {
		var row model.Handling
		err = rows.Scan(&row.HandlingId, &row.Title, &row.ClientType, &row.HandlingDuration)
		if err != nil {
			return nil, errors.Wrapf(err, "scan error")
		}
		list = append(list, &row)
	}
	return list, rows.Err()
}
