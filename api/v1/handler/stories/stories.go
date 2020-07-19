package stories

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Story struct {
	Id           int    `json:"id"`
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

// Maybe private
func getStoryIdx(id int) int {
	for i, s := range stories_list {
		if s.Id == id {
			return i
		}
	}
	return -1
}

// Handlers

// "GET", "/api/v1/stories"
func ListStories(ctx *gin.Context) {
	ctx.JSON(200, stories_list)
}

// "GET", "/api/v1/stories/:id"
func GetStory(ctx *gin.Context) {
	var param, _ = strconv.ParseInt(ctx.Param("id"), 10, 64)
	idx := getStoryIdx(int(param))
	if idx == -1 {
		ctx.JSON(404, nil)
	} else {
		ctx.JSON(200, stories_list[idx])
	}
}

// "POST", "/api/v1/stories"
func PostStory(ctx *gin.Context) {
	var story Story
	ctx.BindJSON(&story)
	story.Id = stories_list[len(stories_list)-1].Id + 1
	stories_list = append(stories_list, story)
	ctx.JSON(200, story)
}

// Delete does not work
// "DELETE", "/api/v1/stories/:id"
func DeleteStory(ctx *gin.Context) {
	param, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	idx := getStoryIdx(int(param))
	if idx == -1 {
		ctx.String(404, "Nothing to remove")
	} else {
		for i, s := range stories_list {
			ctx.String(200, "Removing idx: %d", idx)
			ctx.String(200, "Len is:  %d", len(stories_list))
			if s.Id == idx {
				stories_list = append(stories_list[:i], stories_list[i+1:]...)
			}
		}
		ctx.String(200, "Removed story with id: %d", param)
	}
}

// "PATCH", "/api/v1/stories/:id"
func PatchStory(ctx *gin.Context) {
	param, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	idx := getStoryIdx(int(param))
	if idx == -1 {
		ctx.String(404, "Story with id: %d noy found", param)
	} else {
		var reqBody Story
		ctx.BindJSON(&reqBody)

		currentTime := time.Now()
		if reqBody.Title != "" {
			stories_list[idx].Title = reqBody.Title
			stories_list[idx].LastEdited = currentTime.Format("02-01-2006")
		}

		if reqBody.CurrentEvent != "" {
			stories_list[idx].CurrentEvent = reqBody.CurrentEvent
			stories_list[idx].LastEdited = currentTime.Format("02-01-2006")
		}

		ctx.String(200, "Pathced a story with id: %d", param)
	}
}
