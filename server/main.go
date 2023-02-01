package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/patrickmn/go-cache"
	"github.com/unknwon/com"
	"net/http"
	"os"
	"regexp"
	"time"
)

// Resp /*接口返回数据封装*/
type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

/*错误码*/
const (
	SUCCESS               = 200
	ERROR                 = 500
	ErrorNotExistRouting  = 10001
	ErrorParameterIllegal = 10002
)

var MsgFlags = map[int]string{
	SUCCESS:               "success",
	ERROR:                 "fail",
	ErrorNotExistRouting:  "路由不存在",
	ErrorParameterIllegal: "请求参数错误",
}

var aliClient = resty.New()
var aliDriveAccessTokenKey = "accessToken:aliDrive:ac:"
var aliDriveId = ""
var RefreshToken = ""
var dataFile = "data.json"

// 设置缓存超时时间和清理时间
var cacheClient = cache.New(5*time.Minute, 10*time.Minute)

type Json map[string]interface{}

type AliFile struct {
	DriveId       string     `json:"drive_id"`
	Encrypted     bool       `json:"encrypted"`
	CreatedAt     *time.Time `json:"created_at"`
	FileExtension string     `json:"file_extension"`
	FileId        string     `json:"file_id"`
	Type          string     `json:"type"`
	Name          string     `json:"name"`
	Category      string     `json:"category"`
	ParentFileId  string     `json:"parent_file_id"`
	UpdatedAt     string     `json:"updated_at"`
	Size          int64      `json:"size"`
	Thumbnail     string     `json:"thumbnail"`
	Url           string     `json:"url"`
}
type AliDrive struct{}
type AliFiles struct {
	Items      []AliFile `json:"items"`
	NextMarker string    `json:"next_marker"`
}
type AliRespError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
type TokenResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}
type IndexData struct {
	Name  string     `json:"name"`
	Icon  string     `json:"icon"`
	Lists []ListItem `json:"lists"`
}
type ListItem struct {
	Path  string `json:"path"`
	Logo  string `json:"logo"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

// InitRouter /* 自定义路由*/
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//#debug or release
	gin.SetMode(gin.ReleaseMode)
	//定义一个请求Cors的匿名函数
	cors := func() gin.HandlerFunc {
		return func(context *gin.Context) {
			method := context.Request.Method
			context.Header("Access-Control-Allow-Origin", "*")
			context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			context.Header("Access-Control-Allow-Credentials", "true")
			if method == "OPTIONS" {
				context.AbortWithStatus(http.StatusNoContent)
			}
			context.Next()
		}
	}
	r.Use(cors())
	r.Static("css", "templates/css")
	r.Static("js", "templates/js")
	r.Static("fonts", "templates/fonts")
	r.Static("img", "templates/img")
	r.LoadHTMLGlob("./templates/*.html")
	r.GET("/", IndexPage)
	r.POST("/index", GetIndexDataApi)
	r.POST("/lists", GetFolderDataApi)
	r.POST("/player", GetPlayerDataApi)
	return r
}

func IndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "",
	})
}

func GetIndexDataApi(c *gin.Context) {
	data := getIndexData()
	RespFunc(c, http.StatusOK, SUCCESS, data)
	return
}

// GetFolderDataApi /*路由实现方法*/
func GetFolderDataApi(c *gin.Context) {
	c.Handler()
	path := com.StrTo(c.PostForm("path")).String()

	//禁止手动输入root查询跟目录
	if path == "root" {
		RespFunc(c, http.StatusBadRequest, ErrorParameterIllegal, "")
		return
	}

	req, err := AliDrive{}.getFiles(path)
	fmt.Errorf("%+v:%s", req, err)
	if err != nil {
		RespFunc(c, http.StatusBadRequest, ErrorNotExistRouting, err.Error())
		return
	}
	RespFunc(c, http.StatusOK, SUCCESS, req)
	return
}

func GetPlayerDataApi(c *gin.Context) {
	path := com.StrTo(c.PostForm("path")).String()
	resp, err := AliDrive{}.GetVideoPreviewPlayInfo(path)
	if err != nil {
		RespFunc(c, http.StatusBadRequest, ERROR, err.Error())
		return
	} else {
		RespFunc(c, http.StatusOK, SUCCESS, resp)
		return
	}
}

// RespFunc /*Func返回数据封装*/
func RespFunc(c *gin.Context, httpCode int, code int, data interface{}) {
	c.JSON(httpCode, Resp{
		Code:    code,
		Message: GetMsg(code),
		Data: func(data interface{}) interface{} {
			if data == nil {
				return ""
			}
			return data
		}(data),
	})
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

// GetVideoPreviewPlayInfo /*阿里云盘相关操作 start*/
func (driver AliDrive) GetVideoPreviewPlayInfo(path string) (interface{}, error) {
	var err error
	var resp Json
	var e AliRespError
	_, err = aliClient.R().SetResult(&resp).SetError(&e).
		SetHeader("authorization", "Bearer\t"+driver.GetAliDriveAccessToken()).
		SetBody(Json{
			"drive_id": aliDriveId,
			"file_id":  path,
			"category": "live_transcoding",
		}).Post("https://api.aliyundrive.com/v2/file/get_video_preview_play_info")
	if err != nil {
		return nil, err
	}
	if e.Code != "" {
		if e.Code == "AccessTokenInvalid" {
			err = driver.RefreshToken()
			if err != nil {
				return nil, err
			} else {
				return driver.GetVideoPreviewPlayInfo(path)
			}
		}
		return nil, fmt.Errorf("%s", e.Message)
	}
	return resp, nil
}
func (driver AliDrive) getFiles(path string) ([]AliFile, error) {
	marker := "first"
	res := make([]AliFile, 0)
	for marker != "" {
		if marker == "first" {
			marker = ""
		}
		var resp AliFiles
		var e AliRespError
		_, err := aliClient.R().
			SetResult(&resp).
			SetError(&e).
			SetHeader("authorization", "Bearer\t"+driver.GetAliDriveAccessToken()).
			SetBody(Json{
				"drive_id":                aliDriveId,
				"fields":                  "*",
				"image_thumbnail_process": "image/resize,w_400/format,jpeg",
				"image_url_process":       "image/resize,w_1920/format,jpeg",
				"limit":                   200,
				"marker":                  marker,
				"order_by":                "name",
				"order_direction":         "ASC",
				"parent_file_id":          path,
				"video_thumbnail_process": "video/snapshot,t_0,f_jpg,ar_auto,w_300",
				"url_expire_sec":          14400,
			}).Post("https://api.aliyundrive.com/v2/file/list")
		if err != nil {
			return nil, err
		}
		if e.Code != "" {
			if e.Code == "AccessTokenInvalid" {
				err = driver.RefreshToken()
				if err != nil {
					return nil, err
				} else {
					return driver.getFiles(path)
				}
			}
			return nil, fmt.Errorf("%s", e.Message)
		}
		marker = resp.NextMarker
		res = append(res, resp.Items...)
	}
	return res, nil
}
func (driver AliDrive) RefreshToken() error {

	accessToken := driver.GetAliDriveAccessToken()
	if accessToken != "" {
		return nil
	}
	url := "https://auth.aliyundrive.com/v2/account/token"
	var resp TokenResp
	var e AliRespError

	_, err := aliClient.R().
		SetBody(Json{"refresh_token": RefreshToken, "grant_type": "refresh_token"}).
		SetResult(&resp).
		SetError(&e).
		Post(url)
	if err != nil {
		return err
	}

	if e.Code != "" {
		return fmt.Errorf("failed to refresh token: %s", e.Message)
	} else {
	}
	driver.SetAliDriveAccessToken(resp.AccessToken, time.Second*72)
	return nil
}
func (driver AliDrive) GetAliDriveAccessToken() string {
	value, found := cacheClient.Get(aliDriveAccessTokenKey) // 42,true
	if found {
		return value.(string)
	}
	return ""
}
func (driver AliDrive) SetAliDriveAccessToken(value string, expiration time.Duration) {
	// 设置缓存值并带上过期时间
	cacheClient.Set(aliDriveAccessTokenKey, value, expiration)
}

/*阿里云盘相关操作 end*/

/*获取首页数据*/
func getIndexData() *[]IndexData {
	var retData []IndexData
	file, err := os.Open(dataFile)
	defer file.Close()
	if err != nil {
		panic(fmt.Sprintf("Failed to open config file '%s': %s\n", dataFile, err))
	}
	fi, _ := file.Stat()
	if size := fi.Size(); size > (10 << 20) {
		panic(fmt.Sprintf("config file (%q) size exceeds reasonable limit (%d) - aborting", dataFile, size))
	}
	if fi.Size() == 0 {
		panic(fmt.Sprintf("config file (%q) is empty, skipping", dataFile))
	}
	buffer := make([]byte, fi.Size())
	_, err = file.Read(buffer)

	var StripComments = func(data []byte) ([]byte, error) {
		data = bytes.Replace(data, []byte("\r"), []byte(""), 0)
		lines := bytes.Split(data, []byte("\n")) //split to muli lines
		filtered := make([][]byte, 0)
		for _, line := range lines {
			match, err := regexp.Match(`^\s*#`, line)
			if err != nil {
				return nil, err
			}
			if !match {
				filtered = append(filtered, line)
			}
		}
		return bytes.Join(filtered, []byte("\n")), nil
	}
	buffer, err = StripComments(buffer) //去掉注释
	if err != nil {
		panic(fmt.Sprintf("Failed to strip comments from json: %s\n", err))
	}
	buffer = []byte(os.ExpandEnv(string(buffer))) //特殊
	err = json.Unmarshal(buffer, &retData)        //解析json格式数据
	if err != nil {
		panic(fmt.Sprintf("Failed unmarshalling json: %s\n", err))
	}
	return &retData
}

func InitAliClient() {
	aliClient.
		SetTimeout(time.Second*20).
		SetRetryCount(3).
		SetHeader("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36").
		SetHeader("content-type", "application/json").
		SetHeader("origin", "https://www.aliyundrive.com")
}
func init() {
	InitAliClient()
}
func main() {
	r := InitRouter()

	err := r.Run("0.0.0.0:50288")
	if err != nil {
		fmt.Errorf("failed to start: %s", err.Error())
	}
}
