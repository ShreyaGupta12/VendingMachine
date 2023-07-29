package main

import (
    "fmt"
)

type Item struct {
    name  string
    price int
}

type VendingMachine struct {
    items      []Item
    stocks     map[string]int
    balance    int
    records    map[string]int
    itemPrices map[string]int
}

func NewVendingMachine() *VendingMachine {
    vm := &VendingMachine{
        items: []Item{
            {name: "Coke", price: 100},
            {name: "Pepsi", price: 90},
            {name: "Sprite", price: 80},
        },
        stocks:     make(map[string]int),
        balance:    0,
        records:    make(map[string]int),
        itemPrices: make(map[string]int),
    }
    for _, item := range vm.items {
        vm.stocks[item.name] = 10
        vm.itemPrices[item.name] = item.price
    }
    return vm
}

func (vm *VendingMachine) DisplayItems() {
    fmt.Println("--- Items Available ---")
    for _, item := range vm.items {
        fmt.Printf("%s (Price: %d)\n", item.name, item.price)
    }
    fmt.Println("-----------------------")
}

func (vm *VendingMachine) InsertCoin(amount int) {
    vm.balance += amount
    fmt.Printf("Inserted %d cents. Current balance: %d cents.\n", amount, vm.balance)
}

func (vm *VendingMachine) SelectItem(name string) {
    if vm.stocks[name] == 0 {
        fmt.Printf("Sorry, %s is out of stock.\n", name)
        return
    }
    price, ok := vm.itemPrices[name]
    if !ok {
        fmt.Printf("Sorry, %s is not a valid item.\n", name)
        return
    }
    if vm.balance < price {
        fmt.Printf("Insufficient balance. Please insert at least %d cents.\n", price)
        return
    }
    vm.balance -= price
    vm.stocks[name] -= 1
    vm.records[name] += 1
    fmt.Printf("Dispensing %s. Current balance: %d cents.\n", name, vm.balance)
}

func (vm *VendingMachine) Refund() {
    fmt.Printf("Refunding %d cents.\n", vm.balance)
    vm.balance = 0
}

func (vm *VendingMachine) DisplaySales() {
    fmt.Println("--- Sales Report ---")
    total := 0
    for name, count := range vm.records {
        price, _ := vm.itemPrices[name]
        subtotal := price * count
        fmt.Printf("%s: %d x %d cents = %d cents\n", name, count, price, subtotal)
        total += subtotal
    }
    fmt.Printf("Total sales: %d cents\n", total)
    fmt.Println("--------------------")
}

func main() {
    vm := NewVendingMachine()

    fmt.Println("Welcome to the Vending Machine!")
    vm.DisplayItems()

    vm.InsertCoin(50)
    vm.InsertCoin(100)
    vm.SelectItem("Coke")
    vm.SelectItem("Pepsi")
    vm.SelectItem("Sprite")
    vm.InsertCoin(1000)
    vm.SelectItem("Water")
    vm.SelectItem("Coke")
    vm.Refund()

    vm.DisplaySales()
}
