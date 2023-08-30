package handler

import (
	tech "github.com/LittleMikle/avito_tech_2023"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary CreateSegment
// @Tags create
// @Description create segment
// @ID create-segment
// @Accept  json
// @Produce  json
// @Param input body tech.Segment true "segment info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/segments/create [post]
func (h *Handler) createSegment(c *gin.Context) {
	var input tech.Segment
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Segmentation.CreateSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary DeleteSegment
// @Tags delete
// @Description delete segment
// @ID delete-segment
// @Accept  json
// @Produce  json
// @Param input body tech.Segment true "segment info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/segments/:id [post]
func (h *Handler) deleteSegment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Segmentation.DeleteSegment(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
