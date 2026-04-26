package main

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}

func langPrefix(lang string) string {
	if lang == "pl" {
		return ""
	}
	return "/" + lang
}

func MainPageHandler(c *gin.Context) {
	lang := getLang(c)
	c.HTML(http.StatusOK, "main_page.html", gin.H{
		"T":    langs[lang],
		"Lang": lang,
		"LP":   langPrefix(lang),
	})
}

func CatalogPageHandler(c *gin.Context) {
	lang := getLang(c)
	c.HTML(http.StatusOK, "catalog_page.html", gin.H{
		"T":    langs[lang],
		"Lang": lang,
		"LP":   langPrefix(lang),
	})
}

func AboutPageHandler(c *gin.Context) {
	lang := getLang(c)
	id := c.Param("id")
	c.HTML(http.StatusOK, "about_page.html", gin.H{
		"T":    langs[lang],
		"Lang": lang,
		"LP":   langPrefix(lang),
		"ID":   id,
	})
}

func ReviewsPageHandler(c *gin.Context) {
	lang := getLang(c)
	c.HTML(http.StatusOK, "review_page.html", gin.H{
		"T":    langs[lang],
		"Lang": lang,
		"LP":   langPrefix(lang),
	})
}

func CrudPageHandler(c *gin.Context) {
	lang := getLang(c)
	c.HTML(http.StatusOK, "add_page.html", gin.H{
		"T":    langs[lang],
		"Lang": lang,
		"LP":   langPrefix(lang),
	})
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
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	mainName := uuid.New().String() + filepath.Ext(mainPhoto.Filename)
	err = c.SaveUploadedFile(mainPhoto, "./static/uploads/main-photos/"+mainName)
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
	photosDir := []string{}
	for _, photo := range photos {
		name := uuid.New().String() + filepath.Ext(photo.Filename)
		err := c.SaveUploadedFile(photo, "./static/uploads/other-photos/"+name)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
			return
		}
		photosDir = append(photosDir, "./static/uploads/other-photos/"+name)
	}

	mainPhotoPath := "./static/uploads/main-photos/" + mainName

	vObj.MainPhoto = mainPhotoPath
	vObj.Photos = photosDir
	var v Vehicle
	v, err = UpdateVehicleById(vId, vObj)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, v)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
