package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// LogEntry represents a log entry structure
type LogEntry struct {
	Level         string  `json:"level"`
	Service       string  `json:"serviceName"`
	Timestamp     float64 `json:"ts"`
	Caller        string  `json:"caller"`
	Message       string  `json:"msg"`
	Endpoint      string  `json:"endpoint"`
	Method        string  `json:"method"`
	CountryISO2   string  `json:"countryISO2"`
	UserID        string  `json:"userID"`
	PaymentMethod string  `json:"paymentMethod"`
	UserType      string  `json:"userType"`
	Error         string  `json:"error"`
	Stacktrace    string  `json:"stacktrace"`
}

var (
	levels         = []string{"debug", "info", "error"}
	services       = []string{"UserManagementService", "OrderManagementService"}
	userEndpoints  = []string{"/user/:username/order", "/user/register", "/user/login", "/user/:username"}
	orderEndpoints = []string{"/order/:username", "/order"}
	methods        = []string{"GET", "PUT", "POST"}
	countries      = []string{"ES", "IN", "US", "UK", "AU", "Other"}
	paymentMethods = []string{"VISA", "MasterCard", "Bank", "AMEX"}
	userTypes      = []string{"business", "trial", "corporate", "individual"}
	errors         = []string{"Inbound request failed with status 400", "Internal server error", "Not found", "Unauthorized"}
	successMsgs    = []string{"Request processed successfully", "Operation completed without errors", "Success"}
)

func generateLogEntry() *LogEntry {
	// set random seed
	rand.Seed(time.Now().UnixNano())

	level := levels[rand.Intn(len(levels))]
	var message string
	var err string

	if level == "error" {
		message = errors[rand.Intn(len(errors))]
		err = message
	} else {
		message = successMsgs[rand.Intn(len(successMsgs))]
		err = ""
	}

	// Select random service
	service := services[rand.Intn(len(services))]

	// select the endpoint based on the service
	var endpoint string
	if service == "UserManagementService" {
		endpoint = userEndpoints[rand.Intn(len(userEndpoints))]
	} else {
		endpoint = orderEndpoints[rand.Intn(len(orderEndpoints))]
	}

	return &LogEntry{
		Level:         level,
		Service:       service,
		Timestamp:     float64(time.Now().UnixNano()) / 1e9,
		Caller:        fmt.Sprintf("%s/main.go:%d", service, rand.Intn(100)),
		Message:       message,
		Endpoint:      endpoint,
		Method:        methods[rand.Intn(len(methods))],
		CountryISO2:   countries[rand.Intn(len(countries))],
		UserID:        fmt.Sprintf("%d", rand.Intn(100)),
		PaymentMethod: paymentMethods[rand.Intn(len(paymentMethods))],
		UserType:      userTypes[rand.Intn(len(userTypes))],
		Error:         err,
		Stacktrace:    fmt.Sprintf("%s.func%d", services[rand.Intn(len(services))], rand.Intn(10)),
	}
}

func main() {
	logsFolder := "logs"
	logFileName := "app.log"
	logFilePath := fmt.Sprintf("%s/%s", logsFolder, logFileName)

	// Ensure the logs folder exists
	if err := os.MkdirAll(logsFolder, os.ModePerm); err != nil {
		fmt.Println("Error creating logs folder:", err)
		os.Exit(1)
	}

	// Open the log file for writing
	logFile, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		os.Exit(1)
	}
	defer logFile.Close()

	// Set log output to the file
	fmt.Println("Logging to:", logFilePath)
	fmt.Fprintln(logFile, "Logs:")

	for {
		logEntry := generateLogEntry()
		logJSON, err := json.Marshal(logEntry)
		if err != nil {
			fmt.Println("Error marshaling log entry:", err)
			os.Exit(1)
		}

		// Print to console
		fmt.Println(string(logJSON))

		// Write to log file
		fmt.Fprintln(logFile, string(logJSON))
		time.Sleep(time.Second) // Adjust the delay as needed
	}
}
