package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

type DomainInfo struct {
	Domain       string `json:"domain"`
	HasMX        bool   `json:"hasMX"`
	HasSPF       bool   `json:"hasSPF"`
	SPFRecord    string `json:"spfRecord"`
	HasDMARC     bool   `json:"hasDMARC"`
	DMARCRecord  string `json:"dmarcRecord"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func main() {
	var inputFile string
	var outputFile string
	var concurrency int

	flag.StringVar(&inputFile, "input", "", "Input file containing domains (one per line)")
	flag.StringVar(&outputFile, "output", "", "Output file to write results (in JSON format)")
	flag.IntVar(&concurrency, "concurrency", 10, "Number of concurrent workers")

	flag.Parse()

	var domains []string

	if inputFile != "" {
		file, err := os.Open(inputFile)
		if err != nil {
			log.Fatalf("Error opening input file: %v", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			domains = append(domains, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatalf("Error reading input file: %v", err)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			domains = append(domains, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatalf("Error: could not read from input: %v", err)
		}
	}

	if len(domains) == 0 {
		log.Fatal("No domains provided")
	}

	results := make([]DomainInfo, len(domains))
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, concurrency)

	for i, domain := range domains {
		wg.Add(1)
		semaphore <- struct{}{}
		go func(index int, d string) {
			defer func() {
				<-semaphore
				wg.Done()
			}()

			info := checkDomain(d)
			results[index] = info
		}(i, domain)
	}

	wg.Wait()

	if outputFile != "" {
		file, err := os.Create(outputFile)
		if err != nil {
			log.Fatalf("Error creating output file: %v", err)
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(results); err != nil {
			log.Fatalf("Error writing output file: %v", err)
		}
	} else {
		printResults(results)
	}
}

func checkDomain(domain string) DomainInfo {
	var info DomainInfo

	info.Domain = domain

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		info.ErrorMessage = fmt.Sprintf("Error looking up MX records: %v", err)
		return info
	}

	if len(mxRecords) > 0 {
		info.HasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		info.ErrorMessage = fmt.Sprintf("Error looking up TXT records: %v", err)
		return info
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			info.HasSPF = true
			info.SPFRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		info.ErrorMessage = fmt.Sprintf("Error looking up DMARC records: %v", err)
		return info
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			info.HasDMARC = true
			info.DMARCRecord = record
			break
		}
	}

	return info
}

func printResults(results []DomainInfo) {
	fmt.Printf("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord\n")
	for _, info := range results {
		fmt.Printf("%v, %v, %v, %v, %v, %v\n", info.Domain, info.HasMX, info.HasSPF, info.SPFRecord, info.HasDMARC, info.DMARCRecord)
	}
}
