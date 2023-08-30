package handler

import (
	"fmt"
	tech "github.com/LittleMikle/avito_tech_2023"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createUsersSeg(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	var input tech.UserSegment
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if len(input.Segments) == 0 {
		newErrorResponse(c, http.StatusBadRequest, "segments input can't be empty")
		return
	}
	for _, val := range input.Segments {
		err := h.services.UsersSeg.CreateUsersSeg(userId, val)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "created",
	})
}

func (h *Handler) getUsersSeg(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	segments, err := h.services.UsersSeg.GetUserSeg(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, segments)
	fmt.Println(userId)
}

func (h *Handler) getHistory(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	err = h.services.UsersSeg.GetHistory(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) scheduleDelete(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}
	days := c.Query("days")
	if days != "" && days != "0" {
		days, err := strconv.Atoi(days)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "invalid days param")
			return
		}
		var input tech.Segment
		if err = c.BindJSON(&input); err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		err = h.services.UsersSeg.ScheduleDelete(userId, days, input)
	}

	c.JSON(http.StatusOK, "SCHEDULER SET")
}

func (h *Handler) randomCreate(c *gin.Context) {
	var input tech.Segment

	percent := c.Query("percent")
	if percent != "" && percent != "0" {
		percent, err := strconv.Atoi(percent)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "invalid days param")
			return
		}

		if err = c.BindJSON(&input); err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		percentFloat := float64(percent)
		err = h.services.UsersSeg.RandomSegments(input, percentFloat)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "created",
	})
}
