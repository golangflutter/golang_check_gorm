package controller

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type ImageUploadController struct{}

func NewImageUploadController() *ImageUploadController {
    return &ImageUploadController{}
}

// UploadImage handles the file upload
// func (controller *ImageUploadController) UploadImage(ctx *gin.Context) {
   

 
//     title := ctx.DefaultPostForm("productName", "default_title")
//     category := ctx.DefaultPostForm("category", "default_title")
//     fmt.Println("this is tittle from fe", title)
 
 
   
//     file, err := ctx.FormFile("file")
//     if err != nil {
//         ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
//         return
//     }
//     uploadDir := "uploads/" + category + "/" + title


//     uploadDir1 := "uploads/" + title
//     fmt.Println("this is tittle from test1341234123412", uploadDir1)
//     if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
//         os.Mkdir(uploadDir, os.ModePerm)
//     }

//     err = ctx.SaveUploadedFile(file, filepath.Join(uploadDir, file.Filename))
//     if err != nil {
//         ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
//         return
//     }

//     // Generate the URL for the uploaded file
//     uploadedURL := "/" + uploadDir + "/" + file.Filename

//     ctx.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "url": uploadedURL})
// }


// ServeImage serves the original image
func (controller *ImageUploadController) ServeImage(ctx *gin.Context) {
    imageFilename := ctx.Query("file")
    if imageFilename == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "File not specified"})
        return
    }

    imagePath := filepath.Join("uploads", imageFilename)
    ctx.File(imagePath)
}

// ServeThumbnail serves a thumbnail image
func (controller *ImageUploadController) ServeThumbnail(ctx *gin.Context) {
    imageFilename := ctx.Query("file")
    if imageFilename == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "File not specified"})
        return
    }
}

// DeleteImage deletes an uploaded image


func (controller *ImageUploadController) DeleteImage(ctx *gin.Context) {
    fmt.Println("This is delete image function")
    imageFilename := ctx.Query("file")
    if imageFilename == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "File not specified"})
        return
    }
    fmt.Println("This is delete image function imageFilename11111", imageFilename)
    // URL-decode the filename if needed
    imageFilename, err := url.PathUnescape(imageFilename)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file name"})
        return
    }
    fmt.Println("This is delete image function imageFilename", imageFilename)
    imagePath := filepath.Join( imageFilename)
    fmt.Println("This is delete image function imagePath333333333", imagePath)
    
    // Check if the file exists
    _, err = os.Stat(imagePath)
    if err != nil {
        if os.IsNotExist(err) {
            ctx.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
        } else {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking file"})
        }
        return
    }
    fmt.Println("This is delete image function file exi444444444555555t" )
    // Attempt to delete the file

    
    err = os.Remove(imagePath)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the image"})
        return
    }
    fmt.Println("This is delete image function file exit done" )
    ctx.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully"})
}

// -------------
func (controller *ImageUploadController) UploadImage(ctx *gin.Context) {
    fmt.Println("this is UploadImage new 1111111", )
    // Get the title and category from the request form
    title := ctx.DefaultPostForm("productName", "default_title")
    category := ctx.DefaultPostForm("category", "default_title")

    // Get the uploaded file
    file, err := ctx.FormFile("file")
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
        return
    }

    // Create the upload directory if it doesn't exist
    uploadDir := filepath.Join("uploads", category, title)
    if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
        if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating directory"})
            return
        }
    }

    // Save the uploaded file
    err = ctx.SaveUploadedFile(file, filepath.Join(uploadDir, file.Filename))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving file"})
        return
    }

    // Generate the URL for the uploaded file
    uploadedURL := "/" + uploadDir + "/" + file.Filename

    // Respond with the success message and the uploaded file URL
    ctx.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "url": uploadedURL})
}