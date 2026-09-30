<script setup lang="ts">
import { computed, ref, } from 'vue';
import { store } from '../../../wailsjs/go/models';
import { UpdateProduct, ShowMessage, ShowConfirm, DeleteProduct } from "../../../wailsjs/go/main/App"
import Modal from '../Modal.vue';

const props = defineProps<{
  product?: store.Product
}>();

const emit = defineEmits(["close", "save"])

const formState = ref<store.CreateProductParams>({
  Name: props.product?.Name || "",
  Code: props.product?.Code || "",
  Barcode: props.product?.Barcode || "",
  Description: props.product?.Description || "",
  Price: props.product?.Price || 0,
})
const productID = computed(() => props.product?.ID)

const onSubmit = async () => {
  if (productID.value === undefined) {
    await ShowMessage("Update failed", `Failed to update item "${formState.value.Name}". Please try again later.`)
    emit("close")
    return
  }

  await UpdateProduct({
    Name: formState.value.Name || "",
    Code: formState.value.Code || "",
    Barcode: formState.value.Barcode || "",
    Description: formState.value.Description || "",
    Price: formState.value.Price || 0,
  }, productID.value)
  await ShowMessage("Update complete", `Item ${formState.value.Name} has been updated.`)

  emit("save")
  emit("close")
}

const onDelete = async () => {
  if (productID.value === undefined) {
    await ShowMessage("Delete failed", `Failed to delete item. Please try again later.`)
    emit("close")
    return
  }

  const confirm = await ShowConfirm("Delete", "Do you want to delete this item?")
  if (confirm) {
    await DeleteProduct(productID.value)
    emit("save")
  }

  emit("close")
}
</script>

<template>
  <Modal closeBtnText="Cancel" @close="$emit('close')">
    <form @submit.prevent="onSubmit">
      <div>
        <label for="name">Name</label>
        <input type="text" v-model="formState.Name" />
      </div>
      <div>
        <label for="code">Code</label>
        <input type="text" v-model="formState.Code">
      </div>
      <div>
        <label for="barcode">Barcode</label>
        <input type="text" v-model="formState.Barcode">
      </div>
      <div>
        <label for="category">Category</label>
        <input type="text" v-model="formState.Description">
      </div>
      <div>
        <label for="price">Price</label>
        <input type="number" step="any" v-model="formState.Price">
      </div>
      <button type="submit">Save</button>
    </form>
    <button @click="onDelete">Delete</button>
  </Modal>
</template>
