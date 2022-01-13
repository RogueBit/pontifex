package main

/* GoLang implementation of Bruce Schneier's Solitaire Encryption
   Algorithm (http://www.counterpane.com/solitaire.html).

   Forked from https://github.com/danmux/pontifex

   Found a great write up at https://aarontoponce.org/wiki/crypto/card-ciphers/solitaire
   to help with understanding, he did a youtube as well that makes it easy to digest.

   -- r0gu3b17 (r0gu3b17@gmail.com)
   -----BEGIN PGP PUBLIC KEY BLOCK-----

   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
   =dRQa
   -----END PGP PUBLIC KEY BLOCK-----*/

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	ja     = byte(53) //Numerical position in deck of Joker A
	jb     = byte(54) //Numerical position in deck of Joker B
	maxKey = 26       //Letters in the english alphabet
)

var verbose = false //Verbose off by default

func main() {

	var keyName, inName, outName string
	flag.StringVar(&keyName, "k", "shared.key", "the shared key to use")
	flag.StringVar(&inName, "i", "in.txt", "the input")
	flag.StringVar(&outName, "o", "out.txt", "the output")
	op := flag.String("op", "enc", "[enc, dec, key, dkey] key - takes a dec of cards input and makes a key, dkey, makes a key from a default deck")
	flag.BoolVar(&verbose, "v", false, "verbose prints the deck at each stage")

	orgHelp := flag.Usage
	flag.Usage = func() {
		orgHelp()
		setupMaps(true)
	}

	flag.Parse()

	setupMaps(false)

	in, err := ioutil.ReadFile(inName)
	if err != nil {
		log.Fatal(err)
	}

	switch *op {
	case "enc":
		key := readKey(keyName)
		in = cleanInput(in)
		o := encrypt(key, in)
		err = writeChars(o, outName, true)
	case "dec":
		key := readKey(keyName)
		in = cleanInput(in)
		o := decrypt(key, in)
		err = writeChars(o, outName, false)
	case "key":
		cards, err := ioutil.ReadFile(inName)
		if err != nil {
			break
		}
		key := cardsToKey(cards)
		err = writeKey(key, keyName)
	case "dkey":
		err = defaultKey(keyName)
	default:
		log.Fatal(*op, "unsupported")
	}

	if err != nil {
		log.Fatal(err)
	}
}

func decrypt(k key, in []byte) []byte {
	sz := len(in)
	ks := k.stream(sz)
	o := make([]byte, sz)
	for i, v := range in {
		k := ks[i]
		// only generate up to our max modulus value
		if k > maxKey {
			k -= maxKey
		}
		if v < k {
			v += maxKey
		}
		o[i] = v - k
	}
	return o
}

func encrypt(k key, in []byte) []byte {
	sz := len(in)
	ks := k.stream(sz)
	if verbose {
		for _, v := range ks {
			fmt.Printf("%d ", v)
		}
	}
	o := make([]byte, sz)
	for i, v := range in {
		k := ks[i]
		// only generate up to our max modulus value
		if k > maxKey {
			k -= maxKey
		}
		o[i] = k + v
		if o[i] > maxKey {
			o[i] = o[i] - maxKey
		}
	}
	return o
}

// defaultKey just generates a default key and writes it in base64
func defaultKey(fName string) error {
	k := cardsToKey([]byte(`
		1C 2C 3C 4C 5C 6C 7C 8C 9C 10C JC QC KC
		1D 2D 3D 4D 5D 6D 7D 8D 9D 10D JD QD KD 
		1H 2H 3H 4H 5H 6H 7H 8H 9H 10H JH QH KH 
		1S 2S 3S 4S 5S 6S 7S 8S 9S 10S JS QS KS
		*A *B`))
	return writeKey(k, fName)
}

// readKey reads a base64 encoded key
func readKey(fName string) key {
	keyB64, err := ioutil.ReadFile(fName)
	if err != nil {
		log.Fatal(err)
	}

	dst := make([]byte, base64.StdEncoding.DecodedLen(len(keyB64)))
	if _, err := base64.StdEncoding.Decode(dst, keyB64); err != nil {
		log.Fatal(err)
	}

	return key(dst)
}

// writeKey writes a key to base 64
func writeKey(k key, fName string) error {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(k)))
	base64.StdEncoding.Encode(dst, k)
	return ioutil.WriteFile(fName, dst, 0644)
}

// writeChars writes the bytes out mapped to chars from 'A' onwards
// if block is true then writes 4 blocks of 5 chars per line - like old school encrypted stuff
func writeChars(d []byte, fName string, block bool) error {
	f, err := os.Create(fName)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	for i, v := range d {
		if err := w.WriteByte('A' + v - 1); err != nil {
			return err
		}
		if !block {
			continue
		}
		// block version
		if i%5 == 4 {
			if err := w.WriteByte(' '); err != nil {
				return err
			}
		}
		if i%20 == 19 {
			if err := w.WriteByte('\n'); err != nil {
				return err
			}
		}
	}
	if err := w.WriteByte('\n'); err != nil {
		return err
	}
	return w.Flush()
}

