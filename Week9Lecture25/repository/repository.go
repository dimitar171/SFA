package repository

import (
	story "Lecture25/story"
	"database/sql"
	"log"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}
func (rp *Repository) GetLastStoryTimeStamp() time.Time {
	query := "SELECT s.timestamp FROM stories s ORDER BY s.timestamp DESC LIMIT 1"
	var tmstmp time.Time
	if err := rp.db.QueryRow(query).Scan(&tmstmp); err != nil {
		log.Println(err)
	}
	return tmstmp
}

func (rp *Repository) GetStories() []story.Story {
	query := "SELECT storyid, title, score FROM stories s"
	rows, err := rp.db.Query((query))
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()
	stories := []story.Story{}

	for rows.Next() {
		st := story.Story{}
		if err := rows.Scan(&st.Id, &st.Title, &st.Score); err != nil {
			log.Print(err)
		}
		stories = append(stories, st)
	}
	return stories
}

func (rp *Repository) SaveStories(sList []story.Story) {
	insertQuery := "INSERT INTO stories (storyid,title,score,timestamp) values(?,?,?,?)"
	for _, s := range sList {
		rp.db.Exec(insertQuery, s.Id, s.Title, s.Score, time.Now())
	}
}
