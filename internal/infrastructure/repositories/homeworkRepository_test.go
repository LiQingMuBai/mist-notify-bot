package repositories

import (
	"github.com/spf13/viper"
	"homework_bot/internal/domain"
	"homework_bot/internal/infrastructure/configs"
	"os"
	"testing"
	"time"
)

func TestHomeworkRepository_Create(t *testing.T) {
	type args struct {
		homework      domain.Homework
		extendedID    int
		extendedError error
	}

	tests := []args{
		{
			homework: domain.Homework{
				Id:          0,
				Name:        "Test1",
				Description: "Test1 description",
				Images:      nil,
				Tags:        []string{"tag1", "tag2"},
				Deadline:    time.Date(2024, 9, 15, 0, 0, 0, 0, time.UTC),
			},
			extendedID:    1,
			extendedError: nil,
		},
		{
			homework: domain.Homework{
				Id:          0,
				Name:        "Test2",
				Description: "",
				Images:      []string{"/home/nikita/GolandProjects/homework_bot/media/7b984546-4f68-477a-9cd7-0ec428bf0440.jpg"},
				Tags:        []string{"tag1"},
				Deadline:    time.Date(2024, 9, 18, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	db, err := configs.NewPostgresDB(configs.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		t.Error(err)
	}

	repo := NewHomeworkRepository(db)

	for _, tt := range tests {
		id, err := repo.Create(tt.homework)
		if err != nil {
			t.Error(err)
		} else if id != tt.extendedID {
			t.Errorf("got %d, want %d", id, tt.extendedID)
		}
	}
}
