package errcode

var (
	ErrorGetTagFail     = NewError(20010001, "获取标签失败")
	ErrorGetTagListFail = NewError(20010002, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(20010003, "创建标签失败")
	ErrorUpdateTagFail  = NewError(20010004, "更新标签失败")
	ErrorDeleteTagFail  = NewError(20010005, "删除标签失败")
	ErrorCountTagFail   = NewError(20010006, "统计标签失败")
	ErrorUploadFileFail = NewError(20020001, "上传文件失败")
)
