package client

import (
	"bytes"
	"fmt"
	"gofer/config"
	"gofer/schemas"
	"net/http"
	"os"
	"text/tabwriter"
)

func sendApiRequest(method string, path string, body []byte, params map[string]string) (*http.Response, error) {
	apiUrl := config.GetConfig().Client.ApiUrl

	req, err := http.NewRequest(method, apiUrl+path, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// API key for authentication
	apiKey := os.Getenv("GOFER_API_KEY")
	req.Header.Set("x-api-key", apiKey)

	// Add query parameters to URL object
	query := req.URL.Query()
	for key, value := range params {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func printTasks(tasks []schemas.TaskResponse) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(w, "Task ID\tTitle\tDescription\tCompleted\tCreated By\tCreated At\tUpdated At")
	fmt.Fprintln(w, "-------\t-----\t-----------\t----------\t----------\t-----------\t-----------")

	for _, task := range tasks {
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t%v\n",
			task.TaskId,
			task.Title,
			task.Description,
			task.Completed,
			task.CreatedBy,
			task.CreatedAt.Local().Format("2006-01-02 15:04:05"),
			task.UpdatedAt.Local().Format("2006-01-02 15:04:05"),
		)
	}

	w.Flush()
}
