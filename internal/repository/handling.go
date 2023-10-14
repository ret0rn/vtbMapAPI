package repository

import (
	"context"
	"fmt"

	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/ret0rn/vtbMapAPI/internal/model"
)

func (r *Repository) GetHandlingList(ctx context.Context, filter *model.HandlingFilter) (list []*model.Handling, err error) {
	const q = `SELECT handling_id, title, client_type, handling_duration, offices_ids FROM public.handling WHERE handling_id > 0 %s;`

	var addQ = ""
	var params = make([]interface{}, 0)
	if filter != nil {
		if filter.ClientType != 0 {
			params = append(params, filter.ClientType)
			addQ += fmt.Sprintf(" and client_type = $%d", len(params))
		}
	}
	ctxDB, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	rows, err := r.MasterPool.Query(ctxDB, fmt.Sprintf(q, addQ), params...) // TODO: заменить на Slave
	if err != nil {
		return nil, errors.Wrapf(err, "query error")
	}
	list = make([]*model.Handling, 0)
	for rows.Next() {
		var row model.Handling
		err = rows.Scan(&row.HandlingId, &row.Title, &row.ClientType, &row.HandlingDuration, pq.Array(&row.OfficesIDs))
		if err != nil {
			return nil, errors.Wrapf(err, "scan error")
		}
		list = append(list, &row)
	}
	return list, rows.Err()
}
