<script lang="ts" setup>
import { computed, onMounted, ref } from 'vue'
import { GetAllProducts } from "../../wailsjs/go/main/App"
import { store } from "../../wailsjs/go/models"

const products = ref<Array<store.Product>>([])

onMounted(async () => {
  products.value = await GetAllProducts()
})

const categoryFilter = ref("")
const categories = computed(() => {
  const cats = new Set<string>();
  for (const p of products.value) {
    cats.add(p.Description.String)
    console.log(p.Description.String)
  }
  return [...cats].toSorted()
})

const search = ref("")
const filtered = computed(() => {
  const s = search.value.toLowerCase()

  const filtered = categoryFilter.value === ""
    ? products.value
    : products.value.filter((p) => p.Description.String === categoryFilter.value)

  return filtered.filter((p) => {
    return p.Name.toLowerCase().includes(s)
      || p.Code.String.toLowerCase().includes(s)
      || p.Barcode.String.toLowerCase().includes(s)
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


</script>

<template>
  <input type="text" name="search" id="search" v-model.trim="search">
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
      <tr v-for="entry in sorted" :key="entry.ID">
        <td>{{ entry.Name }}</td>
        <td>{{ entry.Code.String }}</td>
        <td>{{ entry.Barcode.String }}</td>
        <td>{{ entry.Description.String }}</td>
        <td>{{ entry.Price }}</td>
      </tr>
    </tbody>
  </table>
</template>

<style></style>
