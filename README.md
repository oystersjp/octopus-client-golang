# octopus-client-golang

## example

code
```
package main

import (
	"fmt"
	"github.com/yopidax/octopus-client-golang"
)

func main(){
	client := octopus.NewClient("some-cred-xxx")
	articles, err := client.SearchByKeyword("やきそば", 1)
	if err != nil {
		panic(err)
	}

	for _, v := range articles {
		fmt.Println(v.Title)
	}
}
```

result
```
$ go run yakisoba.go
【画像】通常の4倍「ペヤング超超超大盛GIGAMAX」がついに登場！！
ペヤングがガチでやべー商品を来月発売してしまう　超超超大盛GIGAMAX2142kcal
【朗報】ペヤングソースやきそば超超超大盛ＧＩＧＡＭＡＸ　一日のカロリーを400円で摂取できる！
ペヤング「超超超大盛ＧＩＧＡＭＡＸ」、爆誕。カロリー 2142kcal、食塩相当量13.5g
ペヤングがガチでやべー商品を来月発売してしまうｗｗｗｗｗｗｗｗｗ
【速報】ペヤングから2142kcalの超超超大盛 ＧＩＧＡＭＡＸが爆誕ｗｗｗｗｗｗ
ペヤング「超超超大盛ＧＩＧＡＭＡＸ」登場　2142kcal、必要湯量1300ml、食塩相当量13.5g
```
