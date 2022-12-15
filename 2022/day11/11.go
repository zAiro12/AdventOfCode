package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

type scimmia struct {
	oggetti           []int
	isMoltiplicazione bool
	operazione        int
	test              int
	ifTrue            string
	ifFalse           string
	ispezione         int
}

func main() {

	scimmieA := make(map[string]scimmia)
	scimmieB := make(map[string]scimmia)

	file := zairo.Input(os.Args[1])
	sc := bufio.NewScanner(file)

	var in string
	var split []string

	var oggetti []int
	var isMoltiplicazione bool
	var operazione, test int
	var ifTrue, ifFalse string

	for sc.Scan() {

		in = sc.Text()
		split = strings.Split(in, " ")
		split[1] = strings.Trim(split[1], ":")

		sc.Scan()
		in = sc.Text()
		in = strings.Trim(in, "  Starting items: ")
		oggetti = aggiungiVal(in)

		sc.Scan()
		in = sc.Text()
		in = strings.Trim(in, "  Operation: new = old ")
		if in[0] == '*' {
			isMoltiplicazione = true
			in = strings.Trim(in, "* ")
		} else {
			isMoltiplicazione = false
			in = strings.Trim(in, "+ ")
		}

		operazione, _ = strconv.Atoi(in)

		sc.Scan()
		in = sc.Text()
		in = strings.Trim(in, "  Test: divisible by ")
		test, _ = strconv.Atoi(in)

		sc.Scan()
		in = sc.Text()
		in = strings.Trim(in, "    If true: throw to monkey ")
		ifTrue = in

		sc.Scan()
		in = sc.Text()
		in = strings.Trim(in, "    If false: throw to monkey ")
		ifFalse = in

		scimmieA[split[1]] = scimmia{oggetti: oggetti, isMoltiplicazione: isMoltiplicazione, operazione: operazione, test: test, ifTrue: ifTrue, ifFalse: ifFalse, ispezione: 0}
		scimmieB[split[1]] = scimmia{oggetti: oggetti, isMoltiplicazione: isMoltiplicazione, operazione: operazione, test: test, ifTrue: ifTrue, ifFalse: ifFalse, ispezione: 0}

		sc.Scan()
	}

	zairo.StampaA(parteA(scimmieA))
	zairo.StampaB(parteB(scimmieB))
}

func aggiungiVal(s string) []int {
	val := strings.Split(s, " ")
	var num int
	var oggetti []int
	for i := 0; i < len(val); i++ {
		val[i] = strings.Trim(val[i], ",")

		num, _ = strconv.Atoi(val[i])
		oggetti = append(oggetti, num)
	}
	return oggetti
}

func parteA(scimmie map[string]scimmia) int {
	for i := 0; i < 20; i++ {

		for scimmiaAnalizzata := 0; scimmiaAnalizzata < len(scimmie); scimmiaAnalizzata++ {
			nomeScimmia := strconv.Itoa(scimmiaAnalizzata)

			for _, oggetto := range scimmie[nomeScimmia].oggetti {

				var nuovoOggetto int
				if scimmie[nomeScimmia].isMoltiplicazione {
					if scimmie[nomeScimmia].operazione == 0 {
						nuovoOggetto = oggetto * oggetto
					} else {
						nuovoOggetto = oggetto * scimmie[nomeScimmia].operazione
					}
				} else {
					if scimmie[nomeScimmia].operazione == 0 {
						nuovoOggetto = oggetto + oggetto
					} else {
						nuovoOggetto = oggetto + scimmie[nomeScimmia].operazione
					}
				}

				nuovoOggetto = nuovoOggetto / 3

				if nuovoOggetto%scimmie[nomeScimmia].test == 0 {

					appoggio := scimmie[scimmie[nomeScimmia].ifTrue].oggetti
					appoggio = append(appoggio, nuovoOggetto)
					scimmie[scimmie[nomeScimmia].ifTrue] = scimmia{
						oggetti:           appoggio,
						isMoltiplicazione: scimmie[scimmie[nomeScimmia].ifTrue].isMoltiplicazione,
						operazione:        scimmie[scimmie[nomeScimmia].ifTrue].operazione,
						test:              scimmie[scimmie[nomeScimmia].ifTrue].test,
						ifTrue:            scimmie[scimmie[nomeScimmia].ifTrue].ifTrue,
						ifFalse:           scimmie[scimmie[nomeScimmia].ifTrue].ifFalse,
						ispezione:         scimmie[scimmie[nomeScimmia].ifTrue].ispezione}

				} else {
					appoggio := scimmie[scimmie[nomeScimmia].ifFalse].oggetti
					appoggio = append(appoggio, nuovoOggetto)
					scimmie[scimmie[nomeScimmia].ifFalse] = scimmia{
						oggetti:           appoggio,
						isMoltiplicazione: scimmie[scimmie[nomeScimmia].ifFalse].isMoltiplicazione,
						operazione:        scimmie[scimmie[nomeScimmia].ifFalse].operazione,
						test:              scimmie[scimmie[nomeScimmia].ifFalse].test,
						ifTrue:            scimmie[scimmie[nomeScimmia].ifFalse].ifTrue,
						ifFalse:           scimmie[scimmie[nomeScimmia].ifFalse].ifFalse,
						ispezione:         scimmie[scimmie[nomeScimmia].ifFalse].ispezione}
				}

				scimmie[nomeScimmia] = scimmia{
					oggetti:           scimmie[nomeScimmia].oggetti,
					isMoltiplicazione: scimmie[nomeScimmia].isMoltiplicazione,
					operazione:        scimmie[nomeScimmia].operazione,
					test:              scimmie[nomeScimmia].test,
					ifTrue:            scimmie[nomeScimmia].ifTrue,
					ifFalse:           scimmie[nomeScimmia].ifFalse,
					ispezione:         scimmie[nomeScimmia].ispezione + 1}

			}
			scimmie[nomeScimmia] = scimmia{
				oggetti:           []int{},
				isMoltiplicazione: scimmie[nomeScimmia].isMoltiplicazione,
				operazione:        scimmie[nomeScimmia].operazione,
				test:              scimmie[nomeScimmia].test,
				ifTrue:            scimmie[nomeScimmia].ifTrue,
				ifFalse:           scimmie[nomeScimmia].ifFalse,
				ispezione:         scimmie[nomeScimmia].ispezione}

		}
	}

	var iterazioni []int
	for _, v := range scimmie {
		iterazioni = append(iterazioni, v.ispezione)
	}

	sort.Ints(iterazioni)
	zairo.ReverseSlice(iterazioni)
	return iterazioni[0] * iterazioni[1]
}

