package utils

import (
    "fmt"
    "os"
)

func GenerateReport(tenantName, namespace, url string) {
    fmt.Println("DEBUG: Starting report generation")

    fileName := fmt.Sprintf("%s-report.txt", tenantName)

    f, err := os.Create(fileName)
    if err != nil {
        panic(fmt.Errorf("failed to create report: %v", err))
    }
    defer f.Close()

    content := fmt.Sprintf(
        "Tenant: %s\nNamespace: %s\nURL: %s\n",
        tenantName,
        namespace,
        url,
    )

    _, err = f.WriteString(content)
    if err != nil {
        panic(fmt.Errorf("failed to write report: %v", err))
    }

    fmt.Println("DEBUG: Report generated:", fileName)
}
