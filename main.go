package main

import (
	"fmt"
	. "github.com/qiniu/api/conf"
	"github.com/qiniu/api/fop"
	"github.com/qiniu/api/io"
	"github.com/qiniu/api/rs"
	"log"
)

func init() {

	ACCESS_KEY = "gAmUFxyUfpu50dJpeI5_LqQuEyV6V4NxbEXBbkqo"
	SECRET_KEY = "lMC6kUfeVN3uXXHd1mGWB37-B99BEKHFWKOZ4Acb"
}

//GET upload access token
func uptoken(bucketName string) string {
	putPolicy := rs.PutPolicy{
		Scope: bucketName,
		//CallbackUrl: callbackUrl,
		//CallbackBody:callbackBody,
		//ReturnUrl:   returnUrl,
		//ReturnBody:  returnBody,
		//AsyncOps:    asyncOps,
		//EndUser:     endUser,
		//Expires:     expires,
	}
	return putPolicy.Token(nil)
}

func main() {
	//上传本地文件
	upload()

	//5.1 获取文件信息
	//getFileInfo()

	//6.1.1 查看图像属性
	//imageAttr()

	//5.2 删除文件
	//delFile()

}

//6.1.1 查看图像属性
func imageAttr() {
	var imageUrl = "http://attach.qiniudn.com/pic.jpg"
	ii := fop.ImageInfo{}
	infoRet, err := ii.Call(nil, imageUrl)
	if err != nil {
		// 产生错误
		log.Println("fop getImageInfo failed:", err)
		return
	}

	log.Println(infoRet.Height, infoRet.Width, infoRet.ColorModel, infoRet.Format)
}

func makeImageInfoUrl(imageUrl string) string {
	ii := fop.ImageInfo{}
	return ii.MakeRequest(imageUrl)
}

//5.2 删除文件
func delFile() {
	bucket := "attach"
	key := "goupload.jpg"
	var rsCli = rs.New(nil)

	err := rsCli.Delete(nil, bucket, key)
	if err != nil {
		// 产生错误
		log.Println("rs.Copy failed:", err)
		return
	}
}

//5.1 获取文件信息
func getFileInfo() {
	var ret rs.Entry
	bucket := "attach"
	key := "goupload.jpg"
	var rsCli = rs.New(nil)

	ret, err := rsCli.Stat(nil, bucket, key)

	if err != nil {
		// 产生错误
		log.Println("rs.Stat failed:", err)
		return
	}

	// 处理返回值
	log.Println(ret)
}

//上传本地文件
func upload() {
	uptoken := uptoken("attach")
	fmt.Printf("uptoken:%s\n", uptoken)

	var err error
	var ret io.PutRet
	var extra = &io.PutExtra{
	//Params:    params,
	//MimeType:  mieType,
	//Crc32:     crc32,
	//CheckCrc:  CheckCrc,
	}

	var key = "goupload.jpg"
	var localFile = "d:\\566594.jpg"

	// ret       变量用于存取返回的信息，详情见 io.PutRet
	// uptoken   为业务服务器生成的上传口令
	// key       为文件存储的标识
	// localFile 为本地文件名
	// extra     为上传文件的额外信息，详情见 io.PutExtra，可选
	err = io.PutFile(nil, &ret, uptoken, key, localFile, extra)

	if err != nil {
		//上传产生错误
		log.Print("io.PutFile failed:", err)
		return
	}

	//上传成功，处理返回值
	log.Print(ret.Hash, ret.Key)

}
