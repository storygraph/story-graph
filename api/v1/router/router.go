package router

import (
	"github.com/gin-gonic/gin"
	"github.com/storygraph/story-graph/api/v1/handler/stories"
)

type Story struct {
	Id           int64  `json:"id"`
	Title        string `json:"title"`
	LastEdited   string `json:"lastEdited"`
	CurrentEvent string `json:"currentEvent"`
}

var stories_list = []Story{
	Story{
		Id:           1,
		Title:        "Baron Hedgehog",
		LastEdited:   "05-05-2020",
		CurrentEvent: "The battle of Noxville",
	},
	Story{
		Id:           2,
		Title:        "Giggles",
		LastEdited:   "02-05-2020",
		CurrentEvent: "The poisonous apple",
	},
	Story{
		Id:           3,
		Title:        "Wreckage",
		LastEdited:   "02-04-2020",
		CurrentEvent: "The good sailor",
	},
	Story{
		Id:           4,
		Title:        "Friends in arms",
		LastEdited:   "01-04-2020",
		CurrentEvent: "Dead Freddy",
	},
	Story{
		Id:           5,
		Title:        "Neverspring",
		LastEdited:   "02-03-2020",
		CurrentEvent: "A fool's paradise",
	},
}

func getStoryIdx(id int64) int {
	for i, s := range stories_list {
		if s.Id == id {
			return i
		}
	}
	return -1
}
func DeleteStory(id int64) {
	for idx, s := range stories_list {
		if s.Id == id {
			stories_list = append(stories_list[0:idx], stories_list[idx+1:]...)
			return
		}
	}
	return
}

func New() (Router *gin.Engine) {
	Router = gin.Default()
	Router.Handle("GET", "/status", func(ctx *gin.Context) {
		ctx.String(200, "Running")
	})

	Router.Handle("GET", "/api/v1/stories", stories.ListStories)
	Router.Handle("GET", "/api/v1/stories/:id", stories.GetStory)
	Router.Handle("POST", "/api/v1/stories", stories.PostStory)
	Router.Handle("DELETE", "/api/v1/stories/:id", stories.DeleteStory)
	Router.Handle("PATCH", "/api/v1/stories/:id", stories.PatchStory)

	return
}
