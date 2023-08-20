package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

var t int

func init() {
	//オプションで制限時間をできる
	flag.IntVar(&t, "t", 1, "制限時間")
	flag.Parse()
}

func shuffle(data []string) {
	n := len(data) - 1
	rand.Seed(time.Now().UnixNano())
	for i := n; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

func input(Stdin io.Reader) <-chan string {
	channel := make(chan string)
	go func() {
		strings := bufio.NewScanner(Stdin)
		for strings.Scan() {
			channel <- strings.Text()
		}
	}()
	return channel
}

func main() {
	// ゲームの初期設定
	var (
		// questions := []string{"apple", "banana", "cherry", "dog", "elephant", "flower", "giraffe", "hamburger", "ice cream", "juice", "kangaroo", "lion", "monkey", "orange", "pepper", "queen", "rabbit", "snake", "tiger", "umbrella", "violin", "watermelon", "xylophone", "yellow", "zebra"}
		questions = []string{"dog", "test"}
		ch_rcv    = input(os.Stdin)
		tm        = time.After(time.Duration(t) * time.Minute)
		// scanner   = bufio.NewScanner(os.Stdin)
		score   = 0
		penalty = 0
		num     = 1
	)

	shuffle(questions)
	fmt.Println()
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("タイピングゲームを始めます。制限時間は", t, "分。1語1点、", len(questions), "点満点")
	fmt.Println("----------------------------------------------------------------")
	fmt.Println()
	End:
	// 無限ループでゲームを続ける
	for score < len(questions) {
		question := questions[score]
		// var input string
		fmt.Printf("[質問%d] 次の単語を入力してください: %s\n", num, question)
		fmt.Print("[答え]")

		select {
		case input := <-ch_rcv:
			//標準入力に何か入力された時の処理
			// 入力された文字が一致しているかどうかをチェックする
			if input == question {
				fmt.Println(" ----")
				fmt.Println("|正解|")
				fmt.Println(" ----")
				fmt.Println()
				score++
				num++
			} else {
				// 入力が間違っていた場合はペナルティを加算する
				fmt.Println(" ------")
				fmt.Println("|不正解|")
				fmt.Println(" ------")
				fmt.Println()
				penalty++
			}
		case <-tm:
			//制限時間が過ぎた際の処理
			fmt.Println("\n制限時間を過ぎました")
			fmt.Println("----------------------------------------------------------------")
			fmt.Println("再チャレンジしますか? (y/n)")
			fmt.Println("----------------------------------------------------------------")
			select {
			case inp := <-ch_rcv:
				// ゲームを続けるかどうかをユーザーに尋ねる
				if inp == "n" {
					break End
				}
			}
			fmt.Println("----------------------------------------------------------------")
			fmt.Println("タイピングゲームを始めます。制限時間は", t, "分。1語1点、", len(questions), "点満点")
			fmt.Println("----------------------------------------------------------------")
			tm = time.After(time.Duration(t) * time.Minute)
		}
	}

	// ゲーム終了時のスコアを表示する
	finalScore := score - penalty

	fmt.Println("----------------------------------------------------------------")
	fmt.Printf("加点：%d  減点：%d\n", score, penalty)
	fmt.Printf("最終スコア%d点  ", finalScore)
	if finalScore <= 30 {
		fmt.Println("[頑張りましょう！]")
	} else if finalScore <= 60 && finalScore >= 31 {
		fmt.Println("[もう少しです！]")
	} else if finalScore <= 80 && finalScore >= 61 {
		fmt.Println("[なかなかいいですね！]")
	} else if finalScore >= 81 {
		fmt.Println("[素晴らしい！！]")
	}
	fmt.Println("----------------------------------------------------------------")
}
