package storage

import (
	"fmt"
	"strings"
	"time"
)

type statsModel struct {
	ID        int
	CreatedAt string
	Int1      int
	Int2      int
	Limit     int
	Str1      string
	Str2      string
	Result    string
}

func (model statsModel) ToDto() StatsDto {
	t, _ := time.Parse("2006-01-02T15:04", model.CreatedAt)
	return StatsDto{
		ID:        model.ID,
		CreatedAt: t,
		Int1:      model.Int1,
		Int2:      model.Int2,
		Limit:     model.Limit,
		Str1:      model.Str1,
		Str2:      model.Str2,
		Result:    model.Result,
	}
}

type StatsDto struct {
	ID        int
	CreatedAt time.Time
	Int1      int
	Int2      int
	Limit     int
	Str1      string
	Str2      string
	Result    string
}

func (dto StatsDto) ToModel() statsModel {
	return statsModel{
		ID:        dto.ID,
		CreatedAt: dto.CreatedAt.Format("2006-01-02T15:04"),
		Int1:      dto.Int1,
		Int2:      dto.Int2,
		Limit:     dto.Limit,
		Str1:      dto.Str1,
		Str2:      dto.Str2,
		Result:    dto.Result,
	}
}

type StatsCounterDto struct {
	Stats []StatsDto
	Count int
}

func GetLastCalls() (*StatsCounterDto, error) {
	// create SQLite file connection
	c, err := NewStatsClient()
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Insert stats into db
	stats, err := c.GetAll()
	if err != nil {
		return nil, err
	}

	return &StatsCounterDto{
		Stats: stats,
		Count: len(stats),
	}, nil
}

func SaveLastCall(str1 string, str2 string, int1 int, int2 int, limit int, results []string) error {
	// stringify to json
	resultsStr := fmt.Sprintf("[%s]", strings.Join(results, ","))

	// create SQLite file connection
	c, err := NewStatsClient()
	if err != nil {
		return err
	}
	defer c.Close()

	// Create stats dto
	dto := StatsDto{
		CreatedAt: time.Now(),
		Int1:      int1,
		Int2:      int2,
		Limit:     limit,
		Str1:      str1,
		Str2:      str2,
		Result:    resultsStr,
	}

	// Insert stats into db
	_, err = c.Insert(dto)
	if err != nil {
		return err
	}

	return nil
}
