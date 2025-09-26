package main

import "fmt"

type Human struct {
    Name string
    Age  int
}

func (h *Human) Speak() string {
    return fmt.Sprintf("%s says hello!", h.Name)
}

func (h *Human) Walk() string {
    return fmt.Sprintf("%s is walking", h.Name)
}

func (h *Human) Fly() string {
    return fmt.Sprintf("Hi I'm %s, I'm %d, and I can't fly", h.Name, h.Age)
}

type Action struct {
    *Human

    ActionName string
}

func (a *Action) Perform() string {
    return fmt.Sprintf("%s is performing action: %s", a.Name, a.ActionName)
}

func main() {
    human := &Human{
        Name: "Daniil",
        Age:  20,
    }

    action := &Action{
        Human:      human,
        ActionName: "running",
    }
    
    fmt.Println(action.Speak())
    fmt.Println(action.Walk())
    fmt.Println(action.Fly())
    
    fmt.Println(action.Perform())

    fmt.Println(action.Name)
    fmt.Println(action.Age)

    action.Age = 31
    fmt.Println(action.Fly())
}