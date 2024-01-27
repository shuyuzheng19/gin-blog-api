package service

import (
	"common-web-framework/common"
	"common-web-framework/config"
	"common-web-framework/helper"
	"common-web-framework/models"
	"common-web-framework/repository"
	"common-web-framework/request"
	"common-web-framework/response"
	"common-web-framework/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type FileServiceImpl struct {
	repository repository.FileRepository
}

func (fs FileServiceImpl) ClearFileContent(path string) error {
	return os.Truncate(path, 0)
}

func (fs FileServiceImpl) DeleteSystemFile(paths []string) int64 {
	var count int64
	for _, path := range paths {
		if err := os.Remove(path); err == nil {
			count++
		}
	}
	return count
}

func (fs FileServiceImpl) GetSystemFile(req request.SystemFileRequest) []response.SystemFileResponse {

	if req.Path == "" {
		req.Path = config.CONFIG.Upload.Path
	}

	if _, err := os.Stat(req.Path); os.IsNotExist(err) {
		helper.ErrorToResponse(common.NotFoundDir)
	}

	var result = make([]response.SystemFileResponse, 0)

	var files, err = ioutil.ReadDir(req.Path)

	if err != nil {
		helper.ErrorToResponse(common.OpenFileFail)
	}

	for _, info := range files {
		if !info.IsDir() {

			name := info.Name()

			if req.Keyword != "" && !strings.Contains(name, strings.ToLower(req.Keyword)) {
				continue
			}

			size := info.Size()

			createTime := utils.FormatDate(info.ModTime())

			updateTime := utils.FormatDate(info.ModTime())

			path := filepath.Join(req.Path, name)

			result = append(result, response.SystemFileResponse{
				Name:       name,
				Path:       path,
				Ext:        filepath.Ext(name),
				Size:       size,
				CreateTime: createTime,
				UpdateTime: updateTime,
			})
		}
	}

	return result

}

func (fs FileServiceImpl) UpdatePublic(uid int, id int, isPub bool) {

	var count = fs.repository.UpdatePublic(uid, id, isPub)

	if count == 0 {
		helper.ErrorToResponse(common.UpdateFail)
	} else {
		config.LOGGER.Info("修改文件是否公开", zap.Int("user_id", uid),
			zap.Int("id", id), zap.Bool("is_pub", isPub))
	}

}

func (fs FileServiceImpl) DeleteFile(uid int, ids []int) int64 {
	var count = fs.repository.DeleteFile(uid, ids)

	if count == 0 {
		helper.ErrorToResponse(common.UpdateFail)
	} else {
		config.LOGGER.Info("删除文件", zap.Int("user_id", uid),
			zap.Ints("ids", ids))
	}

	return count
}

func (fs FileServiceImpl) GetAdminFileList(uid int, req request.OtherAdminFilter) response.PageInfo {
	var files, count = fs.repository.GetAdminFile(uid, req)

	return response.PageInfo{
		Page:  req.Page,
		Count: count,
		Size:  common.AdminFileCount,
		Data:  files,
	}
}

func (fs FileServiceImpl) GetPublicFile(req request.FileRequest) response.PageInfo {
	var files, count = fs.repository.FindFileInfos(-1, req)

	return response.PageInfo{
		Page:  req.Page,
		Count: count,
		Size:  common.FilePageCount,
		Data:  files,
	}
}

func (fs FileServiceImpl) GetUserFile(uid int, req request.FileRequest) response.PageInfo {
	var files, count = fs.repository.FindFileInfos(uid, req)

	return response.PageInfo{
		Page:  req.Page,
		Count: count,
		Size:  common.FilePageCount,
		Data:  files,
	}
}

func GetFiles(ctx *gin.Context) []*multipart.FileHeader {
	var files, err = ctx.MultipartForm()

	if err != nil {
		config.LOGGER.Warn("请求体找不到文件")
		helper.ResultFailToToResponse(common.NoFile, ctx)
		return nil
	}

	return files.File["files"]
}

var mb = 1024 * 1024