func parteB(scimmie map[string]scimmia) int {
	numeroMagico := 1
	for _, v := range scimmie {
		numeroMagico *= v.test

	}

	for i := 0; i < 10000; i++ {

		for scimmiaAnalizzata := 0; scimmiaAnalizzata < len(scimmie); scimmiaAnalizzata++ {
			nomeScimmia := strconv.Itoa(scimmiaAnalizzata)

			for _, oggetto := range scimmie[nomeScimmia].oggetti {

				var nuovoOggetto int
				if scimmie[nomeScimmia].isMoltiplicazione {
					if scimmie[nomeScimmia].operazione == 0 {
						nuovoOggetto = oggetto * oggetto
					} else {
						nuovoOggetto = oggetto * scimmie[nomeScimmia].operazione
					}
				} else {
					if scimmie[nomeScimmia].operazione == 0 {
						nuovoOggetto = oggetto + oggetto
					} else {
						nuovoOggetto = oggetto + scimmie[nomeScimmia].operazione
					}
				}

				nuovoOggetto = nuovoOggetto % numeroMagico

				if nuovoOggetto%scimmie[nomeScimmia].test == 0 {

					appoggio := scimmie[scimmie[nomeScimmia].ifTrue].oggetti
					appoggio = append(appoggio, nuovoOggetto)
					scimmie[scimmie[nomeScimmia].ifTrue] = scimmia{
						oggetti:           appoggio,
						isMoltiplicazione: scimmie[scimmie[nomeScimmia].ifTrue].isMoltiplicazione,
						operazione:        scimmie[scimmie[nomeScimmia].ifTrue].operazione,
						test:              scimmie[scimmie[nomeScimmia].ifTrue].test,
						ifTrue:            scimmie[scimmie[nomeScimmia].ifTrue].ifTrue,
						ifFalse:           scimmie[scimmie[nomeScimmia].ifTrue].ifFalse,
						ispezione:         scimmie[scimmie[nomeScimmia].ifTrue].ispezione}

				} else {
					appoggio := scimmie[scimmie[nomeScimmia].ifFalse].oggetti
					appoggio = append(appoggio, nuovoOggetto)
					scimmie[scimmie[nomeScimmia].ifFalse] = scimmia{
						oggetti:           appoggio,
						isMoltiplicazione: scimmie[scimmie[nomeScimmia].ifFalse].isMoltiplicazione,
						operazione:        scimmie[scimmie[nomeScimmia].ifFalse].operazione,
						test:              scimmie[scimmie[nomeScimmia].ifFalse].test,
						ifTrue:            scimmie[scimmie[nomeScimmia].ifFalse].ifTrue,
						ifFalse:           scimmie[scimmie[nomeScimmia].ifFalse].ifFalse,
						ispezione:         scimmie[scimmie[nomeScimmia].ifFalse].ispezione}
				}

				scimmie[nomeScimmia] = scimmia{
					oggetti:           scimmie[nomeScimmia].oggetti,
					isMoltiplicazione: scimmie[nomeScimmia].isMoltiplicazione,
					operazione:        scimmie[nomeScimmia].operazione,
					test:              scimmie[nomeScimmia].test,
					ifTrue:            scimmie[nomeScimmia].ifTrue,
					ifFalse:           scimmie[nomeScimmia].ifFalse,
					ispezione:         scimmie[nomeScimmia].ispezione + 1}

			}
			scimmie[nomeScimmia] = scimmia{
				oggetti:           []int{},
				isMoltiplicazione: scimmie[nomeScimmia].isMoltiplicazione,
				operazione:        scimmie[nomeScimmia].operazione,
				test:              scimmie[nomeScimmia].test,
				ifTrue:            scimmie[nomeScimmia].ifTrue,
				ifFalse:           scimmie[nomeScimmia].ifFalse,
				ispezione:         scimmie[nomeScimmia].ispezione}

		}
	}

	var iterazioni []int
	for _, v := range scimmie {
		iterazioni = append(iterazioni, v.ispezione)
	}

	sort.Ints(iterazioni)
	zairo.ReverseSlice(iterazioni)
	return iterazioni[0] * iterazioni[1]
}

func stampaScimmie(scimmie map[string]scimmia) {
	for k, v := range scimmie {
		fmt.Println(k, ":", v.oggetti)
	}
}
