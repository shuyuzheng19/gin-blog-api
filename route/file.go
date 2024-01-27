package route

import (
	"common-web-framework/config"
	"common-web-framework/controller"
	"common-web-framework/middle"
	"common-web-framework/service"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"time"
)

// LoadFileController 加载标签相关api
func (r *RouterSetup) LoadFileController(name string) *RouterSetup {
	var group = r.apiGroup.Group(name)

	var fileService = service.NewFileService()

	if cronJob != nil {
		cronJob.AddFunc("0 0 * * *", func() {
			var logConfig = config.CONFIG.Logger

			var logPath = filepath.Join(logConfig.LoggerDir, logConfig.DefaultName)

			var fileName = time.Now().Add(-time.Minute).Format("2006-01-02") + ".log"

			var file, err = os.OpenFile(filepath.Join(logConfig.LoggerDir, fileName), os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)

			defer file.Close()

			if err != nil {
				config.LOGGER.Info("创建日志文件失败", zap.String("error", err.Error()))
				return
			}

			var buff, _ = os.ReadFile(logPath)

			file.WriteString(string(buff))

			fileService.ClearFileContent(logPath)

			config.LOGGER.Info("已重新创建日志", zap.String("file_name", fileName))
		})
	}

	var fileController = controller.NewFileController(fileService)

	{
		group.GET("public", fileController.GetPublicFiles)
		group.POST("upload/avatar", fileController.UploadAvatarFile)
		group.GET("current", middle.JwtMiddle(middle.AdminRole), fileController.GetCurrentUserFiles)
		group.GET("admin/list", middle.JwtMiddle(middle.AdminRole), fileController.GetAdminFile)
		group.PUT("admin/delete", middle.JwtMiddle(middle.AdminRole), fileController.DeleteFile)
		group.PUT("admin/public", middle.JwtMiddle(middle.AdminRole), fileController.UpdateFilePublic)
		group.POST("upload/image", middle.JwtMiddle(middle.AdminRole), fileController.UploadImageFile)
		group.POST("upload/other", middle.JwtMiddle(middle.AdminRole), fileController.UploadOtherFile)
		group.POST("admin/system_file", middle.JwtMiddle(middle.SuperRole), fileController.GetSystemFile)
		group.POST("admin/system_file/delete", middle.JwtMiddle(middle.SuperRole), fileController.DeleteSystemFile)
		group.GET("admin/system_file/clear_content", middle.JwtMiddle(middle.SuperRole), fileController.ClearSystemFileContent)
		group.GET("admin/system_file/logs", middle.JwtMiddle(middle.SuperRole), fileController.GetLogFileList)
		group.GET("admin/system_file/current_log", middle.JwtMiddle(middle.SuperRole), fileController.GetCurrentLog)
		group.GET("admin/system_file/download", fileController.DownloadSystemFile)
	}

	return r
}
