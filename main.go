package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github/durhunb/emo-blog/internal/conf"
	"github/durhunb/emo-blog/internal/routers"
	"github/durhunb/emo-blog/internal/service"
	"log"
	"net/http"
	"os"
	"strings"
)

type suites []string

//var (
//	noDefaultFeatures bool
//	features          suites
//)

func (s *suites) String() string {
	return strings.Join(*s, ",")
}

func (s *suites) Set(value string) error {
	for _, item := range strings.Split(value, ",") {
		*s = append(*s, strings.TrimSpace(item))
	}
	return nil
}

//func flagParse() {
//	flag.BoolVar(&noDefaultFeatures, "no-default-features", false, "whether use default features")
//	flag.Var(&features, "features", "use special features")
//	flag.Parse()
//}

func init() {
	//flagParse()

	// 初始化服务
	service.Initialize()
}

func main() {
	gin.SetMode(conf.ServerSetting.RunMode)

	//路由配置
	router := routers.NewRouter()

	s := &http.Server{
		Addr:           conf.ServerSetting.HttpIp + ":" + conf.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    conf.ServerSetting.ReadTimeout,
		WriteTimeout:   conf.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Fprintf(os.Stdout, "LiuLi service listen on %s\n",
		fmt.Sprintf("http://%s:%s", conf.ServerSetting.HttpIp, conf.ServerSetting.HttpPort),
	)

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("run app failed: %s", err)
	}
}
