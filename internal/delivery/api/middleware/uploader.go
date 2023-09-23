package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	baseModel "github.com/rochmanramadhani/go-lazisnu-api/internal/model/abstraction"
	"github.com/rochmanramadhani/go-lazisnu-api/pkg/util/ctxval"
	res "github.com/rochmanramadhani/go-lazisnu-api/pkg/util/response"
	"io"
	"os"
	"path/filepath"
	"time"
)

// UploadFiles is a middleware for uploading files
// It will save the files to the public directory
// The feature parameter is used to determine the subdirectory
// If the feature is empty, the files will be saved to the root of the public directory
// Example:
//
//	// Upload files to the public directory with subdirectory file extension
//	e.POST("/upload", middleware.Upload("", handler.Upload))
//
//	// Upload files to the public directory with subdirectory feature
//	e.POST("/upload", middleware.Upload("user", handler.Upload))
func UploadFiles(feature string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Multipart form
			form, err := c.MultipartForm()
			if err != nil {
				return res.ErrorBuilder(res.Constant.Error.InternalServerError, err).Send(c)
			}
			files := form.File["files"]

			// Create a slice to store file information
			var uploadedFiles []baseModel.UploadFileContext

			for _, file := range files {
				// Source
				src, err := file.Open()
				if err != nil {
					return res.ErrorBuilder(res.Constant.Error.InternalServerError, err).Send(c)
				}
				defer src.Close()

				// Determine the destination directory based on the file extension and feature
				extension := filepath.Ext(file.Filename)
				var destinationDir string

				var fileName string

				switch extension {
				case ".pdf":
					fileName = fmt.Sprintf("file-%s%s", time.Now().Format("20060102-150405"), extension)
					destinationDir = os.Getenv("FILE_DIR")
				case ".png", ".jpeg", ".jpg":
					fileName = fmt.Sprintf("img-%s%s", time.Now().Format("20060102-150405"), extension)
					destinationDir = os.Getenv("IMAGE_DIR")
				default:
					return res.ErrorBuilder(res.Constant.Error.BadRequest, fmt.Errorf("invalid file extension: %s", extension)).Send(c)
				}

				// Add feature subdirectory if specified
				if feature != "" {
					destinationDir = filepath.Join(destinationDir, feature)
				}

				// Create the destination directory if it doesn't exist
				if err := os.MkdirAll(destinationDir, os.ModePerm); err != nil {
					return res.ErrorBuilder(res.Constant.Error.InternalServerError, err).Send(c)
				}

				// Create the destination file
				dst, err := os.Create(filepath.Join(destinationDir, fileName))
				if err != nil {
					return res.ErrorBuilder(res.Constant.Error.InternalServerError, err).Send(c)
				}
				defer dst.Close()

				// Append file information to the slice
				uploadedFiles = append(uploadedFiles, baseModel.UploadFileContext{
					FilePath: filepath.Join(destinationDir, fileName),
					FileName: fileName,
				})

				// Copy the file
				if _, err = io.Copy(dst, src); err != nil {
					return res.ErrorBuilder(res.Constant.Error.InternalServerError, err).Send(c)
				}
			}

			ctx := ctxval.SetUploadFileValues(c.Request().Context(), &uploadedFiles)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}
