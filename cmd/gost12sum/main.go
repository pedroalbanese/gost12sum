package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pedroalbanese/gogost/gost34112012256"
	"github.com/pedroalbanese/gogost/gost34112012512"
)

var (
	check     = flag.String("c", "", "Check hashsum file")
	long      = flag.Bool("l", false, "Use 512 bit hash (default 256-bit)")
	recursive = flag.Bool("r", false, "Process directories recursively")
	verbose   = flag.Bool("v", false, "Verbose mode (for CHECK command)")
)

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "GOST12SUM(2) Copyright (c) 2020-2021 ALBANESE Research Lab")
		fmt.Fprintln(os.Stderr, "GOST R 34.11-2012 Streebog 256/512-bit Recursive Hasher\n")
		fmt.Fprintln(os.Stderr, "Usage of", os.Args[0]+":")
		fmt.Fprintf(os.Stderr, "%s [-v] [-c <hash.g12>] [-r] [-l] <file.ext>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	Files := strings.Join(flag.Args(), " ")

	if Files == "-" {
		var h hash.Hash
		if *long == false {
			h = gost34112012256.New()
		} else if *long == true {
			h = gost34112012512.New()
		}
		io.Copy(h, os.Stdin)
		fmt.Println(hex.EncodeToString(h.Sum(nil)) + " (stdin)")
		os.Exit(0)
	}

	if strings.Contains(Files, "*") && *check == "" && *recursive == false {
		files, err := filepath.Glob(Files)
		if err != nil {
			log.Fatal(err)
		}
		for _, match := range files {
			var h hash.Hash
			if *long == false {
				h = gost34112012256.New()
			} else if *long == true {
				h = gost34112012512.New()
			}
			f, err := os.Open(match)
			if err != nil {
				log.Fatal(err)
			}
			file, err := os.Stat(match)
			if file.IsDir() {
			} else {
				if _, err := io.Copy(h, f); err != nil {
					log.Fatal(err)
				}
				fmt.Println(hex.EncodeToString(h.Sum(nil)), "*"+f.Name())
			}
			f.Close()
		}
		os.Exit(0)
	}

	if *check == "" && *recursive == false {
		for _, match := range flag.Args() {
			var h hash.Hash
			if *long == false {
				h = gost34112012256.New()
			} else if *long == true {
				h = gost34112012512.New()
			}
			f, err := os.Open(match)
			if err != nil {
				log.Fatal(err)
			}
			file, err := os.Stat(match)
			if file.IsDir() {
			} else {
				if _, err := io.Copy(h, f); err != nil {
					log.Fatal(err)
				}
				fmt.Println(hex.EncodeToString(h.Sum(nil)), "*"+f.Name())
			}
			f.Close()
		}
		os.Exit(0)
	}

	if *check == "" && *recursive == true {
		err := filepath.Walk(filepath.Dir(Files),
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				file, err := os.Stat(path)
				if file.IsDir() {
				} else {
					for _, match := range flag.Args() {
						filename := filepath.Base(path)
						pattern := filepath.Base(match)
						matched, err := filepath.Match(pattern, filename)
						if err != nil {
							fmt.Println(err)
						}
						if matched {
							var h hash.Hash
							if *long == false {
								h = gost34112012256.New()
							} else if *long == true {
								h = gost34112012512.New()
							}
							f, err := os.Open(path)
							if err != nil {
								log.Fatal(err)
							}
							if _, err := io.Copy(h, f); err != nil {
								log.Fatal(err)
							}
							f.Close()
							fmt.Println(hex.EncodeToString(h.Sum(nil)), "*"+f.Name())
						}
					}
				}
				return nil
			})
		if err != nil {
			log.Println(err)
		}
	}

	if *check != "" {
		var file io.Reader
		var err error
		if *check == "-" {
			file = os.Stdin
		} else {
			file, err = os.Open(*check)
			if err != nil {
				log.Fatalf("failed opening file: %s", err)
			}
		}

		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var txtlines []string

		for scanner.Scan() {
			txtlines = append(txtlines, scanner.Text())
		}

		for _, eachline := range txtlines {
			lines := strings.Split(string(eachline), " *")

			if strings.Contains(string(eachline), " *") {
				var h hash.Hash
				if *long == false {
					h = gost34112012256.New()
				} else if *long == true {
					h = gost34112012512.New()
				}
				_, err := os.Stat(lines[1])
				if err == nil {
					f, err := os.Open(lines[1])
					if err != nil {
						log.Fatal(err)
					}
					io.Copy(h, f)
					f.Close()

					if *verbose {
						if hex.EncodeToString(h.Sum(nil)) == lines[0] {
							fmt.Println(lines[1]+"\t", "OK")
						} else {
							fmt.Println(lines[1]+"\t", "FAILED")
						}
					} else {
						if hex.EncodeToString(h.Sum(nil)) == lines[0] {
						} else {
							os.Exit(1)
						}
					}
				} else {
					if *verbose {
						fmt.Println(lines[1]+"\t", "Not found!")
					} else {
						os.Exit(1)
					}
				}
			}
		}
	}
}
