<script setup lang="ts">
import { computed, ref, } from 'vue';
import { store } from '../../../wailsjs/go/models';
import { UpdateProduct, ShowMessage, ShowConfirm, DeleteProduct, CreateProduct } from "../../../wailsjs/go/main/App"
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
const isEditing = computed(() => !!props.product)

const onSubmit = async () => {
  if (isEditing.value) {
    if (productID.value === undefined) {
      await ShowMessage("Update failed", `Failed to update item "${formState.value.Name}". Please try again later.`)
      emit("close")
      return
    }
    await UpdateProduct(formState.value, productID.value)
    await ShowMessage("Update complete", `Item ${formState.value.Name} has been updated.`)
  } else {
    await CreateProduct(formState.value)
    await ShowMessage("Product created", `Product ${formState.value.Name} has been created.`)
  }

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
    <h2>{{ isEditing ? "Edit product" : "Create product" }}</h2>
    <form @submit.prevent="onSubmit">
      <div>
        <label>
          Name <input type="text" v-model="formState.Name" />
        </label>
      </div>
      <div>
        <label>
          Code <input type="text" v-model="formState.Code">
        </label>
      </div>
      <div>
        <label>
          Barcode <input type="text" v-model="formState.Barcode">
        </label>
      </div>
      <div>
        <label>
          Category <input type="text" v-model="formState.Description">
        </label>
      </div>
      <div>
        <label>
          Price <input type="number" step="any" v-model="formState.Price">
        </label>
      </div>
      <button type="submit">Save</button>
    </form>
    <button v-if="isEditing" @click="onDelete">Delete</button>
  </Modal>
</template>
