package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "tenant-cli",
    Short: "Provision multi-tenant environments",
    Long:  "CLI to automate Kubernetes namespaces and Helm charts per tenant",
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
    }
}
