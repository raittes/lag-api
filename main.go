package main

import (
	"flag"
	"io/ioutil"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

var (
	addr          = flag.String("addr", ":8888", "Listen port")
	static_yml    = flag.String("static-rules", "", "YML with static rules")
	lag           = flag.Duration("lag", 0*time.Millisecond, "response delay in ms")
	static_config StaticConfig
)

type StaticConfig map[string]Rule

type Rule struct {
	StatusCode int               `yaml:"status_code"`
	Headers    map[string]string `yaml:"headers"`
	Body       string            `yaml:"body"`
	Lag        string            `yaml:"lag"`
}

func read_yaml(file_path string) {
	data, err := ioutil.ReadFile(file_path)
	if err != nil {
		log.Fatal("Failed to read YML: ", err)
	}
	err = yaml.Unmarshal(data, &static_config)
	if err != nil {
		log.Fatal("Failed to parse YML: ", err)
	} else {
		log.Println("Loaded", file_path)
	}
}

func staticHandler() *gin.Engine {
	router := gin.Default()
	router.NoRoute(func(c *gin.Context) {

		route := c.Request.URL.Path
		var rule Rule
		if r, ok := static_config[route]; ok {
			rule = r
		} else {
			c.String(404, "Not found")
			return
		}

		// headers
		for header, value := range rule.Headers {
			c.Request.Header.Set(header, value)
		}

		// status code
		status_code := 200
		if code := rule.StatusCode; code != 0 {
			status_code = code
		}

		// body
		body := rule.Body

		// lag
		if l := rule.Lag; l != "" {
			if t, e := time.ParseDuration(l); e == nil {
				time.Sleep(t)
			}
		} else if *lag > 0 {
			time.Sleep(*lag)
		}

		c.String(status_code, body)
	})
	return router
}

func main() {
	flag.Parse()

	var handler *gin.Engine
	gin.SetMode(gin.ReleaseMode)

	if *static_yml != "" {
		read_yaml(*static_yml)
		handler = staticHandler()
	} else {
		log.Fatal("No handler specified")
	}

	if err := handler.Run(*addr); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	} else {
		log.Println("Listening at", *addr)
	}
}
