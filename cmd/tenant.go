package cmd

import (
    "fmt"
    "os"
    "sync"

    "github.com/spf13/cobra"
    "gopkg.in/yaml.v2"

    "github.com/Aishhh-27/go-multi-tenant-cli/internal/gitlab"
    "github.com/Aishhh-27/go-multi-tenant-cli/internal/k8s"
    "github.com/Aishhh-27/go-multi-tenant-cli/internal/terraform"
    "github.com/Aishhh-27/go-multi-tenant-cli/internal/utils"
)

type Tenant struct {
    Name string `yaml:"name"`
}

type Config struct {
    Tenants []Tenant `yaml:"tenants"`
}

var filePath string

var tenantCmd = &cobra.Command{
    Use:   "create",
    Short: "Create tenants from YAML file",
    Run: func(cmd *cobra.Command, args []string) {

        data, err := os.ReadFile(filePath)
        if err != nil {
            panic(err)
        }

        var config Config
        err = yaml.Unmarshal(data, &config)
        if err != nil {
            panic(err)
        }

        var wg sync.WaitGroup

        for _, t := range config.Tenants {
            wg.Add(1)

            go func(tenantName string) {
                defer wg.Done()

                // Recover from panic inside goroutine
                defer func() {
                    if r := recover(); r != nil {
                        fmt.Println("ERROR for tenant:", tenantName)
                        fmt.Println("REASON:", r)
                    }
                }()

                fmt.Println("Provisioning:", tenantName)

                // ✅ Step 1: Create GitLab project
                gitlabURL := gitlab.CreateProject(tenantName)

                // ✅ Step 2: Terraform
                terraform.ApplyTerraform(tenantName)

                // ✅ Step 3: Helm
                k8s.DeployHelmChart(tenantName, "./charts/gitlab")

                // ✅ Step 4: Report
                utils.GenerateReport(tenantName, tenantName, gitlabURL)

            }(t.Name)
        }

        wg.Wait()

        fmt.Println("All tenants provisioned successfully")
    },
}

func init() {
    tenantCmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to tenants YAML")
    tenantCmd.MarkFlagRequired("file")
    rootCmd.AddCommand(tenantCmd)
}
