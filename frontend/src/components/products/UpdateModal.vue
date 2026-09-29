<script setup lang="ts">
import { ref, watch } from 'vue';
import { store } from '../../../wailsjs/go/models';
import { UpdateProduct, ShowMessage, ShowConfirm, DeleteProduct } from "../../../wailsjs/go/main/App"
import Modal from '../Modal.vue';

const props = defineProps<{
  isOpen: boolean
  product?: store.Product
}>();

const emit = defineEmits(["close", "updated"])

const formState = ref({ ...props.product })

watch(
  () => props.product,
  (newValues) => formState.value = { ...newValues }
)

const onSubmit = async () => {
  if (formState.value.ID === undefined) {
    await ShowMessage("Update failed", `Failed to update item "${formState.value.Name}". Please try again later.`)
    emit("close")
    return
  }

  await UpdateProduct({
    ID: formState.value.ID,
    Name: formState.value.Name || "",
    Code: formState.value.Code || "",
    Barcode: formState.value.Barcode || "",
    Description: formState.value.Description || "",
    Price: formState.value.Price || 0,
  })
  await ShowMessage("Update complete", `Item ${formState.value.Name} has been updated.`)

  emit("updated")
  emit("close")
}

const onDelete = async () => {
  if (formState.value.ID === undefined) {
    await ShowMessage("Delete failed", `Failed to delete item. Please try again later.`)
    emit("close")
    return
  }

  const confirm = await ShowConfirm("Delete", "Do you want to delete this item?")
  if (confirm) {
    await DeleteProduct(formState.value.ID)
    emit("updated")
  }

  emit("close")
}

</script>

<template>
  <Modal :isOpen="isOpen" closeBtnText="Cancel" @close="$emit('close')">
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
