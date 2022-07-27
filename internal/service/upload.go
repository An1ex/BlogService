package service

import (
	"BlogService/global"
	"BlogService/pkg/upload"
	"errors"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.EncryptFileName(fileHeader.Filename)
	uploadSavePath := upload.GetSavePath()
	dst := uploadSavePath + "/" + fileName
	if !upload.CheckFileType(fileType) {
		return nil, errors.New("this type of file is not supported")
	}

	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not allowed or wrong")
	}

	if !upload.CheckMaxSize(fileType, fileHeader) {
		return nil, errors.New("exceeded maximum file limit")
	}

	if upload.CheckSavePath(uploadSavePath) {
		//保存文件夹不存在,即IsNotExist(err)=true
		err := upload.CreateSavePath(uploadSavePath, os.ModePerm)
		if err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}

	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions")
	}

	if err := upload.SaveFile(dst, fileHeader); err != nil {
		return nil, err
	}

	accessUrl := global.Config.App.UploadServerUrl + "/" + fileName
	return &FileInfo{
		Name:      fileName,
		AccessUrl: accessUrl,
	}, nil
}
