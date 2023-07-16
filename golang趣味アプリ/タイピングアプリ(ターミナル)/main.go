// package main

// import (
// 	"fmt"
// )

// func main() {
// 	totalScore := 0
// 	ask(1, "dog", &totalScore)
// 	ask(2, "cat", &totalScore)
// 	ask(3, "fish", &totalScore)
// 	ask(4, "Tiger", &totalScore)
// 	ask(5, "Elephant", &totalScore)
// 	ask(6, "Crocodile", &totalScore)
// 	ask(7, "this is a dog", &totalScore)
// 	ask(8, "Amphibians", &totalScore)
// 	ask(9, "Butterfly", &totalScore)
// 	ask(10, "excellentswimmer", &totalScore)

// 	fmt.Println("スコア", totalScore)
// 	if totalScore <= 30 {
// 		fmt.Println("頑張りましょう！")
// 	} else if totalScore <= 60 && totalScore >= 31 {
// 		fmt.Println("もう少しです！")
// 	} else if totalScore <= 80 && totalScore >= 61 {
// 		fmt.Println("なかなかいいですね！")
// 	} else if totalScore >= 81 {
// 		fmt.Println("素晴らしい！！")
// 	}
// }
// func ask(number int, question string, scorePtr *int) {
// 	var ans string
// 	fmt.Printf("[質問%d] 次の単語を入力してください: %s\n", number, question)
// 	fmt.Scan(&ans)
// 	if question == ans {
// 		fmt.Println("正解です！")
// 		*scorePtr += 10
// 	} else {
// 		fmt.Println("不正解です!")
// 	}
// }

package main

import (
	"bufio" // buffered「データ転送」をやるためのもの
	"flag"  // コマンドラインのフラグを解析
	"fmt"   // 文字列の入出力
	// ioパッケージ インターフェース
	"math/rand" //ランダム
	"os"        // osパッケージ
	"time"      // タイマー
)

// 配列をシャッフルする
func shuffle(data []string) {
	n := len(data)
	rand.Seed(time.Now().Unix())
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

var t int

func init() {
	//オプションで制限時間をできる
	flag.IntVar(&t, "t", 1, "制限時間")
	flag.Parse()
}

func main() {
	var (
		ch_rcv = input(os.Stdin)
		tm     = time.After(time.Duration(t) * time.Minute)
		words  = []string{"raccoon", "hogehoge"}
		// words     = []string{ "raccoon", "excellentswimmer is long language", "hogehoge" }
		score = 0
	)
	fmt.Println()
	shuffle(words)
	fmt.Println("タイピングゲームを始めます。制限時間は", t, "分。1語1点、", len(words), "点満点")
	//送信用チャネル
	num := 1
	for i := true; i && score < len(words); {
		question := words[score]
		fmt.Print("[質問", num, "]次の単語を入力してください:", question, "\n")
		fmt.Print("[答え]")
		select {
		case x := <-ch_rcv:
			//標準入力に何か入力された時の処理
			// 入力された文字が一致しているかどうかをチェックする
			if question == x {
				fmt.Println("正解です！")
				score++
				num++
			} else {
				fmt.Println("不正解です！")
			}
		case <-tm:
			//制限時間が過ぎた際の処理
			fmt.Println("\n制限時間を過ぎました")
			i = false
		}
	}
	fmt.Println("あなたの点数:", score, "点 / ", len(words), " 点")
	n := score
	switch {
	case n <= 10:
		fmt.Println("判定 F")
	case n > 45:
		fmt.Println("判定 SSS")
	default:
		fmt.Println("判定 F")
	}
}

// 入力コードを読み出すためのインターフェース
func input(Stdin *os.File) <-chan string {
	fmt.Printf("%T", Stdin)
	channel := make(chan string)
	go func() {
		strings := bufio.NewScanner(Stdin)
		for strings.Scan() {
			channel <- strings.Text()
		}
	}()
	return channel
}
