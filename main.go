package main

import (
	"fmt"
)

type Item struct {
	Name     string
	Price    float64
	Discount float64
}

type Describable interface {
	Description() string
}

func (item Item) Description() string {
	if item.Discount > 0 {
		discountedPrice := calculatePrice(item)
		return fmt.Sprintf("%s - %.2f TL (%.0f%% indirimle %.2f TL)", item.Name, item.Price, item.Discount, discountedPrice)
	}
	return fmt.Sprintf("%s - %.2f TL", item.Name, item.Price)
}

func calculatePrice(item Item) float64 {
	if item.Discount > 0 {
		return item.Price * (1 - item.Discount/100)
	}
	return item.Price
}

func totalPrice(items []Item) float64 {
	total := 0.0
	for _, item := range items {
		total += calculatePrice(item)
	}
	return total
}

func main() {
	items := []Item{
		{Name: "Elma", Price: 100, Discount: 10},
		{Name: "Portakal", Price: 1.00},
		{Name: "Armut", Price: 2.50, Discount: 50},
		{Name: "Muz", Price: 5.00},
		{Name: "Karpuz", Price: 10.00, Discount: 5},
		{Name: "Çilek", Price: 3.00, Discount: 10},
		{Name: "Kiraz", Price: 4.00},
		{Name: "Üzüm", Price: 2.00},
	}

	fmt.Println("Ürünler:")
	for _, item := range items {
		fmt.Println(item.Description())
	}

	total := totalPrice(items)
	fmt.Printf("Toplam Fiyat: %.2f TL\n", total)
}
