package example

import (
	"fmt"
	"log"

	"github.com/zhtfirst/go-packages/handler"
)

func HandlerTest() {
	// 初始化打印日志,记录项目中需要打印的日志,不需要打印日志可以注释掉
	_ = handler.NotesLog()

	log.Println(fmt.Sprintf("UID: %s", 123), 999)

}
