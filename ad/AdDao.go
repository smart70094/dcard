package ad

import (
	"database/sql"
	"fmt"
)

func addAd(tx *sql.Tx, ad Ad) int {
	sqlStatement := `
		INSERT INTO ad (title, "startAt", "endAt")
		VALUES ($1, $2, $3)
		RETURNING id`

	var id int
	err := tx.QueryRow(sqlStatement, ad.Title, ad.StartAt, ad.EndAt).Scan(&id)
	if err != nil {
		return 0
	}

	return id
}

func retrieveAd(vo GetAdReqVo, db *sql.DB) ([]map[string]interface{}, error) {
	sqlStatement := `SELECT hot_ad.id, hot_ad.title, hot_ad."startAt", hot_ad."endAt"
	FROM hot_ad
	JOIN hot_ad_condition ON hot_ad.id = hot_ad_condition.hot_ad_id
	WHERE hot_ad_condition.age_start <= $1 AND hot_ad_condition.age_end >= $2`
	params := []interface{}{vo.Age, vo.Age}

	if vo.Gender != "" {
		sqlStatement += fmt.Sprintf(" AND $%d = ANY(hot_ad_condition.gender)", len(params)+1)
		params = append(params, vo.Gender)
	}
	if vo.Country != "" {
		sqlStatement += fmt.Sprintf(" AND $%d = ANY(hot_ad_condition.country)", len(params)+1)
		params = append(params, vo.Country)
	}
	if vo.Platform != "" {
		sqlStatement += fmt.Sprintf(" AND $%d = ANY(hot_ad_condition.platform)", len(params)+1)
		params = append(params, vo.Platform)
	}

	sqlStatement += fmt.Sprintf(" OFFSET $%d LIMIT $%d", len(params)+1, len(params)+2)
	params = append(params, vo.Offset, vo.Limit)

	rows, err := db.Query(sqlStatement, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	results := make([]map[string]interface{}, 0)
	for rows.Next() {
		columnValues := make([]interface{}, len(columns))
		columnPointers := make([]interface{}, len(columns))
		for i := range columnValues {
			columnPointers[i] = &columnValues[i]
		}
		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}
		item := make(map[string]interface{})
		for i, column := range columns {
			item[column] = columnValues[i]
		}
		results = append(results, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
