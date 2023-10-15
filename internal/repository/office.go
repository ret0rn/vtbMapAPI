package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/ret0rn/vtbMapAPI/internal/model"
)

func (r *Repository) GetOfficeLocation(ctx context.Context, filter *model.OfficeLocationFilter) (list model.OfficeLocationList, err error) {
	const maxDistanceKM = 3000
	const q = `SELECT office_id, longitude, latitude, distance, address, 
       		   office_name, timetable_individual, timetable_enterprise, metro_station, 
       		   has_ramp, client_types, handling_ids FROM (
			   SELECT ST_Distance(
				   ST_GeogFromText('SRID=4326;POINT(' || $1 || ' ' || $2 || ')'),
				   geom
			   ) distance, office_id, longitude, latitude, address, office_name, 
       		   coalesce(timetable_individual, '{}'::jsonb) timetable_individual,
			   coalesce(timetable_enterprise, '{}'::jsonb) timetable_enterprise,
			   metro_station, has_ramp, client_types, handling_ids
			   FROM public."office" WHERE is_active %s
			   ORDER BY location <-> point ('(' || $1 || ', ' || $2 || ')') LIMIT 10)
    		   office_location WHERE distance <= $3;`

	var addQ = ""
	var params = []interface{}{filter.Longitude, filter.Latitude, maxDistanceKM}
	if filter.ClientType != 0 && filter.HandlingType != 0 {
		params = append(params, filter.ClientType)
		addQ += fmt.Sprintf(" and $%d = any(client_types)", len(params))
		params = append(params, filter.HandlingType)
		addQ += fmt.Sprintf(" and $%d = any(handling_ids)", len(params))
	}

	ctxDb, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	rows, err := r.MasterPool.Query(ctxDb, fmt.Sprintf(q, addQ), params...) // TODO: заменить на SlavePool
	if err != nil && err != pgx.ErrNoRows {
		return nil, errors.Wrapf(err, "query error")
	}
	defer rows.Close()
	list = make([]*model.OfficeLocation, 0)
	for rows.Next() {
		var row model.OfficeLocation
		err := rows.Scan(
			&row.OfficeID,
			&row.Longitude,
			&row.Latitude,
			&row.Distance,
			&row.Address,
			&row.OfficeName,
			&row.TimetableIndividual,
			&row.TimetableEnterprise,
			&row.MetroStation,
			&row.HasRamp,
			&row.ClientTypes,
			&row.HandlingTypes,
		)
		if err != nil {
			return nil, errors.Wrapf(err, "scan error")
		}
		list = append(list, &row)
	}
	return list, rows.Err()
}

func (r *Repository) NewOffice(ctx context.Context, office *model.Office) error {
	const q = `INSERT INTO public.office
			(longitude, latitude, "location", geom, address, 
			 office_name, timetable_individual, timetable_enterprise, 
			 metro_station, handling_ids, client_types, has_ramp)
			VALUES($1, $2, '(%f, %f)', ST_SetSRID(ST_MakePoint($1, $2), 4326), $3, $4, $5, $6, $7, $8, $9, $10);`

	ctxDb, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	timetableIndividual, err := json.Marshal(office.TimetableIndividual.Days)
	if err != nil && err != pgx.ErrNoRows {
		return errors.Wrapf(err, "Marshal timetableIndividual error")
	}
	timetableEnterprise, err := json.Marshal(office.TimetableEnterprise.Days)
	if err != nil && err != pgx.ErrNoRows {
		return errors.Wrapf(err, "Marshal timetableEnterprise error")
	}
	_, err = r.MasterPool.Exec(ctxDb, fmt.Sprintf(q, office.Longitude, office.Latitude),
		office.Longitude,
		office.Latitude,
		office.Address,
		office.OfficeName,
		timetableIndividual,
		timetableEnterprise,
		office.MetroStation,
		pq.Array(office.HandlingTypes),
		pq.Array(office.ClientTypes),
		office.HasRamp,
	) // TODO: заменить на SlavePool
	if err != nil && err != pgx.ErrNoRows {
		return errors.Wrapf(err, "query error")
	}
	return nil
}
