package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//输入一个url 返回原网页
func Fetch(url string) ([]byte, error) {
	client := http.Client{}

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	cookie1 := "sid=e02d4ab3-ba69-4e01-b789-c17f4c02c7b8; ec=QobhCTx9-1599189256464-ab96c6f67f411-1869936969; FSSBBIl1UgzbN7NO=5j_.eU.0RY0sxWEVJUQdOSP35uwD.RxqMK.tUbiJPJyUzzprDDebZu_XWXxPqndzsFEpqR7yF.REYr1sJD4eTeA; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1599189261,1599390392; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1599398920; FSSBBIl1UgzbN7NP=5UTs9PmeVyEQqqqm03ynTdACeBrX9y39JNPltUX0c.Se6v7RJcb75jm4be4oBLrieK9uYQrP2XMQ8Vl0OM_nBqQ1KVGxZMKNz70TZpaDTbMCt_KP.lpKh2WVHyt0qi6XALnvXGhXBzSG0xwOuw0dircF5y7KDqcaPDtcfiLGuY2GbqmBSdQiUvQeYmVcsjXpBTxht.FoSmRhFoXf.2WZAe7LnEOCrutZtjY67JD.SkWqK3iiVXAMqEshsb32EcC9OV; _efmdata=p3XiO3x27oLymrskrJPDgioD%2BbG75aorN1zZiAS5PAkImUhGVMdl0RbKnShL%2BnwZbK4VpMH4mSI%2FYCiLeqRBI3iFWIsQq%2FlBQ%2BSdP15QPco%3D; _exid=2%2FD07DmzQruaM0URWx85Cu8tZ6DrU3pEDCfMP%2FDpfKOjOXc6zJV4rpHCrOWV5LVxQUtcZ6Bwx0bgUtWNwluv%2Bg%3D%3D" // 上图找到的cookie
	req.Header.Add("cookie", cookie1)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//不要头部的信息 只需要判code信息即可
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error:status code %d", resp.StatusCode) //fmt.Errorf()
	}
	//乱码时查看教程15-1 下载 go get -g golang.org/x/text,go get -g golang.org/net/html
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//writeStringToFile(all, "lang/basic/regex/w.txt")
	//printCityList(all)
	//fmt.Printf("%s\n", all)
	return all, nil
}
