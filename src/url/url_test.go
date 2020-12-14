package url

import (
	"fmt"
	"net/url"
	"testing"
)

func TestUrl(t *testing.T) {
	URL, _ := url.Parse("http://localhost:8887/v1/api/cloud_api/fetcher?email=1156143589@qq.com")
	fmt.Println("scheme", URL.Scheme)
	fmt.Println("host", URL.Host)
	fmt.Println("port", URL.Port())
	fmt.Println("rawQuery", URL.RawQuery)
	fmt.Println("path", URL.Path)
	fmt.Println("forceQuery", URL.ForceQuery)
	fmt.Println("fragment", URL.Fragment)

	escapeUrl := url.QueryEscape("https://link.zhihu.com/?target=http%3A//www.flysnow.org/2017/08/26/go-1-9-type-alias.html")
	fmt.Println(escapeUrl)

	unescapeUrl, _ := url.QueryUnescape("https://link.zhihu.com/?target=http%3A//www.flysnow.org/2017/08/26/go-1-9-type-alias.html")
	fmt.Println(unescapeUrl)
}
