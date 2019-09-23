package main

import (
	"github.com/515074431/gin-antd/routers"
)
//主程序
func main() {
	r := routers.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	//r.Run(fmt.Sprintf(":%d", setting.HTTPPort))
	_ = r.Run()

	/*s := &http.Server{
		Addr:fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:r,
		ReadTimeout:setting.ReadTimeout,
		WriteTimeout:setting.WriteTimeout,
		MaxHeaderBytes: 1<<20,
	}

	s.ListenAndServe()*/
}
