package queries

import (
	"time"

	"github.com/ha-dev/getslowestdbqueries/models"
	errors "github.com/ha-dev/getslowestdbqueries/pkg/err"
	"gorm.io/gorm"
)

// DbQueries struct for queries from Queries model.
type DbQueries struct {
	*gorm.DB
}

// GetDbQueries method for getting all Slowest Queries.
/* func (q *DbQueries) GetDbQueries() ([]models.ResulyQuery, error) {
	// Define books variable.
	resultQuery := []models.ResulyQuery{}
	q.DB.Table("querys").Scan(&resultQuery)
	// Return query result.
	return resultQuery, nil
}
*/

func (q *DbQueries) GetDbQueries(query_command string, limit uint32, offset uint32) (interface{}, error) {

	resultQuery := []models.ResulyQueryWithPagination{}
	//
	query := `SELECT
		COUNT(*) OVER() AS "totalrecord",
		id,
		query,
		time_to_run
	FROM querys
	WHERE query LIKE ? ORDER BY time_to_run ASC limit ? offset ?`

	q.Raw(query, "%"+query_command+"%", limit, offset).Scan(&resultQuery)
	data := make(map[string]interface{})
	data["list_slowest_query"] = resultQuery

	if len(resultQuery) != 0 {
		data["count"] = resultQuery[0].Totalrecord
	} else {
		data["count"] = 0
	}

	return data, nil
}

// InsertDbQueries method for Insert Slowest Query In Db.
func (q *DbQueries) InsertDbQueries(slowest_query_info *models.Querys) error {
	query := models.Querys{}
	query.Query = slowest_query_info.Query
	query.Time_to_run = slowest_query_info.Time_to_run
	err := q.Save(&query).Error
	if err != nil {
		return errors.ErrorInsertSlowestQuery
	}
	return nil
}

// Run Query and Get time to Run this Query on database
func (q *DbQueries) ExecuteDbQueries(query string) (int64, error) {
	start := time.Now()
	err := q.Exec(query).Error
	if err != nil {
		return time.Since(start).Milliseconds(), err
	}
	elapsed := time.Since(start).Milliseconds()
	return elapsed, nil
}
