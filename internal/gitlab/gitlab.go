package gitlab

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

type GitLabResponse struct {
    WebURL string `json:"web_url"`
}

func CreateProject(projectName string) string {
    fmt.Println("DEBUG: Creating GitLab project:", projectName)

    token := os.Getenv("GITLAB_TOKEN")
    if token == "" {
        fmt.Println("ERROR: GitLab token not set")
        return "FAILED"
    }

    url := "https://gitlab.com/api/v4/projects"

    payload := map[string]string{
        "name": projectName,
    }

    body, _ := json.Marshal(payload)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
    if err != nil {
        fmt.Println("ERROR:", err)
        return "FAILED"
    }

    req.Header.Set("PRIVATE-TOKEN", token)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("ERROR:", err)
        return "FAILED"
    }
    defer resp.Body.Close()

    fmt.Println("GitLab API status:", resp.Status)

    if resp.StatusCode != 201 {
        fmt.Println("GitLab API FAILED")
        return "FAILED"
    }

    var result GitLabResponse
    json.NewDecoder(resp.Body).Decode(&result)

    fmt.Println("GitLab project created:", result.WebURL)

    return result.WebURL
}
