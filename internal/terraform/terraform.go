package terraform

import (
    "context"
    "fmt"
    "path/filepath"

    "github.com/hashicorp/terraform-exec/tfexec"
)

func ApplyTerraform(tenantName string) {
    dir, err := filepath.Abs("./terraform")
    if err != nil {
        panic(err)
    }

    tf, err := tfexec.NewTerraform(dir, "terraform")
    if err != nil {
        panic(err)
    }

    err = tf.Init(context.Background(), tfexec.Upgrade(true))
    if err != nil {
        panic(err)
    }

    // Workspace handling
    workspaces, _, err := tf.WorkspaceList(context.Background())
    if err != nil {
        panic(err)
    }

    exists := false
    for _, ws := range workspaces {
        if ws == tenantName {
            exists = true
            break
        }
    }

    if exists {
        err = tf.WorkspaceSelect(context.Background(), tenantName)
    } else {
        err = tf.WorkspaceNew(context.Background(), tenantName)
    }

    if err != nil {
        panic(err)
    }

    // Apply with lock disabled (for concurrency demo)
    err = tf.Apply(
        context.Background(),
        tfexec.Var("tenant_name="+tenantName),
        tfexec.Lock(false),
    )
    if err != nil {
        panic(err)
    }

    fmt.Println("Terraform applied for tenant:", tenantName)
}