// cleanInput takes the input bytes, removes whitespace
// maps upper case chars to values where 'A' = 1
func cleanInput(in []byte) []byte {
	in = bytes.Join(bytes.Fields(in), nil)
	for i, v := range in {
		in[i] = v - 'A' + 1
	}
	pads := 4 - (len(in)-1)%5
	for i := 0; i < pads; i++ {
		in = append(in, 9+byte(i)) // pads with IJKL
	}

	return in
}

// stream generates a key stream of count size
func (k *key) stream(count int) []byte {
	r := make([]byte, count)
	tot := 0
	for tot < count {
		cardVal := k.oneRound()
		// if we land on any joker - ignore this round
		if cardVal >= ja {
			continue
		}
		r[tot] = cardVal
		tot++
	}

	return r
}

// oneRound generates one key round
func (k *key) oneRound() byte {

	sz := len(*k)

	if verbose {
		fmt.Println("\nstart")
		k.log()
	}

	// move A joker down 1
	k.move(ja, 1)

	if verbose {
		fmt.Println("\nmoved ja")
		k.log()
	}

	// move B joker down 2
	k.move(jb, 2)

	if verbose {
		fmt.Println("\nmoved jb")
		k.log()
	}

	// find the top most joker
	top := k.findVal(ja)
	bot := k.findVal(jb)
	if top >= bot {
		b := bot
		bot = top
		top = b
	}

	if verbose {
		fmt.Printf("top:%d, bot:%d\n", top, bot)
	}

	// tripple cut into new key
	newKey := make([]byte, sz)

	p := 0
	for i := bot + 1; i < sz; i++ {
		newKey[p] = (*k)[i]
		p++
	}
	for i := top; i <= bot; i++ {
		newKey[p] = (*k)[i]
		p++
	}
	for i := 0; i < top; i++ {
		newKey[p] = (*k)[i]
		p++
	}

	if verbose {
		fmt.Println("\ntripple cut")
		key(newKey).log()
	}

	// get the value of the bottom of the deck for the count cut
	// handling jokers
	count := enforceMaxCardVal(newKey[sz-1])

	if verbose {
		fmt.Println("\ncount cut:", count)
	}

	// the count cut copies from the newKey above back into the original key
	p = 0
	// from the count onwards becomes the top
	for i := int(count); i < sz-1; i++ {
		(*k)[p] = newKey[i]
		p++
	}
	// and the top becomes the bottom - up to the last card which stays the same
	for i := 0; i < int(count); i++ {
		(*k)[p] = newKey[i]
		p++
	}
	// put the last card back in place
	(*k)[sz-1] = newKey[sz-1]

	if verbose {
		fmt.Println("\nafter count cut")
		k.log()
	}

	// get the value of the top card to count down
	tc := (*k)[0]
	if verbose {
		fmt.Println("\ntop card", valToCard[tc], tc)
	}
	// handle jokers
	tc = enforceMaxCardVal(tc)

	// get the value of this card
	val := (*k)[tc]

	if verbose {
		fmt.Println("\nkey val: ", valToCard[val], val)
	}
	return val
}

// enforceMaxCardVal simply makes sure both jokers are counted as 53
func enforceMaxCardVal(in byte) byte {
	if in > 52 {
		return 53 // both jokers value 53
	}
	return in
}

// cardsToKey converts the string based deck format to a key
func cardsToKey(cards []byte) key {
	r := key{}
	cs := bytes.Fields(cards)
	for _, c := range cs {
		r = append(r, cardToVal[string(c)])
	}
	return r
}

// key models the deck based key
type key []byte

// findVal returns the index of the card with value val
func (k key) findVal(val byte) int {
	for i, v := range k {
		if val == v {
			return i
		}
	}
	return -1
}

// move moves one card forward in the key deck
func (k key) move(val byte, off int) {
	sz := len(k)
	pos := k.findVal(val)

	// if this will wrap around then becomes an insert
	if pos+off >= sz {
		// slide the deck down one between the old and new positions
		insert := (pos + off) % sz
		for i := pos; i > insert; i-- {
			k[i] = k[i-1]
		}
		// and set the new position to val
		k[insert+1] = val
		return
	}

	// otherwise move everything up one and set the new position to val
	for i := 0; i < off; i++ {
		k[pos+i] = k[pos+i+1]
	}
	k[pos+off] = val
	return
}

// log prints the deck out
func (k key) log() {
	for i, v := range k {
		if i%13 == 0 {
			println()
		}
		print(valToCard[v], " ")
	}
	println()
}

// some maps to go from and to card string to their representative value
var (
	cardToVal = map[string]byte{}
	valToCard = map[byte]string{}
)

// setupMaps sets up the maps from card strings to values and visa versa
func setupMaps(show bool) {
	if show {
		fmt.Fprintln(os.Stderr, "    Card Values:")
	}
	val := byte(1)
	// bridge clubs, diamonds, hearts, and spades
	for _, suit := range []string{"C", "D", "H", "S"} {
		for _, card := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"} {
			c := card + suit
			if show {
				fmt.Fprintf(os.Stderr, "    %s\t%d\n", c, int(val))
			}
			cardToVal[c] = val
			valToCard[val] = c
			val++
		}
	}
	j := "*A"
	cardToVal[j] = ja
	valToCard[ja] = j

	j = "*B"
	cardToVal[j] = jb
	valToCard[jb] = j
}
