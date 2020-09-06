package parser

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestParserProfile(t *testing.T) {
	client := http.Client{}

	req, _ := http.NewRequest(http.MethodGet, "https://album.zhenai.com/u/1608140885", nil)
	cookie1 := "sid=e02d4ab3-ba69-4e01-b789-c17f4c02c7b8; ec=QobhCTx9-1599189256464-ab96c6f67f411-1869936969; FSSBBIl1UgzbN7NO=5j_.eU.0RY0sxWEVJUQdOSP35uwD.RxqMK.tUbiJPJyUzzprDDebZu_XWXxPqndzsFEpqR7yF.REYr1sJD4eTeA; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1599189261,1599390392; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1599398774; FSSBBIl1UgzbN7NP=5UTs9rTeVyq3qqqm03yjeOaw9_O0jif1TLNpgvyMLkrnkDZumjABcpBrn7wev3xxte2JaoLWslwHvxGNk8o_tiF_xc5mAD7mLWpK4XJQJ032Gn1ueixmK_Mav2KwYBfgu.66sToMeWbESAnU3wpLM4Cl3quQCRQJVWB3liRNypD7K1ONafX3zpFVfYEBoR6shStR3QfzU0LwErlCH50hhfFL1tP94JQ6kZ.vLW0HgJEcGpDYOTYEa4pg3H_5ENyNXE; _efmdata=p3XiO3x27oLymrskrJPDgioD%2BbG75aorN1zZiAS5PAkImUhGVMdl0RbKnShL%2BnwZ0YQcJ%2BDwm0L3sf5MbwJduUv%2FiOlk8OUiMi%2FNOfwVJhs%3D; _exid=q2VC2etoLfj%2BfvD5JRB9SBnI7uEdreGZvXePSBmEYbKqvc7LgArj8kGtHrlV8cYsANqFEqXVA6H8Wqyxb0kYJw%3D%3D" // 上图找到的cookie
	req.Header.Add("cookie", cookie1)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	//不要头部的信息 只需要判断头部信息即可
	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("error:status code %d", resp.StatusCode) //fmt.Errorf()

	}
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	//file, err := ioutil.ReadFile("profile_data.html")
	/*	if err != nil {
		panic(err)
	}*/
	ParserProfile(all, "姓名")

	//fmt.Println(result)
}
