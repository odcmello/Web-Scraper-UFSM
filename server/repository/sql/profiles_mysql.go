package profile

import (
	"api-egressos/database"
	"api-egressos/model"
	"api-egressos/repository"
	"context"
	"log"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

func NewSQLProfileRepository() repository.ProfileRepo {
	return &pRepository{}
}

type pRepository struct {
	Db *sqlx.DB
}

var queryFilters = map[string]func(value interface{}, sb *sqlbuilder.SelectBuilder) string{
	"nome": func(value interface{}, sb *sqlbuilder.SelectBuilder) string {
		return sb.Like("nome", "%"+value.(string)+"%")
	},
	"empresa": func(value interface{}, sb *sqlbuilder.SelectBuilder) string {
		return sb.Like("company", "%"+value.(string)+"%")
	},
	"localizacao": func(value interface{}, sb *sqlbuilder.SelectBuilder) string {
		return sb.Like("location", "%"+value.(string)+"%")
	},
	"curso": func(value interface{}, sb *sqlbuilder.SelectBuilder) string {
		return sb.Equal("curso", value.(string))
	},
	"ano_evasao": func(value interface{}, sb *sqlbuilder.SelectBuilder) string {
		return sb.Equal("ano_evasao", value)
	},
}

func (p *pRepository) GetProfiles(ctx context.Context) ([]*model.Profile, error) {

	sb := sqlbuilder.NewSelectBuilder()
	// SELECT * FROM egressos
	sb.Select("*").From("egressos")
	// apply query filters
	filters := ctx.Value("filters").(map[string]interface{})
	applyFilters(filters, sb)

	// Build and exec query
	query, args := sb.Build()
	log.Println(query)

	rows, err := database.Db.Conn.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	profiles := make([]*model.Profile, 0)
	for rows.Next() {
		var profile model.Profile
		if err = rows.Scan(&profile.ID, &profile.Name, &profile.JobTitle, &profile.Company, &profile.Location,
			&profile.URL, &profile.Curso, &profile.AnoEvasao); err != nil {
			return nil, err
		}
		profiles = append(profiles, &profile)
	}

	return profiles, nil
}

func applyFilters(filters map[string]interface{}, sb *sqlbuilder.SelectBuilder) {
	if len(filters) > 0 {
		where := make([]string, 0)
		for k, v := range filters {
			filter := queryFilters[k](v, sb)
			where = append(where, filter)
		}
		sb.Where(sb.And(where...))
	}
}
