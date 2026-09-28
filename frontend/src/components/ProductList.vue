<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import { GetAllProducts } from "../../wailsjs/go/main/App"
import { store } from "../../wailsjs/go/models"

let products = ref<Array<store.Product>>()

onMounted(async () => {
  products.value = await GetAllProducts()
})

</script>

<template>
  <table>
    <thead>
      <tr>
        <th>Name</th>
        <th>Code</th>
        <th>Barcode</th>
        <th>String</th>
        <th>Price</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="entry in products" :key="entry.ID">
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
