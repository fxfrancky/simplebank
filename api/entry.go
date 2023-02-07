package api

import (
	"database/sql"
	"net/http"

	db "github.com/fxfrancky/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createEntryRequest struct {
	AccountID int64 `json:"accountid" binding:"required"`
	Amount    int64 `json:"amount" binding:"required"`
}

func (server *Server) createEntry(ctx *gin.Context) {

	var req createEntryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateEntryParams{
		AccountID: req.AccountID,
		Amount:    req.Amount,
	}

	entry, err := server.store.CreateEntry(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, entry)

}

type getEntryRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getEntry(ctx *gin.Context) {

	var req getEntryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	entry, err := server.store.GetEntry(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, entry)
}

type listEntryRequest struct {
	AccountID int64 `form:"accountid" binding:"required"`
	Limit     int32 `form:"limit" binding:"required"`
	Offset    int32 `form:"offset" binding:"required"`
}

func (server *Server) listEntry(ctx *gin.Context) {

	var req listEntryRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListEntriesParams{
		AccountID: req.AccountID,
		Limit:     req.Limit,
		Offset:    req.Offset,
	}

	listEntries, err := server.store.ListEntries(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, listEntries)

}
