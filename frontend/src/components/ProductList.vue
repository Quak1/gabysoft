<script lang="ts" setup>
import { computed, onMounted, ref } from 'vue'
import * as Product from "../../wailsjs/go/tasks/Product"
import { store } from "../../wailsjs/go/models"
import UpdateModal from './products/UpdateModal.vue'
import { useCart } from '../composables/useCart'

const products = ref<Array<store.Product>>([])
const isModalOpen = ref(false)
const currentItem = ref<store.Product>()

onMounted(async () => {
  products.value = await Product.GetAll()
})

const categoryFilter = ref("")
const categories = computed(() => {
  const cats = new Set<string>();
  for (const p of products.value) {
    cats.add(p.Description)
  }
  return [...cats].toSorted()
})

const search = ref("")
const filtered = computed(() => {
  const s = search.value.toLowerCase()

  const filtered = categoryFilter.value === ""
    ? products.value
    : products.value.filter((p) => p.Description === categoryFilter.value)

  return filtered.filter((p) => {
    return p.Name.toLowerCase().includes(s)
      || p.Code.toLowerCase().includes(s)
      || p.Barcode.toLowerCase().includes(s)
  })
})

const priceFilterState = ref(0)
const priceConfig = [
  { label: "Price" },
  { label: "Price ↑" },
  { label: "Price ↓" },
]
const priceText = computed(() => priceConfig[priceFilterState.value])
const togglePriceFilterState = () => { priceFilterState.value = (priceFilterState.value + 1) % 3 }

const sorted = computed(() => {
  if (priceFilterState.value === 0) {
    return filtered.value
  }

  return filtered.value.toSorted((a, b) => {
    if (priceFilterState.value === 1) {
      return a.Price - b.Price
    } else {
      return b.Price - a.Price
    }
  })
})

const openUpdateModal = (item: store.Product) => {
  currentItem.value = item
  isModalOpen.value = true
}

const openCreateModal = () => {
  currentItem.value = undefined
  isModalOpen.value = true
}

const handleSave = async () => {
  products.value = await Product.GetAll()
}

const { cartItems, addToCart, removeFromCart } = useCart()

const inCart = computed(() => sorted.value.map(item => ({
  item,
  isInCart: cartItems.value.has(item.ID)
})))
</script>

<template>
  <label>
    Search product:
    <input type="text" name="search" id="search" v-model.trim="search">
  </label>
  <UpdateModal v-if="isModalOpen" :product="currentItem" @close="isModalOpen = false" @save="handleSave" />
  <p>{{ isModalOpen }}</p>
  <button @click="openCreateModal">Create New Product</button>
  <table>
    <thead>
      <tr>
        <th>Name</th>
        <th>Code</th>
        <th>Barcode</th>
        <th>
          <button @click="categoryFilter = ''">Category</button>
          <select name="category" id="category" v-model="categoryFilter">
            <option v-for="cat in categories">{{ cat }}</option>
          </select>
        </th>
        <th>
          <button @click="togglePriceFilterState">
            {{ priceText.label }}
          </button>
        </th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="{ item, isInCart } in inCart" :key="item.ID">
        <td>{{ item.Name }}</td>
        <td>{{ item.Code }}</td>
        <td>{{ item.Barcode }}</td>
        <td>{{ item.Description }}</td>
        <td>{{ item.Price }}</td>
        <td><button @click="openUpdateModal(item)">Edit</button></td>
        <td>
          <button v-if="!isInCart" @click="addToCart(item)">Add to Cart</button>
          <button v-else @click="removeFromCart(item.ID)">Remove from Cart</button>
        </td>
      </tr>
    </tbody>
  </table>
</template>

<style></style>
