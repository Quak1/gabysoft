<script setup lang="ts">
import { computed, ref, } from 'vue';
import { store } from '../../../wailsjs/go/models';
import * as Product from "../../../wailsjs/go/tasks/Product"
import * as App from '../../../wailsjs/go/main/App'
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
      await App.ShowMessage("Update failed", `Failed to update item "${formState.value.Name}". Please try again later.`)
      emit("close")
      return
    }
    await Product.Update(formState.value, productID.value)
    await App.ShowMessage("Update complete", `Item ${formState.value.Name} has been updated.`)
  } else {
    await Product.Create(formState.value)
    await App.ShowMessage("Product created", `Product ${formState.value.Name} has been created.`)
  }

  emit("save")
  emit("close")
}

const onDelete = async () => {
  if (productID.value === undefined) {
    await App.ShowMessage("Delete failed", `Failed to delete item. Please try again later.`)
    emit("close")
    return
  }

  const confirm = await App.ShowConfirm("Delete", "Do you want to delete this item?")
  if (confirm) {
    await Product.Delete(productID.value)
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
