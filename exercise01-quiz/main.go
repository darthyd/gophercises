package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Question struct {
	Question string
	Answer   string
}

func main() {
	shuffled := flag.Bool("s", false, "Shuffle the questions")
	timeLimit := flag.Int("t", 30, "Time limit in seconds")
	flag.Parse()

	csvFile, err := os.Open("problems.csv")
	defer func(csvFile *os.File) {
		err := csvFile.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(csvFile)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("\nIt's time for some math! "+
		"\nHow many answers you can give correctly in", *timeLimit, "seconds."+
		"\nLet's figured it out!"+
		"\n\nPress enter to start.")
	_, err = fmt.Scanln()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	correct := 0
	answered := 0
	points := 0

	ticker := time.NewTicker(time.Duration(*timeLimit) * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println(
					"\nYou answered correctly", correct, "of", answered,
					"questions and made", points, "points.",
				)
				os.Exit(0)
			}
		}
	}()

	if *shuffled {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(csvLines), func(i, j int) { csvLines[i], csvLines[j] = csvLines[j], csvLines[i] })
	}

	for _, columns := range csvLines {
		q := Question{
			Question: columns[0],
			Answer:   columns[1],
		}
		fmt.Println(q.Question)
		var answer string

		_, err := fmt.Scanf("%s", &answer)
		if err != nil {
			continue
		}

		answer = strings.ToLower(answer)
		answer = strings.TrimSpace(answer)

		if answer == q.Answer {
			answered++
			correct++
			points += 5
		} else {
			answered++
			points -= 2
		}
	}
}
