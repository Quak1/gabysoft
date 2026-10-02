import { ref } from "vue";
import { store } from "../../wailsjs/go/models";

interface CartEntry {
  item: store.Product
  quantity: number
  cost: number
  costTax: number
}

const cartItems = ref<Map<number, CartEntry>>(new Map())
const TAX = 0.16

const getCosts = (price: number, quantity: number) => {
  return [price * quantity, price * quantity * (1 + TAX)]

}
const updateCosts = (entry: CartEntry) => {
  [entry.cost, entry.costTax] = getCosts(entry.item.Price, entry.quantity)
}

export function useCart() {
  const addToCart = (item: store.Product) => {
    let entry = cartItems.value.get(item.ID)
    if (entry) {
      entry.quantity++
      updateCosts(entry)
    } else {
      const [cost, costTax] = getCosts(item.Price, 1)
      cartItems.value.set(item.ID, { item, quantity: 1, cost, costTax })
    }
  }

  const removeFromCart = (itemId: number) => {
    cartItems.value.delete(itemId)
  }

  const updateQuantity = (itemId: number, newQuantity: number) => {
    if (newQuantity <= 0) {
      removeFromCart(itemId)
      return
    }

    const entry = cartItems.value.get(itemId)
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
