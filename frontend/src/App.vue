<script lang="ts" setup>
import { ref } from 'vue';
import CartList from './components/cart/CartList.vue';
import ProductList from './components/ProductList.vue';

const tabs = {
  ProductList,
  CartList,
}

const currentTab = ref<keyof typeof tabs>('ProductList')
</script>

<template>
  <div>
    <button v-for="(_, tab) in tabs" :key="tab" @click="currentTab = tab">{{ tab }}</button>
  </div>
  <KeepAlive include="ProductList">
    <component :is="tabs[currentTab]"></component>
  </KeepAlive>
</template>

<style></style>
