package main

import (
	"log/slog"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}

func getLang(c *gin.Context) string {
	path := c.Request.URL.Path

	if strings.HasPrefix(path, "/ru/") || path == "/ru/" {
		return "ru"
	}
	if strings.HasPrefix(path, "/en/") || path == "/en/" {
		return "en"
	}
	return "pl"
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
		slog.Error("get all vehicles", "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal sever error"})
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
	vObj, err := GetVehicleById(vId)
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
	vObj, err := CreateVehicle(v)
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
	vObj, err := GetVehicleById(vId)
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
	v, err = UpdateVehicleById(vId, vObj)
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
	vObj, err := UpdateVehicleById(vId, v)
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
	err = DeleteVehicleById(vId)
	if err != nil {
		slog.Error("delete vehicle by id", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "status bad request"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "vehicle deleted"})
}
