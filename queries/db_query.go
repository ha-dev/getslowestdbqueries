package queries

import (
	"time"

	"github.com/ha-dev/getslowestdbqueries/models"
	errors "github.com/ha-dev/getslowestdbqueries/pkg/err"
	"gorm.io/gorm"
)

// BookQueries struct for queries from Book model.
type DbQueries struct {
	*gorm.DB
}

// GetDbQueries method for getting all Slowest Queries.
func (q *DbQueries) GetDbQueries() ([]models.ResulyQuery, error) {
	// Define books variable.
	resultQuery := []models.ResulyQuery{}
	q.DB.Table("querys").Scan(&resultQuery)
	// Return query result.
	return resultQuery, nil
}

// InsertDbQueries method for Insert Slowest Query In Db.
func (q *DbQueries) InsertDbQueries(slowest_query_info *models.Querys) error {
	// Define books variable.
	query := models.Querys{}
	query.Query = slowest_query_info.Query
	query.Time_to_run = slowest_query_info.Time_to_run
	err := q.Save(&query).Error
	if err != nil {
		return errors.ErrorInsertSlowestQuery
	}

	return nil
}

func (q *DbQueries) ExecuteDbQueries(query string) (int64, error) {

	start := time.Now()
	err := q.Exec(query).Error
	if err != nil {
		return time.Since(start).Milliseconds(), err
	}
	elapsed := time.Since(start).Milliseconds()

	return elapsed, nil
}
