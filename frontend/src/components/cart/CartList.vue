<script setup lang="ts">
import { useCart } from '../../composables/useCart';

const emit = defineEmits(['test'])
const { cartItems, removeFromCart, updateQuantity } = useCart()

const handleQuantityChange = (itemId: number, event: Event) => {
  const target = event.target as HTMLInputElement | null
  if (target) {
    const newQuantity = parseFloat(target.value) || 0
    updateQuantity(itemId, newQuantity)
  }
}
</script>

<template>
  <table>
    <thead>
      <tr>
        <th>Name</th>
        <th>Code</th>
        <th>Barcode</th>
        <th>Category</th>
        <th>Price</th>
        <th>Quantity</th>
        <th>Cost</th>
        <th>Cost TAX</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="{ item, quantity, cost, costTax } in cartItems" :key="item.ID">
        <td>{{ item.Name }}</td>
        <td>{{ item.Code }}</td>
        <td>{{ item.Barcode }}</td>
        <td>{{ item.Description }}</td>
        <td>${{ item.Price }}</td>
        <td>
          <input type="number" step="any" :value="quantity" @change="handleQuantityChange(item.ID, $event)" />
        </td>
        <td>${{ cost.toFixed(2) }}</td>
        <td>${{ costTax.toFixed(2) }}</td>
        <td><button @click="removeFromCart(item.ID)">Remove</button></td>
      </tr>
    </tbody>
  </table>
</template>
