<script lang="ts" setup>
import { computed, onMounted, ref } from 'vue'
import { GetAllProducts } from "../../wailsjs/go/main/App"
import { store } from "../../wailsjs/go/models"
import UpdateModal from './products/UpdateModal.vue'

const products = ref<Array<store.Product>>([])
const isModalOpen = ref(false)
const currentItem = ref<store.Product>()

onMounted(async () => {
  products.value = await GetAllProducts()
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
  products.value = await GetAllProducts()
}

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
      <tr v-for="item in sorted" :key="item.ID">
        <td>{{ item.Name }}</td>
        <td>{{ item.Code }}</td>
        <td>{{ item.Barcode }}</td>
        <td>{{ item.Description }}</td>
        <td>{{ item.Price }}</td>
        <td><button @click="openUpdateModal(item)">Edit</button></td>
      </tr>
    </tbody>
  </table>
</template>

<style></style>
