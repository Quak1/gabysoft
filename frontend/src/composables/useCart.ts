import { ref } from "vue";
import { store } from "../../wailsjs/go/models";

interface CartEntry {
  item: store.Product
  quantity: number
  cost: number
  costTax: number
}

const cartItems = ref<CartEntry[]>([])
const TAX = 0.16

const updateCosts = (entry: CartEntry) => {
  entry.cost = entry.item.Price * entry.quantity
  entry.costTax = entry.cost + entry.cost * TAX
}

export function useCart() {
  const addToCart = (item: store.Product) => {
    let entry = cartItems.value.find((e) => e.item.ID === item.ID)
    if (entry) {
      entry.quantity++
    } else {
      const i = cartItems.value.push({ item, quantity: 1, cost: 0, costTax: 0 })
      entry = cartItems.value[i - 1]
    }

    updateCosts(entry)
  }

  const removeFromCart = (itemId: number) => {
    cartItems.value = cartItems.value.filter((e) => e.item.ID !== itemId)
  }

  const updateQuantity = (itemId: number, newQuantity: number) => {
    if (newQuantity <= 0) {
      removeFromCart(itemId)
      return
    }

    const entry = cartItems.value.find(e => e.item.ID === itemId)
    if (entry) {
      entry.quantity = newQuantity
      updateCosts(entry)
    }
  }

  return {
    cartItems,
    addToCart,
    removeFromCart,
    updateQuantity
  }
}
