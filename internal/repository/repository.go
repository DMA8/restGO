package repository

import (
	"context"
	"errors"
	"time"

	"testTask/internal/domain"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	dbConn *pgxpool.Pool
}

func NewRepository(dbConn *pgxpool.Pool) *Repository {
	return &Repository{
		dbConn: dbConn,
	}
}

func (r *Repository) GetByIDs(ids []uint) ([]*domain.Prop, error) {
	ans, err := r.dbConn.Query(context.TODO(), "SELECT * FROM props WHERE id =  any($1);", ids)
	if err != nil {
		return nil, err
	}
	var out []*domain.Prop
	for ans.Next() {
		var prop domain.Prop
		err = ans.Scan(&prop.ID, &prop.InstalDate)
		if err != nil {
			return nil, err
		}
		out = append(out, &prop)
	}
	return out, nil
}

func (r *Repository) InsertProp(props *domain.Props) error {
	r.dbConn.Exec(context.TODO(), "INSERT INTO props VALUES($1, $2)", props.Props[0].ID, props.Props[0].InstalDate)
	return nil
}

func (r *Repository) UpdateProps(props *domain.Props) error {
	var ids []uint
	var dates []time.Time

	for _, v := range props.Props {
		ids = append(ids, v.ID)
		dates = append(dates, v.InstalDate)
	}
	_, err := r.dbConn.Exec(context.TODO(),
						`update props as p
						set    id=new.id, instalation_date=new.instalation_date
						from (select unnest($1::int[]) as id,unnest($2::date[]) as instalation_date) as new
						where p.id=new.id;`, 
						ids, dates)
	return err
}

func (r *Repository) InsertProps(props *domain.Props) error {
	const tableName = "props"
	count, err := r.dbConn.CopyFrom(
		context.Background(),
		pgx.Identifier{tableName},
		[]string{"id", "instalation_date"},
		pgx.CopyFromSlice(len(props.Props), func(i int) ([]interface{}, error) {
			return []interface{}{props.Props[i].ID, props.Props[i].InstalDate}, nil
		}))
	if err != nil {
		return err
	}
	if int(count) != len(props.Props) {
		return errors.New("not all props were inserted to DB")
	}
	return nil
}
