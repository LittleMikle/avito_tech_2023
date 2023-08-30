package handler

import (
	tech "github.com/LittleMikle/avito_tech_2023"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Create User Segment
// @Tags create
// @Description create segment
// @ID create-user-segment
// @Accept  json
// @Produce  json
// @Param input body tech.UserSegment true "segment info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/users/create/:id [post]
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

// @Summary Delete User Segment
// @Tags delete
// @Description delete segment
// @ID delete-user-segment
// @Accept  json
// @Produce  json
// @Param input body tech.UserSegment true "segment info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/users/delete/:id [delete]
func (h *Handler) deleteUsersSeg(c *gin.Context) {
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
		err := h.services.UsersSeg.DeleteUsersSeg(userId, val)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "created",
	})
}

// @Summary Get User Segment
// @Tags get
// @Description get segments
// @ID get-segment
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/users/:id [get]
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
}

// @Summary Get User History Segment
// @Tags history
// @Description get user history
// @ID history-segment
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/users/history/:id [get]
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

	c.JSON(http.StatusOK, "HISTORY SAVED")
}

// @Summary Schedule Post and Delete User Segment
// @Tags post delete
// @Description schedule post delete segment
// @ID schedule-post-segment
// @Accept  json
// @Produce  json
// @Param input body tech.Segment true "segment info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/users/schedule/:id [post]
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

// @Summary Random User Segment
// @Tags random
// @Description random segment
// @ID random-segment
// @Accept  json
// @Produce  json
// @Param input body tech.Segment true "segment info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/users/random/:id [post]
func (h *Handler) randomCreate(c *gin.Context) {
	var input tech.Segment

	percent := c.Query("percent")
	if percent != "" && percent != "0" {
		percent, err := strconv.Atoi(percent)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "invalid percent param")
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
