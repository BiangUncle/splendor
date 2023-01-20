package model

// Init 初始化默认数据
func Init() (err error) {
	err = LoadDefaultDevelopmentCard()
	if err != nil {
		return
	}
	err = LoadDefaultNobleTiles()
	if err != nil {
		return
	}
	InitTokenStack()
	return nil
}
