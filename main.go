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

	//初始化数据库、服务器等配置
	err := conf.Init()
	if err != nil {
		return
	}

	// 初始化服务
	service.Initialize(conf.Conf)

}

func main() {
	gin.SetMode(conf.Conf.Sever.RunMode)

	//路由配置
	router := routers.NewRouter()

	s := &http.Server{
		Addr:           conf.Conf.Sever.HttpIp + ":" + conf.Conf.Sever.HttpPort,
		Handler:        router,
		ReadTimeout:    conf.Conf.Sever.ReadTimeout,
		WriteTimeout:   conf.Conf.Sever.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Fprintf(os.Stdout, "emo-blog service listen on %s\n",
		fmt.Sprintf("http://%s:%s", conf.Conf.Sever.HttpIp, conf.Conf.Sever.HttpPort),
	)

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("run app failed: %s", err)
	}
}
