workflow "Go Things" {
  on = "push"
  resolves = ["Go Test"]
}

action "Go Test" {
  uses = "./actions/test"
}
