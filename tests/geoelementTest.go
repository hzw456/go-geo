// package tests

// import "testing"

// func TestCreatePoint(t *testing.T) {
// 	if files := fileOperator.GetFilesInLesson(0); len(files) != 0 {
// 		t.Error("获取文件函数测试没通过")
// 	} else {
// 		t.Log("测试通过")
// 	}
// 	files := fileOperator.GetFilesInLesson(3400)
// 	if len(files) != 3 {
// 		t.Error("获取文件函数测试没通过")
// 	} else {
// 		if files[0].Path == "asda" && files[1].Path == "asd" && files[2].Path == "as" {
// 			t.Log("测试通过")
// 		} else {
// 			t.Error("校验文件顺序错误")
// 		}

// 	}
// }
