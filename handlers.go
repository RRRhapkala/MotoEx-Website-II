package main

import (
	"log/slog"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}

func GetAllVehiclesHandler(c *gin.Context) {
	retVal, err := GetAllVehicles(c.Request.Context())
	if err != nil {
		slog.Error("get all vehicles", "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, retVal)
}

func GetVehicleByIdHandler(c *gin.Context) {
	vId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		slog.Error("get vehicle by id", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "status bad request"})
		return
	}
	vObj, err := GetVehicleById(c.Request.Context(), vId)
	if err != nil {
		slog.Error("get vehicle by id", "err", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "status not found"})
		return
	}
	c.JSON(http.StatusOK, vObj)
}

func CreateVehicleHandler(c *gin.Context) {
	var v Vehicle
	err := c.ShouldBindJSON(&v)
	if err != nil {
		slog.Error("create vehicle", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "status bad request"})
		return
	}
	vObj, err := CreateVehicle(c.Request.Context(), v)
	if err != nil {
		slog.Error("create vehicle", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "status bad request"})
		return
	}
	c.JSON(http.StatusCreated, vObj)
}

func UploadPhotosByIdHandler(c *gin.Context) {
	vId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		slog.Error("upload photos by vehicle id", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "status bad request"})
		return
	}
	vObj, err := GetVehicleById(c.Request.Context(), vId)
	if err != nil {
		slog.Error("upload photos by vehicle id", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "status bad request"})
		return
	}
	mainPhoto, err := c.FormFile("main_photo")
	if err != nil {
		slog.Error("upload photos by vehicle id", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "status bad request"})
		return
	}
	mainName := uuid.New().String() + filepath.Ext(mainPhoto.Filename)
	err = c.SaveUploadedFile(mainPhoto, "./static/uploads/main-photos/"+mainName)
	if err != nil {
		slog.Error("upload photos by vehicle id", "err", err)
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "status not acceptable"})
		return
	}
	form, err := c.MultipartForm()
	if err != nil {
		slog.Error("upload photos by vehicle id", "err", err)
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "status not acceptable"})
		return
	}
	photos := form.File["photos"]
	photosDir := []string{}
	for _, photo := range photos {
		name := uuid.New().String() + filepath.Ext(photo.Filename)
		err := c.SaveUploadedFile(photo, "./static/uploads/other-photos/"+name)
		if err != nil {
			slog.Error("upload photos by vehicle id", "err", err)
			c.JSON(http.StatusNotAcceptable, gin.H{"error": "status not acceptable"})
			return
		}
		photosDir = append(photosDir, "./static/uploads/other-photos/"+name)
	}

	mainPhotoPath := "./static/uploads/main-photos/" + mainName

	vObj.MainPhoto = mainPhotoPath
	vObj.Photos = photosDir
	var v Vehicle
	v, err = UpdateVehicleById(c.Request.Context(), vId, vObj)
	if err != nil {
		slog.Error("upload photos by vehicle id", "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "status internal server error"})
		return
	}
	c.JSON(http.StatusAccepted, v)
}

func UpdateVehicleByIdHandler(c *gin.Context) {
	var v Vehicle
	vId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		slog.Error("update vehicle by id", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "status bad request"})
		return
	}
	err = c.ShouldBindJSON(&v)
	if err != nil {
		slog.Error("update vehicle by id", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "status bad request"})
		return
	}
	vObj, err := UpdateVehicleById(c.Request.Context(), vId, v)
	if err != nil {
		slog.Error("update vehicle by id", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "status bad request"})
		return
	}
	c.JSON(http.StatusOK, vObj)
}

func DeleteVehicleByIdHandler(c *gin.Context) {
	vId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		slog.Error("delete vehicle by id", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "status bad request"})
		return
	}
	err = DeleteVehicleById(c.Request.Context(), vId)
	if err != nil {
		slog.Error("delete vehicle by id", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "status bad request"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "vehicle deleted"})
}
