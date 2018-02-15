task "ci" {
  provider = "github-launch/microsoft-vsts-ci"
  config = ".vsts/microsoft.yml"
}
