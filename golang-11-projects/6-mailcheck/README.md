# Domain Checker

This is a command-line tool written in Go for checking various DNS records of domain names. It performs MX, SPF, and DMARC record lookups for the specified domains.

## Prerequisites

- Go (version 1.15 or higher) should be installed on your system.

## Usage

1. Clone the repository or download the `main.go` file.

2. Open a terminal or command prompt and navigate to the directory where the `main.go` file is located.

3. Build the executable using the following command:

```
go build -o domain_checker main.go
```

4. Run the executable with one of the following methods:

- **Method 1: Provide domains through standard input:**

  ```
  ./domain_checker
  ```

  The program will start running, waiting for you to enter domains through standard input. Each domain should be entered on a new line. To stop the program, press `Ctrl + C`.

- **Method 2: Provide domains through an input file:**

  ```
  ./domain_checker -input input.txt
  ```

  Replace `input.txt` with the path to your input file containing domains (one domain per line). The program will read the domains from the file and process them.

- **Method 3: Specify an output file:**

  ```
  ./domain_checker -input input.txt -output output.json
  ```

  Replace `input.txt` with the path to your input file and `output.json` with the desired output file path. The program will read the domains from the input file, process them, and write the results in JSON format to the output file.

- **Method 4: Adjust concurrency level:**

  By default, the program uses a concurrency level of 10. You can adjust this value by specifying the `-concurrency` flag:

  ```
  ./domain_checker -concurrency 20
  ```

  This command will set the concurrency level to 20, allowing the program to process more domains concurrently.

5. The results will be displayed in the terminal or written to the specified output file, depending on the options you chose.

## Example Output

The program will output the results in a comma-separated format with headers:

```
domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord
example.com, true, true, v=spf1 mx -all, true, v=DMARC1; p=quarantine; rua=mailto:dmarc@example.com; ruf=mailto:dmarc@example.com; fo=1
google.com, true, true, v=spf1 include:_spf.google.com ~all, true, v=DMARC1; p=reject; rua=mailto:dmarc@google.com; ruf=mailto:dmarc@google.com; fo=1
...
```


## License

This project is licensed under the [MIT License](LICENSE).
