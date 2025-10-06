package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pretorian41/goaggregate/models"
	"github.com/pretorian41/goaggregate/services"
	"sync"
)

func GetLoadAggregate(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	sources := []models.Job{
		{
			URL: "127.0.0.1:8080/api/agg/1",
			Priorities: map[string]int{
				"email": 0,
				"name":  2,
			},
		},
		{
			URL: "127.0.0.1:8080/api/agg/2",
			Priorities: map[string]int{
				"name": 0,
			},
		},
		{
			URL: "127.0.0.1:8080/api/agg/3",
			Priorities: map[string]int{
				"avatar_url": 0,
				"name":       1,
			},
		},
	}

	const numWorkers = 4

	jobs := make(chan models.Job, len(sources))
	results := make(chan models.ApiResult, len(sources))

	var wg sync.WaitGroup

	// === Запуск воркерів ===
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for job := range jobs {
				result := services.FetchFromAPI(job.URL, id)
				result.Priorities = job.Priorities
				results <- result
			}
		}()
	}

	// === Надсилаємо задачі ===
	for _, src := range sources {
		jobs <- src
	}
	close(jobs)

	// === Очікуємо завершення ===
	wg.Wait()
	close(results)

	// === Агрегуємо результати ===
	final := services.Reduce(results)
	final["id"] = id

	return ctx.JSON(final)
}

func GetSourceFirst(ctx *fiber.Ctx) error {
	jsonStr := `{"email": "test@test.com", "name":" Bar Dor"}`
	time.Sleep(1 * time.Second)

	ctx.Type("json")
	return ctx.SendString(jsonStr)
}

func GetSourceSecond(ctx *fiber.Ctx) error {
	jsonStr := `{"name": "John Foo"}`
	time.Sleep(3 * time.Second)

	ctx.Type("json")
	return ctx.SendString(jsonStr)
}

func GetSourceThird(ctx *fiber.Ctx) error {
	jsonStr := `{"avatar_url": "https://i.pravatar.cc/300", "name": "John Bar"}`
	time.Sleep(4 * time.Second)

	ctx.Type("json")
	return ctx.SendString(jsonStr)
}
