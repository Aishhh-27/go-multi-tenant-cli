package k8s

import (
    "fmt"
    "os/exec"
)

func DeployHelmChart(tenantName, chartPath string) {
    cmd := exec.Command(
        "helm", "upgrade", "--install",
        tenantName,
        chartPath,
        "--namespace", tenantName,
        "--create-namespace",
    )

    output, err := cmd.CombinedOutput()

    if err != nil {
        fmt.Println("WARNING: Helm failed for", tenantName)
        fmt.Println(string(output))
        return
    }

    fmt.Println("Helm chart deployed for tenant:", tenantName)
}
