package helper

import (
	"mime/multipart"
	"path/filepath"
	"strings"
)

// ComposeParams 組合上傳檔案參數
//func ComposeParams(c *gin.Context, filePath string) (uploadOption *structs.UploadFileOption, apiError errorcode.Error) {
//	uploadOption = &structs.UploadFileOption{}
//
//	file, err := c.FormFile("file")
//	if err != nil {
//		go FatalLog(fmt.Sprintf("Get Upload File Error: %v", err))
//		apiError = errorcode.GetAPIError("GET_UPLOAD_FILE_ERROR")
//		return uploadOption, apiError
//	}
//
//	uploadOption.File = file
//	uploadOption.FileName = Md5EncryptionWithTime(file.Filename)
//	uploadOption.FileSize = file.Size
//	uploadOption.FileExt = getFileExt(file)
//	uploadOption.FilePath = filePath
//
//	return uploadOption, apiError
//}
//
//// CreateFile 建立新檔案
//func CreateFile(c *gin.Context, params *structs.UploadFileOption) (apiError errorcode.Error) {
//	// 組合路徑
//	path := params.FilePath + "/" + params.FileName + "." + params.FileExt
//
//	// 建立檔案
//	if err := c.SaveUploadedFile(params.File, path); err != nil {
//		FatalLog(fmt.Sprintf("Create File Error: %v", err))
//		apiError = errorcode.GetAPIError("CREATE_FILE_ERROR")
//	}
//
//	return apiError
//}

// getFileExt 取檔案副檔名(移除「.」符號)
func getFileExt(file *multipart.FileHeader) string {
	fileName := file.Filename
	fileExt := strings.Replace(filepath.Ext(fileName), ".", "", -1)
	return fileExt
}
