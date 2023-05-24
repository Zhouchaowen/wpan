package handler

//// Index 首页
//func Index(c *gin.Context) {
//	userId := c.GetString(middleware.UserIdKey)
//	//获取用户信息
//	//user := model.GetUserInfoById(fi.UserId)
//
//	//获取当前目录所有文件
//	//files := model.GetUserFile(fi.FolderId, user.FileStoreId)
//
//	//获取当前目录所有文件夹
//	//fileFolder := model.GetFileFolder(fi.FolderId, user.FileStoreId)
//
//	//获取父级的文件夹信息
//	parentFolder := model.GetParentFolder(fi.FolderId)
//
//	//获取当前目录所有父级
//	currentAllParent := model.GetCurrentAllParent(parentFolder, make([]model.FileFolder, 0))
//
//	//获取当前目录信息
//	currentFolder := model.GetCurrentFolder(fi.FolderId)
//
//	//获取用户文件使用明细数量
//	//fileDetailUse := model.GetFileDetailUse(user.FileStoreId)
//
//	utils.ToResponse(c, gin.H{
//		"currAll": "active",
//		//"user":             user,
//		"fId":   currentFolder.Id,
//		"fName": currentFolder.FileFolderName,
//		//"files":            files,
//		//"fileFolder":       fileFolder,
//		"parentFolder":     parentFolder,
//		"currentAllParent": currentAllParent,
//		//"fileDetailUse":    fileDetailUse,
//	})
//}
