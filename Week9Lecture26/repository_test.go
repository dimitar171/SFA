package repository

import (
	story "Lecture25/story"
	"database/sql"
	"testing"
	"time"

	_ "modernc.org/sqlite"
)

const (
	createStoryTable = "CREATE TABLE IF NOT EXISTS stories(storyid INT PRIMARY KEY,title TEXT,score INT,timestamp DATETIME,UNIQUE(storyid) ON CONFLICT REPLACE);"
	insertStory      = "INSERT INTO stories(storyid,title,score,timestamp) values(?,?,?,?)"
)

func TestLastStoryTimeStamp(t *testing.T) {
	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal("Failed to open in memory DB")
	}
	_, err = mockDb.Exec(createStoryTable)
	if err != nil {
		t.Fatal("Failed to create table")
	}

	repo := NewRepository(mockDb)
	res := repo.GetLastStoryTimeStamp()
	if res == time.Now() {
		t.Fatal("Failed to create initial conditions")
	}
	wantedTime := time.Now().Add(time.Hour)

	_, err = mockDb.Exec(insertStory, 0, "UnitTest", 15, wantedTime)
	if err != nil {
		t.Fatal(err)
	}
	_, err = mockDb.Exec(insertStory, 1, "UnitTest", 15, time.Now().Add(-time.Hour))
	if err != nil {
		t.Fatal(err)
	}
	res = repo.GetLastStoryTimeStamp()
	if !res.Equal(wantedTime) {
		t.Fatal("Failed to get lates timestamp")
	}
}

func TestGetStories(t *testing.T) {
	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal("Failed to open in memory DB")
	}
	_, err = mockDb.Exec(createStoryTable)
	if err != nil {
		t.Fatal("Failed to create table")
	}

	repo := NewRepository(mockDb)

	_, err = mockDb.Exec(insertStory, 0, "UnitTest", 15, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	res := repo.GetStories()
	wantedStorie := story.Story{Id: 0, Title: "UnitTest", Score: 15}
	if res[0] != wantedStorie {
		t.Fatal("Failed to get Story")
	}
}

func TestSaveStories(t *testing.T) {
	mockStorie := []story.Story{{Id: 0, Title: "UnitTest", Score: 15}}

	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal("Failed to open in memory DB")
	}
	_, err = mockDb.Exec(createStoryTable)
	if err != nil {
		t.Fatal("Failed to create table")
	}
	repo := NewRepository(mockDb)
	repo.SaveStories(mockStorie)
	res := repo.GetStories()
	if res[0] != mockStorie[0] {
		t.Fatal("Failed to save stories in DB")
	}
}