func (fs FileServiceImpl) GlobalUploadFile(ctx *gin.Context, isImage bool, uid *int, isPub bool) (frs []response.SimpleFileResponse) {

	var u = config.CONFIG.Upload

	var files = GetFiles(ctx)

	var infos []models.FileInfo

	for _, file := range files {

		var size = file.Size

		var fileName = file.Filename

		var suffix = filepath.Ext(fileName)

		var create = time.Now()

		var errResponse response.SimpleFileResponse

		if isImage {
			if !utils.IsImageFile(suffix) {
				errResponse = response.SimpleFileResponse{
					Status:  "fail",
					Message: "这不是一个图片文件",
					Name:    fileName,
					Create:  utils.FormatDate(create),
				}
				frs = append(frs, errResponse)
				config.LOGGER.Info("上传文件错误 这不是一个图片文件", zap.String("fileName", fileName))
				continue
			} else if size > int64(u.MaxImageSize*mb) {
				errResponse = response.SimpleFileResponse{
					Status:  "fail",
					Message: "图片文件大小超出",
					Name:    fileName,
					Create:  utils.FormatDate(create),
				}
				frs = append(frs, errResponse)
				config.LOGGER.Info("上传文件错误 图片文件大小超出", zap.String("fileName", fileName))
				continue
			}
		} else {
			if size > int64(u.MaxFileSize*mb) {
				errResponse = response.SimpleFileResponse{
					Status:  "fail",
					Message: "文件大小超出",
					Name:    fileName,
					Create:  utils.FormatDate(create),
				}
				frs = append(frs, errResponse)
				config.LOGGER.Info("上传文件错误 文件大小超出", zap.String("fileName", fileName))
				continue
			}
		}

		var f, err = file.Open()

		if err != nil {
			frs = append(frs, response.SimpleFileResponse{
				Status:  "fail",
				Message: "文件打开失败",
				Name:    fileName,
				Create:  utils.FormatDate(create),
			})
			config.LOGGER.Info("上传文件错误 文件打开失败", zap.String("fileName", fileName))
			continue
		}

		var md5 = utils.GetFileMd5(f)

		var newName = md5 + suffix

		var saveFilePath = u.Path + "/" + newName

		var url string

		var existsMd5 = false

		url = u.Uri + "/" + newName

		if uid != nil {
			if dbUrl := fs.repository.FindByMd5(md5); dbUrl == "" {
				var uploadError = ctx.SaveUploadedFile(file, saveFilePath)
				if uploadError != nil {
					frs = append(frs, response.SimpleFileResponse{
						Status:  "fail",
						Message: "文件上传失败",
						Name:    fileName,
						Create:  utils.FormatDate(create),
					})
					config.LOGGER.Info("上传文件错误 文件上传失败", zap.String("fileName", fileName))
					continue
				}
			} else {
				existsMd5 = true
				url = dbUrl
			}
		}

		var fileInfo = models.FileInfo{
			OldName: fileName,
			NewName: newName,
			Suffix:  suffix,
			Size:    size,
			UserID:  uid,
			FileMd5: md5,
			IsPub:   isPub,
		}

		if !existsMd5 {
			fileInfo.FileMd5Info = models.FileMd5Info{
				Md5:          md5,
				Url:          url,
				AbsolutePath: saveFilePath,
			}
		}

		infos = append(infos, fileInfo)

		config.LOGGER.Info("文件上传成功", zap.String("fileName", fileInfo.NewName),
			zap.String("url", fileInfo.FileMd5Info.Url),
			zap.String("md5", fileInfo.FileMd5Info.Md5),
			zap.Int64("size", fileInfo.Size))

		frs = append(frs, response.SimpleFileResponse{
			Status:  "ok",
			Message: "上传成功",
			Name:    fileName,
			Create:  utils.FormatDate(create),
			Url:     url,
		})
	}

	if uid != nil && len(infos) > 0 {
		fs.repository.BatchSave(infos)
	}

	return frs
}

func (fs FileServiceImpl) UploadAvatarFile(ctx *gin.Context) []response.SimpleFileResponse {
	return fs.GlobalUploadFile(ctx, true, nil, false)
}

func (fs FileServiceImpl) UploadImageFile(ctx *gin.Context) []response.SimpleFileResponse {

	var user = utils.GetUserInfo(ctx)

	return fs.GlobalUploadFile(ctx, true, &user.ID, false)
}

func (fs FileServiceImpl) UploadFile(ctx *gin.Context) []response.SimpleFileResponse {
	var user = utils.GetUserInfo(ctx)

	var pubStr = ctx.DefaultQuery("is_pub", "false")

	var isPub, err = strconv.ParseBool(pubStr)

	if err != nil {
		isPub = false
	}

	return fs.GlobalUploadFile(ctx, false, &user.ID, isPub)
}

func (fs FileServiceImpl) AddFile(fileInfo models.FileInfo) {
	if err := fs.repository.Save(fileInfo); err != nil {
		helper.ErrorToResponse(common.AddFileFail)
	}
}

func NewFileService() FileService {
	var repository = repository.NewFileInfoRepository(config.DB)
	return FileServiceImpl{repository: repository}
}
