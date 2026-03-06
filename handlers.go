package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}

func GetAllVehiclesHandler(c *gin.Context) {
	retVal, err := GetAllVehicles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, retVal)
}

func GetVehicleByIdHandler(c *gin.Context) {
	vId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	vObj, err := GetVehicleById(vId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "can't find car"})
		return
	}
	c.JSON(http.StatusOK, vObj)
}

func CreateVehicleHandler(c *gin.Context) {
	var v Vehicle
	err := c.ShouldBindJSON(&v)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	vObj, err := CreateVehicle(v)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't create vehicle"})
		return
	}
	c.JSON(http.StatusCreated, vObj)
}

func UploadPhotosByIdHandler(c *gin.Context) {
	vId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	vObj, err := GetVehicleById(vId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mainPhoto, err := c.FormFile("main_photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = c.SaveUploadedFile(mainPhoto, "./static/uploads/main-photos/"+mainPhoto.Filename)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	photos := form.File["photos"]
	for _, photo := range photos {
		err := c.SaveUploadedFile(photo, "./static/uploads/other-photos/"+photo.Filename)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
			return
		}
	}
	photosDir := []string{}
	mainPhotoPath := "./static/uploads/main-photos/" + mainPhoto.Filename
	for _, photo := range photos {
		path := "./static/uploads/other-photos/" + photo.Filename
		photosDir = append(photosDir, path)
	}

	vObj.MainPhoto = mainPhotoPath
	vObj.Photos = photosDir
	UpdateVehicleById(vId, vObj)
}

func UpdateVehicleByIdHandler(c *gin.Context) {
	var v Vehicle
	vId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = c.ShouldBindJSON(&v)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	vObj, err := UpdateVehicleById(vId, v)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't update vehicle"})
		return
	}
	c.JSON(http.StatusOK, vObj)
}

func DeleteVehicleByIdHandler(c *gin.Context) {
	vId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = DeleteVehicleById(vId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "vehicle deleted"})
}
