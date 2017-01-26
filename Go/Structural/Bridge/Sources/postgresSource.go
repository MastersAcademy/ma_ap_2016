package Sources

import (
	"github.com/MastersAcademy/ma_ap_2016/Go/Structural/Bridge/ReportEngines"
	"github.com/MastersAcademy/ma_ap_2016/Go/Structural/Bridge/Models"
	"database/sql"
	_ "github.com/lib/pq"
	"errors"
	"fmt"
	"log"
)

type PostgresSource struct {
	DSN string
	DB *sql.DB
	engine ReportEngines.ReportEngine
}

func NewPostgresSource(dsn string, engine ReportEngines.ReportEngine) (Source, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Panicf("Failed to connect to datastore: %s", err.Error())
		return nil, errors.New("Unnable connect to postresql database")
	}

	return &PostgresSource{
		DSN:            dsn,
		DB:             db,
		engine:         engine,
	}, nil

}

func (s *PostgresSource) GetVideos() (string, error) {
	videos, err := s.getVideos()
	if err != nil {
		return "", err
	}
	return s.engine.BuildReport(videos), nil
}

func (s *PostgresSource) getVideos() ([]Models.Video, error){
	var id, title, url string
	var videos []Models.Video

	q_s := fmt.Sprintf("SELECT original_id, title, url FROM video limit 10")
	rows, err := s.DB.Query(q_s)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Videos not found")
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &title, &url)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New("Videos not found")
			}
			return nil, err
		}
		videos = append(videos, Models.Video{
			ID: id,
			Title: title,
			URL: url,
		})
	}
	return videos, nil
}
